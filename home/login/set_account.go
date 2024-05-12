package login

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"strings"

	"github.com/Gophercraft/core/crypto/hashutil"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

func Credentials(username, password string) []byte {
	I := strings.ToUpper(username)
	P := strings.ToUpper(password)
	return []byte(I + ":" + P)
}

// Returns true if matches hash
func Verify(username, password string, againstHash []byte) bool {
	webLogin := Credentials(username, password)
	ok := bcrypt.CompareHashAndPassword(againstHash, webLogin) == nil
	return ok
}

// Reset an accounts username-password hashes as well as its authorization tier
func SetAccount(account *models.Account, email, username, password string, tier auth.AccountTier) error {
	if err := PasswordValidate(password); err != nil {
		return err
	}

	if err := UsernameValidate(username); err != nil {
		return err
	}

	if password == "" {
		return ErrPasswordCannotBeEmpty
	}

	credentials := Credentials(username, password)

	account.Email = strings.ToLower(email)

	// Set username (account name must always be lowercase in database)
	account.Username = strings.ToLower(username)

	// Set account tier
	account.Tier = tier

	// Generate more secure Bcrypt hash for web-based logins
	// This is for supporting Gophercraft itself and has nothing to do with supporting legacy protocols
	var err error
	account.Identity_Bcrypt, err = bcrypt.GenerateFromPassword(
		credentials,
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// Generate SRP identity hashes (BNet SRP v1)
	account.Identity_SHA_1 = hashutil.H(sha1.New, credentials)
	account.Identity_SHA_256 = hashutil.H(sha256.New, credentials)
	account.Identity_SHA_512 = hashutil.H(sha512.New, credentials)

	// Now we need to generate login information for BNet SRP6a (version 2)
	// first create an "srp username" i.e. toupper name using SHA-256/512 (512 is used in Gophercraft, 256 in Trinity/Cypher core)
	// TrinityCore appears to require uppername hashes here specifically (unknown)
	srp_username := strings.ToUpper(
		hex.EncodeToString(
			hashutil.H(sha512.New, []byte(strings.ToUpper(username))),
		),
	)

	// Todo: check how client interacts with password case here
	bnet_v2_credentials := srp_username + ":" + password

	// Read random bytes to sprinkle some salt into HMAC
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt[:]); err != nil {
		panic(err)
	}

	// Use PBKDF2 HMAC to generate salted password hash
	account.Identity_PBKDF2_SHA_512 = pbkdf2.Key([]byte(bnet_v2_credentials), salt, 15000, sha512.Size, sha512.New)
	account.Identity_PBKDF2_Salt = salt

	// voila!

	return nil
}
