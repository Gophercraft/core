package packet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

// In the legacy protocol, a WorldPacket's length cannot exceed this size
const MaxLength = 32766

// Connection provides a buffered communication pipe with which to send and receive packets through the game protocol.
// The purpose of this struct is to massively simplify the process of sending and receiving packets across different protocol revisions.
// In the legacy protocol, the size and WorldType are encrypted while the packet body is transmitted in plaintext.
// Starting in 8.2.5, the WorldType and packet buffer are encrypted in transit with AES-128-GCM, prefixed with an unencrypted size header.
type Connection struct {
	Build      vsn.Build
	Conn       net.Conn
	Reader     *bufio.Reader
	WorldTypes *WorldTypeDescriptor
	SessionKey []byte
	Cipher     crypto.Cipher
	Server     bool
	write      sync.Mutex
	closed     bool

	// params
	opcodeReadSize  uint8
	opcodeWriteSize uint8
	sizePrefixSize  uint8
	extraDataSize   uint8
	flag            uint8
}

const (
	flagCata = 1 << iota
	flagCryptHeader
)

func NewConnection(version vsn.Build, c net.Conn, server bool) (*Connection, error) {
	cr := new(Connection)
	cr.Build = version
	cr.Conn = c
	cr.Reader = bufio.NewReaderSize(c, 65535)
	cr.Server = server
	cr.Cipher = crypto.DummyCipher{}
	var err error

	err = vsn.QueryDescriptors(version, WorldTypeDescriptors, &cr.WorldTypes)
	if err != nil {
		return nil, err
	}

	if err = cr.setupParams(); err != nil {
		return nil, err
	}

	return cr, nil
}

// set up the basic parameters of the connection based on what build this connection is targeting
func (c *Connection) setupParams() error {
	// Size of opcode
	// c.wCodeSize = 2
	c.sizePrefixSize = 2
	c.extraDataSize = 0
	c.flag = 0

	// set wCodeSize and sizePrefixSize - variable between versions
	switch {
	//
	case classicRange.Contains(c.Build):
		if c.Server {
			c.opcodeReadSize = 4
			c.opcodeWriteSize = 2
		} else {
			c.opcodeReadSize = 2
			c.opcodeWriteSize = 4
		}
		c.sizePrefixSize = 2
		c.flag |= flagCryptHeader
	case cataRange.Contains(c.Build):
		c.flag |= flagCata
		// cata has a very strange variable-length encoding of large sizes
		// if (2 + p.Len()) > 0x7FFF {
		// 	c.sizePrefixSize = 3
		// } else {
		// 	c.sizePrefixSize = 2
		// }
		c.sizePrefixSize = 2

		if c.Server {
			c.opcodeReadSize = 4
			c.opcodeWriteSize = 2
		} else {
			c.opcodeReadSize = 2
			c.opcodeWriteSize = 4
		}
		c.flag |= flagCryptHeader
	case mopRange.Contains(c.Build):
		// mop
		c.opcodeReadSize = 2
		c.opcodeWriteSize = 2
		c.sizePrefixSize = 2
		c.flag |= flagCryptHeader
	case wodRange.Contains(c.Build):
		// wod
		switch c.Cipher.(type) {
		case crypto.DummyCipher:
			c.sizePrefixSize = 2
		default:
			c.sizePrefixSize = 4
		}
		c.opcodeReadSize = 2
		c.opcodeWriteSize = 2
		c.flag |= flagCryptHeader
	case legionBfARange.Contains(c.Build):
		// legion - early BfA
		c.opcodeReadSize = 2
		c.opcodeWriteSize = 2
		c.sizePrefixSize = 4
		c.flag |= flagCryptHeader
	case aesCryptRange.Contains(c.Build):
		// watch for future changes - here is the default behavior for all new protocols
		c.opcodeReadSize = 2
		c.opcodeWriteSize = 2
		c.sizePrefixSize = 4
	default:
		return fmt.Errorf("packet: unknown msg code size for build %s", c.Build)
	}

	// set extraDataSize (alpha quirk, unknown what this is for)
	switch {
	case vsn.Range(0, 3368).Contains(c.Build):
		c.extraDataSize = 2
	}

	return nil
}

