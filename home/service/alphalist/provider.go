package alphalist

type ServiceProvider interface {
	ListRealms() ([]Realm, error)
}
