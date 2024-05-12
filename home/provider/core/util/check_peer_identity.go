package util

import (
	"context"

	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/home/protocol"
	"github.com/Gophercraft/log"
	"github.com/Gophercraft/phylactery/database"
	"github.com/Gophercraft/phylactery/database/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckPeerIdentity(db *database.Container, realm_ID uint64, peer_ctx context.Context) error {
	peer_fingerprint, err := protocol.GetPeerFingerprint(peer_ctx)
	if err != nil {
		return err
	}

	var realm models.EnlistedRealm
	ok, err := db.Table("EnlistedRealm").Where(query.Eq("ID", realm_ID)).Get(&realm)
	if err != nil {
		log.Warn(err)
		return status.Errorf(codes.NotFound, "error querying database")
	}
	if !ok {
		return status.Errorf(codes.NotFound, "no realm exists corresponding to %d. You can solve this by registering via the web portal", realm_ID)
	}

	if !protocol.FingerprintsEqual(peer_fingerprint, realm.Fingerprint) {
		return status.Errorf(codes.PermissionDenied, "realm %d exists, but this realm's fingerprint does not match the one on record. you can solve this by enlisting a new realm with the correct fingerprint", realm_ID)
	}

	return nil
}
