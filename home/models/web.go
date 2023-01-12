package models

type WebSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WebSignInResponse struct {
	Error    string `json:"error"`
	WebToken string `json:"webToken"`
}

type WebEnlistRealmRequest struct {
	WebToken    string `json:"webToken"`
	Name        string `json:"string"`
	Fingerprint string `json:"fingerprint"`
}

type WebEnlistRealmResponse struct {
	RealmID uint64 `json:"realmID"`
}

type WebRegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	CaptchaID       string `json:"captchaID"`
	CaptchaSolution string `json:"captchaSolution"`
}

type WebRegisterResponse struct {
	Error        string `json:"error"`
	ResetCaptcha bool   `json:"resetCaptcha"`
	WebToken     string `json:"webToken,omitempty"`
}