func (c *Connection) ident(serv bool) []byte {
	if !c.Build.AddedIn(20886) {
		return nil
	}

	ident := []byte("WORLD OF WARCRAFT CONNECTION - ")

	if serv {
		ident = append(ident, []byte("SERVER TO CLIENT")...)
	} else {
		ident = append(ident, []byte("CLIENT TO SERVER")...)
	}

	if c.Build.AddedIn(31478) {
		ident = append(ident, []byte(" - V2")...)
	}

	if c.Build.AddedIn(22248) {
		ident = append(ident, '\n')
	}

	return ident
}

func (c *Connection) ConfirmProtocol() error {
	if !c.Build.AddedIn(20886) {
		return nil
	}

	if c.Server {
		if _, err := c.Conn.Write(c.ident(true)); err != nil {
			return err
		}
	} else {
		expectedServerIdent := c.ident(true)
		actualServerIdent := make([]byte, len(expectedServerIdent))
		if _, err := io.ReadFull(c.Reader, actualServerIdent[:]); err != nil {
			return err
		}
		if !bytes.Equal(expectedServerIdent, actualServerIdent) {
			return fmt.Errorf("packet: server sent us an incorrect ident")
		}
	}

	if c.Server {
		expectedClientIdent := c.ident(false)
		actualClientIdent := make([]byte, len(expectedClientIdent))
		if _, err := io.ReadFull(c.Reader, actualClientIdent[:]); err != nil {
			return err
		}
		if !bytes.Equal(expectedClientIdent, actualClientIdent) {
			return fmt.Errorf("packet: client sent us an incorrect ident")
		}
	} else {
		if _, err := c.Conn.Write(c.ident(false)); err != nil {
			return err
		}
	}

	return nil
}

func (c *Connection) SetNagle(ok bool) {
	c.Conn.(*net.TCPConn).SetNoDelay(!ok)
}

func (c *Connection) InitEncryption(sessionKey []byte) error {
	var err error
	c.Cipher, err = crypto.NewCipher(c.Build, sessionKey, c.Server)
	if err != nil {
		panic(err)
	}
	return nil
}

// The header that precedes a packet, always unencrypted.
type Header struct {
	Size uint32
	Tag  [12]byte
}

var (
	classicRange   = vsn.Range(0, 12340)
	cataRange      = vsn.Range(13164, 15595)
	mopRange       = vsn.Range(15851, 18414)
	wodRange       = vsn.Range(19027, 21742)
	legionBfARange = vsn.Range(22248, 30706)
	aesCryptRange  = vsn.Range(31478, vsn.Max)
)

func (c *Connection) Send(p *WorldPacket) error {
	wCode, err := c.WorldTypes.Code(p.Type)
	if err != nil {
		return err
	}

	dataPrefixSize := int(c.opcodeWriteSize) + int(c.extraDataSize)

	// Copy data to be encrypted
	data := make([]byte, p.Len()+dataPrefixSize)
	copy(data[dataPrefixSize:], p.Bytes())

	// set world code
	switch c.opcodeWriteSize {
	case 2:
		binary.LittleEndian.PutUint16(data[0:2], uint16(wCode))
	case 4:
		binary.LittleEndian.PutUint32(data[0:4], uint32(wCode))
	default:
		panic(c.opcodeWriteSize)
	}

	header := Header{
		Size: uint32(len(data)),
	}

	var headerData []byte
	//cata requires special behavior to encode the size header server-side (wtf were they smoking)
	if c.Server && c.flag&flagCata != 0 {
		sz := uint32(len(data))
		if sz > 0x7FFF {
			headerData = append(headerData, uint8(0x80|0xFF&(sz>>16)))
		}

		var normalData [2]byte
		binary.LittleEndian.PutUint16(normalData[0:2], uint16(sz))
		headerData = append(headerData, normalData[:]...)
	} else {
		// "normal" behavior
		headerData = make([]byte, c.sizePrefixSize)
		switch c.sizePrefixSize {
		case 2:
			binary.BigEndian.PutUint16(headerData[:], uint16(len(data)))
		case 4:
			binary.LittleEndian.PutUint32(headerData[:], uint32(len(data)))
		default:
			panic(c.sizePrefixSize)
		}
	}

	if c.flag&flagCryptHeader != 0 {
		if err := c.Cipher.Encrypt(headerData, nil); err != nil {
			return err
		}

		if err := c.Cipher.Encrypt(data[0:int(c.opcodeWriteSize)], nil); err != nil {
			return err
		}
	} else {
		if err := c.Cipher.Encrypt(data, header.Tag[:]); err != nil {
			return err
		}
		// AES encryption is enabled, pack tag after header
		// If handshake is in process, this is 12 zero bytes
		if aesCryptRange.Contains(c.Build) {
			headerData = append(headerData, header.Tag[:]...)
		}
	}

	// fmt.Println(spew.Sdump(headerData))
	// fmt.Println(spew.Sdump(data))
	// log.Dump(p.Type.String()+" headerData", headerData)
	// log.Dump(p.Type.String()+" data", data)

	if _, err := c.Conn.Write(headerData); err != nil {
		return err
	}

	if _, err := c.Conn.Write(data); err != nil {
		return err
	}

	// if _, err := c.Conn.Write(append(headerData, data...)); err != nil {
	// 	return err
	// }

	return nil
}

