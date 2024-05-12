package models

import (
	"time"

	"github.com/Gophercraft/core/home/protocol/pb/auth"
)

// Account represents the authentication account
// This account tracks what privileges you are allowed, and what your password is.
type Account struct {
	ID        uint64 `database:"1:auto_increment,index,exclusive"`
	Tier      auth.AccountTier
	CreatedAt time.Time
	// Login temporarily disabled. Triggered by security threats. May require email verification to continue
	Locked bool
	// Permanently banned. Cannot log in through the web app.
	Banned bool
	// Temporarily suspended, account access is restored at UnsuspendAt
	Suspended bool
	// Determines if the email has been validated
	EmailVerified bool
	// Determines if TOTP is enabled for this account
	Authenticator bool
	// Marks the time
	EmailVerificationCode       string
	EmailVerificationCodeSentAt time.Time
	// When the suspension is going to be lifted
	UnsuspendAt time.Time
	// Language string
	Locale string
	// Grunt tag only (Win/Mac)
	OS string
	// Grunt tag only (x86/x64)
	Architecture string
	// BNet short string (Wn64, Mc64)
	Platform string
	// Account email address (may or may not be valid)
	Email    string
	Username string

	AuthenticatorSecret []byte
	// TODO: disable the generation of this field when user does not wish to provide Grunt compatiblity
	// Grunt Identity hash SHA1(TOUPPER(user) : TOUPPER(password))
	Identity_SHA_1 []byte
	// Same as Grunt but with SHA256 instead of SHA1, used for BNet REST login
	Identity_SHA_256        []byte
	Identity_SHA_512        []byte
	Identity_PBKDF2_SHA_512 []byte
	Identity_PBKDF2_Salt    []byte

	// Bcrypt hash - should be used when SRP IdentityHash is not required
	// Used for the Web API and wherever else credentials are transmitted in plaintext:
	Identity_Bcrypt []byte
	// Session Key
	SessionKey []byte
}
