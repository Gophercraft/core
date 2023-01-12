package wizard

type RegistrarParams struct {
	DBDriver, DBURL string
}

type ValidationCheck struct {
	Question string
	Object   string
}

// type Registrar interface {
// 	Init()
// 	NeedsAuth() bool

// }

type Registrar interface {
	Begin(wc *WorldConfigurator, location string) error
	Check() (*ValidationCheck, error)
	Confirm(check *ValidationCheck) error
	MustAuth() bool
	Auth(username, password string) error
	Enlist() error
}
