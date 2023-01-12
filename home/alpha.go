package home

import (
	"context"
	"fmt"
	"strings"

	"github.com/Gophercraft/core/home/alphalist"
	"github.com/Gophercraft/core/home/login"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) listenAlphaList() {
	if s.Config.AlphaRealmlistListen != "" {
		list := alphalist.NewServer(s.Config.AlphaRealmlistListen, s)
		if err := list.Start(); err != nil {
			log.Warn(err)
		}
	}
}

func sanitizeCred(i string) string {
	i = strings.ReplaceAll(i, "\r", "")
	i = strings.ReplaceAll(i, " ", "")
	return i
}

func parseAlphaLogin(lg string) (username, password string, err error) {
	lines := strings.Split(lg, "\n")

	if len(lines) < 2 {
		log.Dump("lines", lines)
		err = fmt.Errorf("cannot parse: invalid alpha login")
		return
	}

	username = sanitizeCred(lines[0])
	password = sanitizeCred(lines[1])

	return
}

func (h *rpcServer) verifyWorldAlpha(ctx context.Context, req *rpcnet.VerifyWorldQuery) (*rpcnet.VerifyWorldResponse, error) {
	us, pass, err := parseAlphaLogin(req.Account)
	if err != nil {
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Disabled,
		}, err
	}

	creds := login.Credentials(us, pass)

	var user models.Account
	found, _ := h.DB.Where("username = ?", us).Get(&user)
	if !found {
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Unauthorized,
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword(user.WebLoginHash, creds); err != nil {
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Disabled,
		}, err
	}

	var ga models.GameAccount
	found, _ = h.DB.Where("owner = ?", user.ID).Where("active = 1").Get(&ga)
	if !found {
		return &rpcnet.VerifyWorldResponse{}, fmt.Errorf("home: no active game account")
	}

	return &rpcnet.VerifyWorldResponse{
		Status:      rpcnet.Status_OK,
		Tier:        user.Tier,
		Account:     user.ID,
		GameAccount: ga.ID,
	}, nil
}
