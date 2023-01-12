package home

import (
	"fmt"

	"github.com/Gophercraft/core/home/models"
)

func (h *Server) HandleEnlistRealm(acc *models.Account, enlistRequest *models.WebEnlistRealmRequest) (*models.WebEnlistRealmResponse, error) {
	if !h.CanEnlistRealm(acc) {
		return nil, fmt.Errorf("home: Account %d lacks the permission to enlist a realm", acc.ID)
	}

	var enlisted models.EnlistedRealm

	enlisted.Note = enlistRequest.Name
	enlisted.Owner = acc.ID
	enlisted.Fingerprint = enlistRequest.Fingerprint

	_, err := h.DB.Insert(&enlisted)
	if err != nil {
		return nil, fmt.Errorf("home: EnlistRealm failed")
	}

	var response models.WebEnlistRealmResponse
	response.RealmID = enlisted.ID

	return &response, nil
}
