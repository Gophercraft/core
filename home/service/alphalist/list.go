package alphalist

import (
	"fmt"

	"github.com/Gophercraft/core/serialization"
)

type Realm struct {
	Name            string
	RedirectAddress string
	Players         uint32
}

func EncodeList(list []Realm) (data []byte, err error) {
	if len(list) > 255 {
		return nil, fmt.Errorf("alphalist: cannot encode more than 255 realms due to hard uint8 limit")
	}

	var buffer serialization.Buffer
	buffer.WriteUint8(uint8(len(list)))

	for _, realm := range list {
		buffer.WriteCString(realm.Name)
		buffer.WriteCString(realm.RedirectAddress)
		buffer.WriteUint32(realm.Players)
	}

	data = buffer.Bytes()
	return
}
