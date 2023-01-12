package home

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"time"

	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
)

func (s *Server) SignOutAll(account uint64) {
	s.DB.Where("account = ?", account).Delete(new(models.WebToken))
}

func (s *Server) genToken(account uint64) string {
	data := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(data)
}

func (s *Server) HandleWebLogin(username, password string) (*models.Account, string, error) {
	account, _, err := s.GetAccount(username)
	if err != nil {
		return nil, "", err
	}

	webToken, err := s.AccountLogin(account, username, password)
	return account, webToken, err
}

func (s *Server) AccountLogin(account *models.Account, username, password string) (string, error) {
	if !login.Verify(username, password, account.WebLoginHash) {
		return "", fmt.Errorf("home: %s verification failed", username)
	}

	s.SignOutAll(account.ID)

	var token models.WebToken

	token.Account = account.ID
	token.Expiry = time.Now().Add(time.Hour * 12)
	token.Token = s.genToken(account.ID)

	s.DB.Insert(&token)

	return token.Token, nil
}

func (s *Server) QueryTokenAccount(token string) (*models.Account, error) {
	var acc models.Account

	found, err := s.DB.Where("token = ?", token).Get(&acc)
	if err != nil {
		panic(err)
	}

	if found {
		return &acc, nil
	}

	return nil, fmt.Errorf("home: couldn't find token")
}
