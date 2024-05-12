package grunt

import (
	"crypto/sha1"
	"fmt"
	"net"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Gophercraft/core/crypto/hashutil"
	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/version"
)

//go:generate stringer -type=ClientEvent
type ClientEvent uint8

const (
	ClientEventNone ClientEvent = iota
	ClientServerOffline
	ClientDisconnected
	ClientConnecting
	ClientConnected
	ClientHandshaking
	ClientSuccess
	ClientFailure
	ClientRequestingRealmlist
	ClientGotRealmlist
	NumClientEvent
)

type ClientHandlerFunc func(*Client)

type Client struct {
	state        int
	options      ClientOptions
	client_proof []byte
	session_key  []byte
	credentials  credentials
	list         []Realm
	link         ClientLink
	handlers     [int(NumClientEvent)][]ClientHandlerFunc
}

type ClientOptions struct {
	Endpoint     string
	Build        version.Build
	OS           OS
	Architecture Architecture
	Program      Program
	Locale       Locale
	IP           net.IP
	TimezoneBias int32
	// the time
	LinkInterval time.Duration
}

type credentials struct {
	user          string
	password      string
	identity_hash []byte
}

func NewClient(options *ClientOptions) (client *Client) {
	// create client
	client = new(Client)
	// set options
	if options != nil {
		client.options = *options
	}

	// set default options
	if client.options.Architecture == 0 {
		client.options.Architecture = X86
	}
	if client.options.OS == 0 {
		client.options.OS = Windows
	}
	if client.options.Locale == 0 {
		client.options.Locale = Locale_enUS
	}
	if client.options.Program == 0 {
		client.options.Program = WoW
	}
	if client.options.Endpoint == "" {
		client.options.Endpoint = "127.0.0.1:3724"
	}
	if client.options.IP == nil {
		client.options.IP = net.IPv4(127, 0, 0, 1)
	}
	if client.options.Build == 0 {
		client.options.Build = 5875
	}

	return
}

func (client *Client) Credentials(user, password string) (err error) {
	if !utf8.ValidString(user) || user == "" {
		return fmt.Errorf("grunt: invalid username")
	}

	if !utf8.ValidString(password) || password == "" {
		return fmt.Errorf("grunt: invalid password")
	}

	client.credentials = credentials{
		user:     strings.ToUpper(user),
		password: strings.ToUpper(password),
	}

	return
}

func (client *Client) Handle(event ClientEvent, handler_func ClientHandlerFunc) {
	client.handlers[event] = append(client.handlers[event], handler_func)
}

func (client *Client) dispatch(event ClientEvent) {
	for _, handler := range client.handlers[event] {
		handler(client)
	}
}

func (client *Client) set_state(state int) {
	client.state = state
}

func (client *Client) get_state() int {
	return client.state
}

func (client *Client) handle_auth_logon_challenge(logon_challenge *AuthLogonChallenge_Server) (err error) {
	if client.get_state() != state_unauthorized {
		return fmt.Errorf("grunt: server sent challenge while we're authorized")
	}

	client.credentials.identity_hash = hashutil.H(sha1.New, hashutil.Credentials(client.credentials.user, client.credentials.password))
	_, K, A, M1 := srp.HashCalculate(
		client.credentials.user,
		client.credentials.identity_hash,
		logon_challenge.ServerPublicKey[:],
		logon_challenge.LargeSafePrime,
		logon_challenge.Generator,
		logon_challenge.Salt[:])

	client.session_key = K
	client.client_proof = M1
	var logon_proof AuthLogonProof_Client
	copy(logon_proof.ClientPublicKey[:], A[:32])
	copy(logon_proof.ClientProof[:], M1[:20])

	// see if we can do a version challenge
	build_info := client.options.Build.BuildInfo()
	if build_info != nil {
		var version_seed []byte
		switch client.options.OS {
		case Windows:
			version_seed = build_info.WinChecksumSeed
		case MacOS:
			version_seed = build_info.MacChecksumSeed
		}

		if len(version_seed) > 0 {
			version_proof := sha1.New()
			version_proof.Write(logon_proof.ClientPublicKey[:])
			version_proof.Write(version_seed)
			copy(logon_proof.VersionProof[:], version_proof.Sum(nil))
		}

	}
	if err = WriteMessageType(&client.link, LogonProof); err != nil {
		return
	}
	if err = WriteAuthLogonProof_Client(&client.link, &logon_proof); err != nil {
		return
	}
	if err = client.link.Send(); err != nil {
		return
	}
	client.dispatch(ClientHandshaking)
	client.set_state(state_logon_challenging)

	return
}

