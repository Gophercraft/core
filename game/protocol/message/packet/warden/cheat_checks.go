package warden

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"time"

	wardencrypto "github.com/Gophercraft/core/crypto/warden"
	"github.com/Gophercraft/core/version"
	"github.com/superp00t/etc"
)

type CheckType uint8

const (
	CheckSkip      = 0   // SKIP
	CheckTiming    = 87  // nyi
	CheckDriver    = 113 // uint Seed + byte[20] SHA1 + byte driverNameIndex (check to ensure driver isn't loaded)
	CheckProcess   = 126 // nyi
	CheckLuaEval   = 139 // evaluate arbitrary Lua check
	CheckPaQ       = 152 // get hash of MPQ file (to check it is not modified)
	CheckAllPages  = 178 // scans all pages for specified SHA1 hash
	CheckMZPEPages = 191 // scans only pages starts with MZ+PE headers for specified hash
	CheckModule    = 217 // check to make sure module isn't injected
	CheckMem       = 243 // retrieve specific memory
)

type CheatCheck struct {
	Type    CheckType
	String  string
	Address uint64
	Size    uint32
	Data    []byte
}

type ServerRequestCheatChecks struct {
	Checks []CheatCheck
}

func (srcc *ServerRequestCheatChecks) Command() Command {
	return CServerRequestCheatChecks
}

func (srcc *ServerRequestCheatChecks) Decode(build version.Build, in *Reader) (err error) {
	var stringTable []string
	/// Read string table
	for {
		c, err := lByte(in)
		if err != nil {
			return err
		}
		if c == 0x00 {
			break
		}
		data := make([]byte, c)
		in.Read(data[:])
		stringTable = append(stringTable, string(data))
	}

	// We'll need these to parse objects
	xorByte := in.ClientKey[0]

	for {
		// Read type info
		c, err := lByte(in)
		if err != nil {
			return err
		}
		check := CheatCheck{}
		check.Type = CheckType(c ^ xorByte)

		switch check.Type {
		case CheckTiming:

		case CheckMem:
			lByte(in)
			address := uint32(0)
			lLe(in, &address)
			check.Address = uint64(address)
			lLe(in, &check.Size)
		case CheckAllPages, CheckMZPEPages:
			check.Data = make([]byte, 20)
			if _, err := in.Read(check.Data[:]); err != nil {
				return err
			}
			address := uint32(0)
			lLe(in, &address)
			check.Address = uint64(address)
			lLe(in, &check.Size)
		case CheckPaQ, CheckLuaEval:
			_index, err := lByte(in)
			if err != nil {
				return err
			}
			index := int(_index)

			if index >= len(stringTable) {
				return fmt.Errorf("packet/warden: index outside of string table %d", index)
			}
			check.String = stringTable[index]
		case CheckDriver:
			_index, err := lByte(in)
			if err != nil {
				return err
			}
			index := int(_index)

			if index >= len(stringTable) {
				return fmt.Errorf("packet/warden: index outside of string table %d", index)
			}
			check.String = stringTable[index]
		case CheckModule:
			// 1-way encryption here
			// we'd have to know good hashes to get an answer
			var seed [4]byte
			in.Read(seed[:])
			var digest [20]byte
			in.Read(digest[:])
			check.Data = append(seed[:], digest[:]...)
		default:
			return fmt.Errorf("packet/warden: tried to encoded unknown check %d", check.Type)
		}
	}

	xb, err := lByte(in)
	if err != nil {
		return err
	}
	if xb != xorByte {
		return fmt.Errorf("packet/warden: xor byte in packet different than selected xor byte")
	}
	return nil
}

func (srcc *ServerRequestCheatChecks) Encode(build version.Build, out *Writer) (err error) {
	for _, check := range srcc.Checks {
		switch check.Address {
		case CheckLuaEval, CheckPaQ, CheckDriver:
			out.Write([]byte{uint8(len(check.String))})
			out.Write([]byte(check.String))
		}
	}

	var index int
	xorByte := out.CryptoData.ClientKey[0]

	// terminate string table
	out.Write([]byte{0x00})
	for _, check := range srcc.Checks {
		// stupid "obfuscation"
		out.Write([]byte{xorByte ^ uint8(check.Type)})
		switch check.Type {
		case CheckTiming:
		case CheckMem:
			out.Write([]byte{0x00})
			sLe(out, uint32(check.Address))
			sLe(out, uint8(check.Size))
		case CheckAllPages, CheckMZPEPages:
			out.Write(check.Data[:])
			sLe(out, uint32(check.Address))
			sLe(out, uint32(check.Size))
		case CheckPaQ, CheckLuaEval:
			out.Write([]byte{uint8(index)})
			index++
		case CheckDriver:
			out.Write(check.Data)
			out.Write([]byte{uint8(index)})
			index++
		case CheckModule:
			var seed [4]byte
			rand.Read(seed[:])
			out.Write(seed[:])
			sh := hmac.New(sha1.New, seed[:])
			digest := sh.Sum([]byte(check.String))
			out.Write(digest[:])
		default:
			err = fmt.Errorf("packet/warden: unknown check type %s", check.Type)
			return
		}
	}

	out.Write([]byte{xorByte})

	return
}

