package realm

import (
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func (s *Session) GetLoc(str string) string {
	var loc *models.LocString
	s.DB().Lookup(wdb.BucketKeyStringID, str, &loc)
	if loc == nil {
		return str
	}

	return loc.Text.GetLocalized(s.Locale)
}
