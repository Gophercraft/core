package wizard

type Configurator struct {
	Dir string
}

func NewConfigurator(d string) *Configurator {
	return &Configurator{d}
}
