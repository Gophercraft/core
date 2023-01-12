package wizard

import (
	"github.com/Gophercraft/core/home/config"
)

// List all home configs in the chosen Gophercraft root directory.
// On a typical deployment, it will return []string{"Home"}, nil
func (c *Configurator) ListHomeConfigs() ([]string, error) {
	return config.ListHomeConfigs(c.Dir)
}
