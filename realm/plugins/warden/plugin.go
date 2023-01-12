package warden

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm"
)

type Plugin struct {
	Module *Module
	init   bool
}

// Activated returns
func (c *Plugin) Activated() (bool, error) {
	// If Warden is disabled,
	if !c.init {
		return false, fmt.Errorf("warden plugin is disabled")
	}
	return true, nil
}

func (c *Plugin) Terminate() error {
	return nil
}

func (c *Plugin) Init(s *realm.Server, info *realm.PluginInfo) error {
	info.ID = "gophercraft.realm.plugins.warden"
	info.Name = "Gophercraft Warden Server"
	info.Description = "This plugin manages the Warden anticheat system. Custom checks can be added or removed via datapacks. As a server admin, you can choose what penalty is given for failing Warden scan, or disable penalties altogether."

	if s.Config.WardenEnabled {
		c.init = true
	}

	s.On(realm.SessionCreateEvent, nil, InitWardenSession)

	s.ProtocolHandlers.On(packet.CMSG_WARDEN_DATA, realm.Authed, HandleWardenData)

	return nil
}

{ 
	On Server.SessionCreate
	If {
		{ 
			@s.HasItem demo:gopher_book
		}
	}
	Do {
		{ @s.Teleport }
	}
}