type ClientCheckResult struct {
	Result         uint8
	Response       []byte
	NewClientTicks uint32
	Time           time.Time
}

type ClientCheatChecksResult struct {
	CurrentChecks *ServerRequestCheatChecks
	Length        uint16
	Checksum      uint32
	CheckResults  []ClientCheckResult
}

func (cccr *ClientCheatChecksResult) Command() Command {
	return CClientCheatChecksResult
}

const PageCheckOK = 0xE9

func (cccr *ClientCheatChecksResult) Decode(build version.Build, in *Reader) (err error) {
	if cccr.CurrentChecks == nil {
		return ErrNeedCurrentChecks
	}

	var length uint16
	var checksum uint32
	lLe(in, &length)
	lLe(in, &checksum)

	data := make([]byte, length)
	if _, err := in.Read(data[:]); err != nil {
		return err
	}

	if checksum != wardencrypto.Checksum(data) {
		return ErrCheckBadChecksum
	}

	reader := bytes.NewReader(data)

	cccr.CheckResults = make([]ClientCheckResult, len(cccr.CurrentChecks.Checks))
	for i := 0; i < len(cccr.CurrentChecks.Checks); i++ {
		serverCheck := cccr.CurrentChecks.Checks[i]
		var err error
		var checkResult ClientCheckResult
		checkResult.Result, err = lByte(reader)
		if err != nil {
			return err
		}

		switch serverCheck.Type {
		case CheckTiming:
			// time between server check and client response
			lLe(in, checkResult.NewClientTicks)
			checkResult.Time = time.Now()
			// memory check
		case CheckMem:
			if checkResult.Result != 0x00 {
				cccr.CheckResults[i] = checkResult
				continue
			}

			checkResult.Response = make([]byte, serverCheck.Size)
			reader.Read(checkResult.Response[:])
		case CheckAllPages, CheckMZPEPages, CheckDriver, CheckModule:
		case CheckLuaEval:
			if checkResult.Result == 0 {
				luaStrSize, err := lByte(reader)
				if err != nil {
					return err
				}
				var luaStr = make([]byte, luaStrSize)
				reader.Read(luaStr[:])
				checkResult.Response = luaStr
			}
		case CheckPaQ:
			if checkResult.Result != 0 {
				cccr.CheckResults[i] = checkResult
				continue
			}
			var hsh [20]byte
			reader.Read(hsh[:])
			checkResult.Response = hsh[:]
		default:
			return fmt.Errorf("packet/warden: unknown check type from client %d", serverCheck.Type)
		}
		cccr.CheckResults[i] = checkResult
	}

	return nil
}

func (cccr *ClientCheatChecksResult) Encode(build version.Build, out *Writer) (err error) {
	if cccr.CurrentChecks == nil {
		return ErrNeedCurrentChecks
	}

	data := etc.NewBuffer()

	cccr.CheckResults = make([]ClientCheckResult, len(cccr.CurrentChecks.Checks))
	for i := 0; i < len(cccr.CurrentChecks.Checks); i++ {
		serverCheck := cccr.CurrentChecks.Checks[i]
		checkResult := cccr.CheckResults[i]

		data.WriteUint8(checkResult.Result)

		switch serverCheck.Type {
		case CheckTiming:
			// time between server check and client response
			sLe(out, checkResult.NewClientTicks)
			// memory check
		case CheckMem:
			if checkResult.Result != 0x00 {
				continue
			}

			out.Write(checkResult.Response)
		case CheckAllPages, CheckMZPEPages, CheckDriver, CheckModule:

		case CheckLuaEval:
			if checkResult.Result == 0x00 {
				data.WriteUint8(uint8(len(checkResult.Response)))
				data.Write(checkResult.Response)
			}
		case CheckPaQ:
			if checkResult.Result != 0 {
				continue
			}
			data.Write(checkResult.Response[:20])
		default:
			return fmt.Errorf("packet/warden: unknown check type from client %d", serverCheck.Type)
		}
		cccr.CheckResults[i] = checkResult
	}

	sLe(out, uint16(data.Len()))
	sLe(out, uint32(wardencrypto.Checksum(data.Bytes())))
	out.Write(data.Bytes())

	return nil
}
