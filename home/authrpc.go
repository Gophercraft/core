package home

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"time"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/log"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/vsn"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authCheckSeed        = []byte{0xC5, 0xC6, 0x98, 0x95, 0x76, 0x3F, 0x1D, 0xCD, 0xB6, 0xA1, 0x37, 0x28, 0xB3, 0x12, 0xFF, 0x8A}
	sessionKeySeed       = []byte{0x58, 0xCB, 0xCF, 0x40, 0xFE, 0x2E, 0xCE, 0xA6, 0x5A, 0x90, 0xB8, 0x01, 0x68, 0x6C, 0x28, 0x0B}
	continuedSessionSeed = []byte{0x16, 0xAD, 0x0C, 0xD4, 0x46, 0xF9, 0x4F, 0xB2, 0xEF, 0x7D, 0xEA, 0x2A, 0x17, 0x66, 0x4D, 0x2F}
	encryptionKeySeed    = []byte{0xE9, 0x75, 0x3C, 0x50, 0x90, 0x93, 0x61, 0xDA, 0x3B, 0x07, 0xEE, 0xFA, 0xFF, 0x9D, 0x41, 0xB8}
)

type rpcServer struct {
	*Server
	rpcnet.UnimplementedHomeServiceServer
}

func (h *rpcServer) GetVersionData(ctx context.Context, _ *empty.Empty) (*rpcnet.VersionData, error) {
	return &rpcnet.VersionData{
		CoreVersion: vsn.GophercraftVersion.String(),
	}, nil
}

func (h *rpcServer) SignIn(ctx context.Context, req *rpcnet.Credentials) (*rpcnet.SessionInfo, error) {
	acc, token, err := h.HandleWebLogin(req.Account, req.Password)
	if err != nil {
		return &rpcnet.SessionInfo{
			Status: rpcnet.Status_Unauthorized,
		}, err
	}

	return &rpcnet.SessionInfo{
		Status:   rpcnet.Status_OK,
		WebToken: token,
		Tier:     acc.Tier,
	}, nil
}

func (h *rpcServer) Ping(ctx context.Context, req *rpcnet.PingMsg) (*rpcnet.PingMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func (h *rpcServer) CheckPeerIdentity(ctx context.Context, realmID uint64) (*rpcnet.StatusMsg, error) {
	finger, err := rpcnet.GetPeerFingerprint(ctx)
	if err != nil {
		return nil, err
	}

	var realm models.EnlistedRealm

	ok, err := h.DB.Where("id = ?", realmID).Get(&realm)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, status.Errorf(codes.NotFound, "no realm exists corresponding to %d. You can solve this by registering via the web portal", realmID)
	}

	if !rpcnet.FingerprintsEqual(finger, realm.Fingerprint) {
		return nil, status.Errorf(codes.PermissionDenied, "realm %d exists, but this realm's fingerprint does not match the one on record.", realmID)
	}

	return rpcnet.Code(rpcnet.Status_OK), nil
}

func (h *rpcServer) AnnounceRealm(ctx context.Context, req *rpcnet.AnnounceRealmMsg) (*rpcnet.StatusMsg, error) {
	smsg, err := h.CheckPeerIdentity(ctx, req.RealmID)
	if err != nil {
		return smsg, err
	}

	var rlm models.Realm
	found, err := h.DB.Where("id = ?", req.RealmID).Get(&rlm)
	if err != nil {
		panic(err)
	}

	rlm.Name = req.RealmName
	rlm.Address = req.Address
	rlm.RedirectAddress = req.RedirectAddress
	rlm.ClientVersion = vsn.Build(req.Build)
	rlm.Description = req.RealmDescription
	rlm.ActivePlayers = req.ActivePlayers
	rlm.Type = config.RealmType(req.Type)
	rlm.LastUpdated = time.Now()

	if !found {
		rlm.ID = req.RealmID
		h.DB.Insert(&rlm)
	} else {
		if _, err := h.DB.ID(req.RealmID).AllCols().Update(&rlm); err != nil {
			panic(err)
		}
	}

	return smsg, err
}

