package web

import (
	"fmt"
	"strconv"
	"time"

	home_models "github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol/pb/auth"
	api_models "github.com/Gophercraft/core/home/service/web/models"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database/query"
)

func (provider *service_provider) GetRealmStatusList(user_info *api_models.UserInfo) (list *api_models.RealmStatusList, err error) {
	var realms []home_models.Realm
	if err = provider.home_db.Table("Realm").Where(query.Lte("RequiredTier", auth.AccountTier_NORMAL)).OrderBy("ID", false).Find(&realms); err != nil {
		log.Warn(err)
		err = fmt.Errorf("home/provider/web: failed to get realms from database")
		return
	}

	list = new(api_models.RealmStatusList)
	list.Realms = make([]api_models.RealmStatus, len(realms))

	now := time.Now()

	for i := range realms {
		realm := &realms[i]
		status := &list.Realms[i]

		status.ID = strconv.FormatUint(realm.ID, 10)
		status.Name = realm.Name
		status.Description = realm.Description
		status.Build = realm.Build.DBD()
		status.Expansion = int32(realm.Build.BuildInfo().MajorVersion)
		status.Online = now.Sub(realm.LastUpdated) > home_models.RealmOnlineDelayMeansOffline
	}

	return
}
