package models

type Error struct {
	Message string `json:"error_message,omitempty"`
}

type RegistrationChallenge struct {
	EmailRequired     bool   `json:"email_required"`
	MaxEmailLength    int32  `json:"max_email_length,omitempty"`
	MaxUsernameLength int32  `json:"max_username_length,omitempty"`
	MaxPasswordLength int32  `json:"max_password_length,omitempty"`
	CaptchaID         string `json:"captcha_id,omitempty"`
}

type LoginChallenge struct {
	CaptchaID string `json:"captcha_id,omitempty"`
}

type LoginRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	CaptchaID       string `json:"captcha_id"`
	CaptchaSolution string `json:"captcha_solution"`
}

type LoginResponse struct {
	WebToken              string `json:"web_token,omitempty"`
	TwoFactorAuthRequired bool   `json:"two_factor_auth_required"`
}

type LogoutResponse struct {
}

type EnlistRealmRequest struct {
	WebToken    string `json:"web_token"`
	Name        string `json:"string"`
	Fingerprint string `json:"fingerprint"`
}

type EnlistRealmResponse struct {
	RealmID uint64 `json:"realm_id"`
}

type RegistrationRequest struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	CaptchaID       string `json:"captcha_id"`
	CaptchaSolution string `json:"captcha_solution"`
}

type RegistrationResponse struct {
	EmailVerificationNeeded bool `json:"email_verification_needed"`
}

type CredentialStatus struct {
	CredentialIsValid bool   `json:"credential_is_valid"`
	WebTokenStatus    string `json:"web_token_status"`
}

type GameAccountStatus struct {
	Name               string `json:"name"`
	ID                 string `json:"id"`
	Active             bool   `json:"active"`
	Characters         uint32 `json:"characters"`
	Banned             bool   `json:"banned"`
	Suspended          bool   `json:"suspended"`
	SuspensionLiftDate string `json:"suspension_lift_date,omitempty"`
}

type AccountStatus struct {
	Username                               string              `json:"username"`
	ID                                     string              `json:"account_id"`
	Tier                                   string              `json:"account_tier"`
	Email                                  string              `json:"email"`
	CreationDate                           string              `json:"creation_date"`
	EmailVerified                          bool                `json:"email_verified"`
	PreferredTwoFactorAuthenticationMethod string              `json:"preferred_two_factor_authentication_method"`
	Authenticator                          bool                `json:"authenticator"`
	Locked                                 bool                `json:"locked"`
	Suspended                              bool                `json:"suspended"`
	SuspensionLiftDate                     string              `json:"suspension_lift_date,omitempty"`
	GameAccounts                           []GameAccountStatus `json:"game_accounts"`
}

type NewGameAccountRequest struct {
	Name string `json:"name"`
}

type NewGameAccountResponse struct {
	ID string `json:"id"`
}

type RenameGameAccountRequest struct {
	Name string `json:"name"`
}

type VersionInfo struct {
	CoreVersion string `json:"core_version"`
	Brand       string `json:"brand"`
	ProjectURL  string `json:"project_url"`
}

type RealmStatusList struct {
	Realms []RealmStatus `json:"realms"`
}

type RealmStatus struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Build       string `json:"build"`
	Expansion   int32  `json:"expansion"`
	Online      bool   `json:"online"`
}

type ServiceAddresses struct {
	Addresses map[string]string `json:"addresses"`
}

type CredentialAuthenticationRequest struct {
	TwoFactorAuthenticationMethod string `json:"two_factor_authentication_method"`
	AuthenticatorPassword         string `json:"authenticator_password"`
}

type CredentialAuthenticationResponse struct {
	Authenticated bool `json:"authenticated"`
}

type TwoFactorAuthenticationEnrollmentRequest struct {
	Secret   string `json:"totp_secret"`
	Password string `json:"totp_password"`
}

type TwoFactorAuthenticationEnrollmentResponse struct {
	Enrolled bool `json:"enrolled"`
}

type TwoFactorAuthMethods struct {
	Methods []string `json:"methods"`
}

type SentTwoFactorAuthenticationCodeEmailResponse struct {
	CensoredEmailAddress string `json:"censored_email_address,omitempty"`
	Sent                 bool   `json:"sent"`
}