func (h *rpcServer) VerifyWorld(ctx context.Context, req *rpcnet.VerifyWorldQuery) (*rpcnet.VerifyWorldResponse, error) {
	smsg, err := h.CheckPeerIdentity(ctx, req.RealmID)
	if err != nil {
		return &rpcnet.VerifyWorldResponse{
			Status: smsg.Status,
		}, err
	}

	// Alpha lacks a sophisticated login system
	if vsn.Build(req.Build) == vsn.Alpha {
		return h.verifyWorldAlpha(ctx, req)
	}

	var sessionKey []byte

	var user models.Account
	found, _ := h.DB.Where("username = ?", req.Account).Get(&user)
	if !found {
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Unauthorized,
		}, nil
	}

	build := vsn.Build(req.Build)

	var sk models.SessionKey
	found, err = h.DB.Where("id = ?", user.ID).Get(&sk)
	if !found {
		log.Println("Database error: ", err)
		log.Warn("connection authentication failure: no session key", req.Account, "from", req.IP)
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Unauthorized,
		}, nil
	}

	// Use new calculations
	if build.AddedIn(vsn.NewAuthSystem) {
		buildInfo := build.BuildInfo()
		if buildInfo == nil {
			err := fmt.Errorf("build info for %s not found", build)
			return &rpcnet.VerifyWorldResponse{
				Status: rpcnet.Status_Unauthorized,
			}, err
		}
		if len(buildInfo.Win64AuthSeed) == 0 || len(buildInfo.Mac64AuthSeed) == 0 {
			err := fmt.Errorf("auth seed for %s not found", build)
			return &rpcnet.VerifyWorldResponse{
				Status: rpcnet.Status_Unauthorized,
			}, err
		}

		localChallenge := req.Seed
		serverChallenge := req.Salt
		digest := req.Digest

		sessionKeyHash := sha256.New()
		skl, _ := sessionKeyHash.Write(sk.K)
		if skl != 64 {
			panic("invalid key length")
		}

		log.Dump("req.Digest", req.Digest)
		log.Dump("localChallenge", localChallenge)
		log.Dump("serverChallenge", serverChallenge)
		log.Dump("buildInfo.Win64AuthSeed", buildInfo.Win64AuthSeed)

		if user.Platform == "Wn64" {
			sessionKeyHash.Write(buildInfo.Win64AuthSeed)
		} else if user.Platform == "Mc64" {
			sessionKeyHash.Write(buildInfo.Mac64AuthSeed)
		} else {
			return &rpcnet.VerifyWorldResponse{
				Status: rpcnet.Status_Unauthorized,
			}, fmt.Errorf("invalid user platform %s", user.Platform)
		}

		digestKeyHash := sessionKeyHash.Sum(nil)

		hmc := hmac.New(sha256.New, digestKeyHash)
		hmc.Write(localChallenge)  //localChallenge
		hmc.Write(serverChallenge) //serverChallenge
		hmc.Write(authCheckSeed)
		authCheckHash := hmc.Sum(nil)

		if subtle.ConstantTimeCompare(authCheckHash[:24], digest[:24]) == 0 {
			err := errors.New(fmt.Sprintln("connection authentication failure: phony connection attempt to", req.Account, "from", req.IP))
			log.Warn(err)
			return &rpcnet.VerifyWorldResponse{
				Status: rpcnet.Status_Unauthorized,
			}, err
		}

		keyDataDigest := sha256.Sum256(sk.K)

		sessionKeyHmac := hmac.New(sha256.New, keyDataDigest[:])
		sessionKeyHmac.Write(serverChallenge)
		sessionKeyHmac.Write(localChallenge)
		sessionKeyHmac.Write(sessionKeySeed)

		sessionKey = make([]byte, 40)
		skg := crypto.NewSessionKeyGenerator(sha256.New, sessionKeyHmac.Sum(nil))
		skg.Read(sessionKey)

		log.Dump("sessionKey", sessionKey)

		encryptKeyGen := hmac.New(sha256.New, sessionKey)
		encryptKeyGen.Write(localChallenge)
		encryptKeyGen.Write(serverChallenge)
		encryptKeyGen.Write(encryptionKeySeed)

		encryptKeyHash := encryptKeyGen.Sum(nil)

		sessionKey = encryptKeyHash[:16]
	} else {
		// Use old calculation
		digest := hash(
			[]byte(req.Account),
			[]byte{0, 0, 0, 0},
			req.Seed,
			req.Salt,
			sk.K,
		)

		if subtle.ConstantTimeCompare(digest, req.Digest) == 0 {
			err := errors.New(fmt.Sprintln("connection authentication failure: phony connection attempt to", req.Account, "from", req.IP))
			log.Warn(err)
			return &rpcnet.VerifyWorldResponse{
				Status: rpcnet.Status_Unauthorized,
			}, err
		}

		sessionKey = sk.K
	}

	var ga models.GameAccount

	if req.GameAccount == "" {
		// If a specific GameAccount is not being requested, select one where Active is true.
		// TODO: ensure that only one account is active.
		found, _ = h.DB.Where("owner = ?", user.ID).Where("active = 1").Get(&ga)
	} else {
		found, _ = h.DB.Where("owner = ?", user.ID).Where("name = ?", req.GameAccount).Get(&ga)
	}

	if !found {
		return &rpcnet.VerifyWorldResponse{
			Status: rpcnet.Status_Unauthorized,
		}, fmt.Errorf("no GameAccount detected")
	}

	return &rpcnet.VerifyWorldResponse{
		Status:      rpcnet.Status_OK,
		Tier:        user.Tier,
		SessionKey:  sessionKey,
		Account:     user.ID,
		GameAccount: ga.ID,
	}, nil
}

func (h *rpcServer) EnlistRealm(ctx context.Context, req *rpcnet.EnlistRealmRequest) (*rpcnet.EnlistRealmResponse, error) {
	start := time.Now()
	timeout := 5 * time.Second

	// TODO: impose increased penalty for incorrect login

	acc, err := h.QueryTokenAccount(req.WebToken)
	if err != nil {
		fixedSleep(start, timeout)
		return &rpcnet.EnlistRealmResponse{
			Status: rpcnet.Status_Unauthorized,
		}, err
	}

	response, err := h.HandleEnlistRealm(acc, &models.WebEnlistRealmRequest{
		Name:        req.RealmName,
		Fingerprint: req.RealmFingerprint,
	})

	if err != nil {
		return &rpcnet.EnlistRealmResponse{
			Status: rpcnet.Status_Disabled,
		}, err
	}

	return &rpcnet.EnlistRealmResponse{
		Status:  rpcnet.Status_OK,
		RealmID: response.RealmID,
	}, nil
}

func (h *rpcServer) SubmitTicket(ctx context.Context, req *rpcnet.SupportTicket) (*rpcnet.TicketSubmissionResult, error) {
	return &rpcnet.TicketSubmissionResult{}, nil
}

func (h *rpcServer) QueryTickets(ctx context.Context, req *rpcnet.TicketQuery) (*rpcnet.TicketQueryResponse, error) {
	return &rpcnet.TicketQueryResponse{}, nil
}

func hash(input ...[]byte) []byte {
	bt := sha1.Sum(bytes.Join(input, nil))
	return bt[:]
}
