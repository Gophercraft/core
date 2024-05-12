package rest

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/Gophercraft/core/bnet/pb/login"
	"github.com/Gophercraft/core/bnet/util"
	"github.com/Gophercraft/core/crypto/hashutil"
	"github.com/Gophercraft/core/crypto/srp"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/service/bnet_rest"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *Provider) LoginSRP(user_info *bnet_rest.UserInfo, login_form *login.LoginForm) (srp_login_challenge *login.SrpLoginChallenge, err error) {
	// Block unwanted users
	if err = provider.check_user_info(user_info); err != nil {
		return
	}

	account_name := get_login_form_input(login_form, "account_name")
	platform_id := login_form.GetPlatformId()

	// Select hash function
	hash_function := "SHA-512"

	// Lookup identity hash

	var (
		account models.Account
		found   bool
	)
	found, err = provider.home_db.Table("Account").Where(query.Eq("Username", strings.ToLower(account_name))).Get(&account)
	if err != nil {
		return
	}

	if !found {
		err = bnet_rest.ErrorMessage(http.StatusOK, "INVALID_ACCOUNT_OR_CREDENTIALS", "We couldn't verify your account with that information.")
		return
	}

	identity_hash := account.Identity_PBKDF2_SHA_512
	salt := account.Identity_PBKDF2_Salt

	// generate SRP session parameters
	srp_session := new(srp.Session)
	switch hash_function {
	case "SHA-256":
		srp_session.Hash = sha256.New
	case "SHA-512":
		srp_session.Hash = sha512.New
	default:
		panic(hash_function)
	}

	srp_session.Version = 2
	srp_session.Iterations = 15000
	srp_session.LargeSafePrime = srp.NewIntFromHexString("AC6BDB41324A9A9BF166DE5E1389582FAF72B6651987EE07FC3192943DB56050A37329CBB4A099ED8193E0757767A13DD52312AB4B03310DCD7F48A9DA04FD50E8083969EDB767B0CF6095179A163AB3661A05FBD5FAAAE82918A9962F0B93B855F97993EC975EEAA80D740ADBF4FF747359D041D5C33EA71D281E446B14773BCA97B43A23FB801676BD207A436C6481F1D2B9078717461A5B9D32E688F87748544523B524B0D57D5EA77A2775D2ECFA032CFBDBF52FB3786160279004E57AE6AF874E7303CE53299CCC041C7BC308D82A5698F3A8D0C38271AE35F8E9DBFBB694B5C803D89F7AE435DE236D525F54759B65E372FCD68EF20FA7111F9E4AFF73")
	srp_session.Generator = srp.NewInt(2)
	srp_session.Multiplier = srp.DeriveMultiplier(srp_session.Hash, srp_session.LargeSafePrime, srp_session.Generator)
	srp_session.Salt = salt

	_, srp_session.Verifier = srp.CalculateVerifier2(srp_session.Hash, salt, identity_hash, srp_session.Generator, srp_session.LargeSafePrime)

	srp_session.PrivateB = srp.CalculatePrivateB(srp_session.LargeSafePrime)
	srp_session.PublicB = srp.CalculatePublicB(srp_session.Verifier, srp_session.PrivateB, srp_session.LargeSafePrime, srp_session.Generator, srp_session.Multiplier)

	// Save SRP session parameters for this account
	// (store platform ID along with it)
	if err = provider.store_srp_session(user_info, account_name, hash_function, platform_id, srp_session); err != nil {
		log.Warn("could not store srp session", err)
		return
	}

	// srp_username := H(TOUPPER(login_username))
	srp_username := strings.ToUpper(
		hex.EncodeToString(
			hashutil.H(sha512.New, []byte(strings.ToUpper(account_name))),
		),
	)

	log.Println("srp_username", srp_username)

	srp_login_challenge = new(login.SrpLoginChallenge)

	// Create SRP challenge response
	util.Set(&srp_login_challenge.Version, srp_session.Version)
	util.Set(&srp_login_challenge.Iterations, srp_session.Iterations)
	util.Set(&srp_login_challenge.Modulus, srp_session.LargeSafePrime.HexString())
	util.Set(&srp_login_challenge.Generator, srp_session.Generator.HexString())
	util.Set(&srp_login_challenge.HashFunction, hash_function)
	util.Set(&srp_login_challenge.Username, srp_username)
	util.Set(&srp_login_challenge.Salt, hex.EncodeToString(salt))
	util.Set(&srp_login_challenge.Public_B, srp_session.PublicB.HexString())
	// util.Set(&login_challenge.EligibleCredentialUpgrade, true)

	return
}
