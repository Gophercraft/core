package alphalist

import "github.com/superp00t/etc"

type Realm struct {
	Name            string
	RedirectAddress string
	Players         uint32
}

type List struct {
	Realms []Realm
}

func (l *List) Encode(wr *etc.Buffer) error {
	if len(l.Realms) > 255 {
		return ErrTooManyRealms
	}

	wr.Write([]byte{uint8(len(l.Realms))})

	for _, realm := range l.Realms {
		wr.WriteCString(realm.Name)
		wr.WriteCString(realm.RedirectAddress)
		wr.WriteUint32(realm.Players)
	}

	return nil
}