func (c *Connection) Recv() (*WorldPacket, error) {
	var headerData [16]byte
	// var wCodeSize int
	// var sizePrefixSize int
	// var decryptHeader bool

	if c.flag&flagCata != 0 {
		// support weird cata behavior
		if !c.Server {
			if _, err := c.Reader.Read(headerData[0:1]); err != nil {
				return nil, err
			}
			if err := c.Cipher.Decrypt(headerData[0:1], nil); err != nil {
				return nil, err
			}
			var hiMask, size uint32
			if headerData[0]&0x80 != 0 {
				hiMask = ((uint32(headerData[0]) & ^uint32(0x80)) << 16)
				if _, err := c.Reader.Read(headerData[1:3]); err != nil {
					return nil, err
				}
				if err := c.Cipher.Decrypt(headerData[1:3], nil); err != nil {
					return nil, err
				}
				size = uint32(binary.LittleEndian.Uint16(headerData[1:3]))
				size |= hiMask
			} else {
				if _, err := c.Reader.Read(headerData[1:2]); err != nil {
					return nil, err
				}
				if err := c.Cipher.Decrypt(headerData[1:2], nil); err != nil {
					return nil, err
				}
				size = uint32(binary.LittleEndian.Uint16(headerData[0:2]))
			}
			content := make([]byte, size)
			if _, err := c.Reader.Read(content[:]); err != nil {
				return nil, err
			}
			if err := c.Cipher.Decrypt(content[0:2], nil); err != nil {
				return nil, err
			}
			wType, err := c.WorldTypes.LookupCode(uint32(binary.LittleEndian.Uint16(content[0:2])))
			if err != nil {
				return nil, err
			}
			return &WorldPacket{
				Type:   wType,
				Buffer: etc.OfBytes(content),
			}, nil
		}
		// c.send = 4
		// sizePrefixSize = 2
		// decryptHeader = true
	}

	var size uint32

	if c.flag&flagCryptHeader != 0 {
		if _, err := c.Reader.Read(headerData[0:c.sizePrefixSize]); err != nil {
			return nil, err
		}
		if err := c.Cipher.Decrypt(headerData[0:c.sizePrefixSize], nil); err != nil {
			return nil, err
		}
		switch c.sizePrefixSize {
		case 2:
			size = uint32(binary.BigEndian.Uint16(headerData[0:2]))
		case 4:
			size = uint32(binary.LittleEndian.Uint32(headerData[0:4]))
		default:
			panic(c.sizePrefixSize)
		}
	} else {
		// Use AES
		if _, err := c.Reader.Read(headerData[:]); err != nil {
			return nil, err
		}

		size = binary.LittleEndian.Uint32(headerData[0:4])
	}

	content := make([]byte, size)
	if _, err := c.Reader.Read(content); err != nil {
		return nil, err
	}

	decryptOffset := int(c.opcodeReadSize)
	if aesCryptRange.Contains(c.Build) {
		decryptOffset = len(content)
	}

	if err := c.Cipher.Decrypt(content[:decryptOffset], headerData[4:]); err != nil {
		return nil, err
	}

	var wCode uint32
	switch c.opcodeReadSize {
	case 2:
		wCode = uint32(binary.LittleEndian.Uint16(content[0:2]))
	case 4:
		wCode = uint32(binary.LittleEndian.Uint32(content[0:4]))
	default:
		panic(c.opcodeReadSize)
	}

	wType, err := c.WorldTypes.LookupCode(wCode)
	if err != nil {
		return nil, err
	}

	content = content[c.opcodeReadSize:]
	return &WorldPacket{
		Type:   wType,
		Buffer: etc.OfBytes(content),
	}, nil
}

func (c *Connection) Close() error {
	return c.Conn.Close()
}
