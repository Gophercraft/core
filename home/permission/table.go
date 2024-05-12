package permission

import "github.com/Gophercraft/core/home/protocol/pb/auth"

type List []Permission

type Table struct {
	tier_lists map[auth.AccountTier]List
}

func MakeTable(table_config *config.PermissionsTable) (table *Table)

func (table *Table) HasPermission(tier auth.AccountTier, permission Permission) bool {

}
