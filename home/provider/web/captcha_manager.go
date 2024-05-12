package web

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/steambap/captcha"
)

const sweep_interval time.Duration = 20 * time.Second
const captcha_life time.Duration = 100 * time.Second

type captcha_challenge struct {
	created time.Time
	text    string
	data    []byte
}

func (challenge *captcha_challenge) is_expired() bool {
	return time.Since(challenge.created) > captcha_life
}

type captcha_manager struct {
	last_sweep       time.Time
	running          atomic.Bool
	guard_challenges sync.Mutex
	challenges       map[string]*captcha_challenge
}

func (mgr *captcha_manager) sweep() {

	for id, challenge := range mgr.challenges {
		if challenge.is_expired() {
			delete(mgr.challenges, id)
		}
	}
	mgr.last_sweep = time.Now()

}

func (mgr *captcha_manager) clean() {
	mgr.guard_challenges.Lock()
	if time.Since(mgr.last_sweep) > sweep_interval {
		mgr.sweep()
	}
	mgr.guard_challenges.Unlock()
}

func new_random_id() string {
	var data [32]byte
	if _, err := io.ReadFull(rand.Reader, data[:]); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(data[:])
}

func (manager *captcha_manager) new_challenge() (id string, err error) {
	manager.clean()

	id = new_random_id()

	challenge := new(captcha_challenge)
	challenge.created = time.Now()

	var buffer bytes.Buffer

	var captcha_data *captcha.Data
	captcha_data, err = captcha.New(280, 80)
	if err = captcha_data.WriteImage(&buffer); err != nil {
		return
	}

	challenge.text = captcha_data.Text
	challenge.data = buffer.Bytes()
	manager.guard_challenges.Lock()
	manager.challenges[id] = challenge
	manager.guard_challenges.Unlock()

	return
}

func (manager *captcha_manager) verify(id, solution string) bool {
	var found bool
	var challenge *captcha_challenge
	manager.guard_challenges.Lock()
	challenge, found = manager.challenges[id]
	if !found {
		manager.guard_challenges.Unlock()
		return false
	}

	valid := challenge.text == solution
	delete(manager.challenges, id)
	manager.guard_challenges.Unlock()

	return valid
}

func (manager *captcha_manager) get_image(id string) (data []byte, err error) {
	var found bool
	var challenge *captcha_challenge
	manager.guard_challenges.Lock()
	challenge, found = manager.challenges[id]
	manager.guard_challenges.Unlock()
	if !found {
		err = fmt.Errorf("captcha not found")
		return
	}

	return challenge.data, nil
}

func new_captcha_manager() (manager *captcha_manager) {
	manager = new(captcha_manager)
	manager.challenges = make(map[string]*captcha_challenge)
	return
}