func (client *Client) handle_auth_logon_proof(logon_proof *AuthLogonProof_Server) (err error) {
	if client.get_state() != state_logon_challenging {
		return fmt.Errorf("grunt: server sent proof but we haven't sent ours")
	}

	if logon_proof.LoginResult != LoginOk {
		client.dispatch(ClientFailure)
		return fmt.Errorf("grunt: server sent invalid login proof")
	}

	client.set_state(state_authorized)
	client.dispatch(ClientSuccess)

	return
}

func (client *Client) handle_realm_list(realm_list *RealmList_Server) (err error) {
	client.list = realm_list.Realms

	client.dispatch(ClientGotRealmlist)

	return
}

func (client *Client) operate() (err error) {
	var challenge AuthLogonChallenge_Client
	challenge.Protocol = 8
	challenge.Info.AccountName = client.credentials.user
	challenge.Info.Architecture = client.options.Architecture
	challenge.Info.Locale = client.options.Locale
	challenge.Info.OS = client.options.OS
	challenge.Info.Program = client.options.Program
	challenge.Info.IP = client.options.IP
	challenge.Info.TimezoneBias = client.options.TimezoneBias
	challenge.Info.Build = client.options.Build
	challenge.Info.Version = Version(client.options.Build)

	if err = WriteMessageType(&client.link, LogonChallenge); err != nil {
		return
	}
	if err = WriteAuthLogonChallenge_Client(&client.link, &challenge); err != nil {
		return
	}
	if err = client.link.Send(); err != nil {
		return
	}

	for {
		if client.get_state() == state_authorized {
			var request RealmList_Client
			request.Request = 0
			if err = WriteMessageType(&client.link, RealmList); err != nil {
				return
			}
			if err = WriteRealmList_Client(&client.link, &request); err != nil {
				break
			}
			if err = client.link.Send(); err != nil {
				return
			}
		}

		var message_type MessageType
		message_type, err = client.link.ReadMessageType()
		if err != nil {
			break
		}

		switch message_type {
		case LogonChallenge:
			var logon_challenge AuthLogonChallenge_Server
			if err = ReadAuthLogonChallenge_Server(&client.link, &logon_challenge); err != nil {
				break
			}
			if err = client.handle_auth_logon_challenge(&logon_challenge); err != nil {
				break
			}
			continue
		case LogonProof:
			var logon_proof AuthLogonProof_Server
			if err = ReadAuthLogonProof_Server(&client.link, client.options.Build, &logon_proof); err != nil {
				break
			}
			if err = client.handle_auth_logon_proof(&logon_proof); err != nil {
				break
			}
			continue
		case RealmList:
			var realm_list RealmList_Server
			if err = ReadRealmList_Server(&client.link, client.options.Build, &realm_list); err != nil {
				break
			}

			if err = client.handle_realm_list(&realm_list); err != nil {
				break
			}

			time.Sleep(10 * time.Second)
		}
	}

	// link has been broken
	client.set_state(state_unauthorized)
	client.dispatch(ClientDisconnected)
	// return error
	return
}

func (client *Client) Run() (err error) {
	client.dispatch(ClientConnecting)

	for {
		err = client.link.Establish(client.options.Endpoint)
		if err != nil {
			time.Sleep(client.options.LinkInterval)
			continue
		}
		client.dispatch(ClientConnected)

		if err = client.operate(); err != nil {
			continue
		}
	}
}

func (client *Client) RealmList() (list []Realm) {
	list = client.list
	return
}
