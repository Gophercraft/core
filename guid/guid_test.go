package guid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

func TestGUID(t *testing.T) {
	var guid1 uint64 = 0x0000BADC0DECBBBB
	g := Classic(guid1)

	if g.HighType() != Player {
		t.Fatal("mismatch", g.HighType())
	}

	if g.Classic() != guid1 {
		fmt.Printf("0x%16X\n", guid1)
		panic("data loss in encoded guid")
	}

	fmt.Println(g.HiClassic())

	fmt.Println(g)

	buffer := etc.NewBuffer()
	g.EncodePacked(5875, buffer)
	log.Dump("buffer.Bytes()", buffer.Bytes())

	gd, err := DecodePacked(5875, buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println(g, gd)

	if gd != g {
		t.Fatal("GUID data loss detected", g, gd)
	}

	guid2 := uint64(0x1FC0000000000002)
	fmt.Println(Classic(guid2))

	str := "Player-6-0034445"

	guid3, err := FromString(str)
	if err != nil {
		panic(err)
	}

	if guid3.HighType() != Player {
		panic(guid3.HighType())
	}

	if guid3.Counter() != 0x0034445 {
		panic(guid3.Counter())
	}

	if guid3.RealmID() != 6 {
		panic(guid3.RealmID())
	}

	guid4 := uint64(0x70000000027C21C)

	fmt.Println("guid4", Classic(guid4))

	filled := uint64(0xCCCCBBBBAAAAAAAA)

	bmask, bits := encodeMasked64(filled)
	orig := decodeMasked64(bmask, bytes.NewReader(bits))
	if orig != filled {
		panic(fmt.Errorf("0x%16X encoded as %8b %s: decoded as 0x%16X", filled, bmask, hex.EncodeToString(bits), orig))
	}
}

// func testEncoding(t *testing.T, guid1 GUID) {
// 	fmt.Println("Before encoding: ", guid1)

// 	bts := guid1.EncodePacked()
// 	fmt.Println("Encoded as bytes:", bts)
// 	e2 := etc.MkBuffer(bts)
// 	// append some data, so we can be sure that the decoder works
// 	// even when the packed GUID data is followed immediately by other data
// 	e2.WriteUint32(0xFFFFFFFF)
// 	g2 := DecodePacked(e2)
// 	fmt.Println("Decoded from bytes:", g2)

// 	safeData := e2.ReadUint32()
// 	if safeData != 0xFFFFFFFF {
// 		t.Fatal("mismatch")
// 	}
// }

// func TestEncoding(t *testing.T) {
// 	for _, g := range []GUID{
// 		0x0000000000521BC0,
// 		0xDEADBEEF1337BADC,
// 	} {
// 		testEncoding(t, g)
// 	}
// }
