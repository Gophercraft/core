package warden

import "fmt"

func newServerRequest(cmd Command) (Request, error) {
	switch cmd {
	case CServerModuleUse:
		return new(ServerModuleUse), nil
	case CServerModuleTransfer:
		return new(ServerModuleTransfer), nil
	case CServerRequestCheatChecks:
		return new(ServerRequestCheatChecks), nil
	case CServerModuleInitialize:
		return new(ServerModuleInitialize), nil
	case CServerRequestMemoryChecks:
		return new(ServerRequestMemoryChecks), nil
	case CServerHashRequest:
		return new(ServerHashRequest), nil
	default:
		return nil, fmt.Errorf("packet/warden: no server request found for command %d", cmd)
	}
}

func newClientResult(cmd Command) (ClientResult, error) {
	switch cmd {
	case CClientModuleMissing, CClientModuleFailed, CClientModuleOK:
		return emptyRequest(cmd), nil
	case CClientHashResult:
		return new(ClientHashResult), nil
	case CClientCheatChecksResult:
		return new(ClientCheatChecksResult), nil
	default:
		return nil, fmt.Errorf("packet/warden: unknown client command type %d", cmd)
	}
}
