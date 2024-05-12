package grunt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
	"strings"

	"github.com/Gophercraft/core/format/tag"
	"github.com/Gophercraft/core/version"
)

type LogonInfo struct {
	Program      Program
	Version      [3]byte
	Build        version.Build
	Architecture Architecture
	OS           OS
	Locale       Locale
	TimezoneBias int32
	IP           net.IP
	AccountName  string
}

// AuthLogonChallenge_Client is the first packet sent by a client
// while initiating a connection to an authserver.
type AuthLogonChallenge_Client struct {
	//
	Protocol uint8
	Info     LogonInfo
}

func ReadLogonInfo(reader io.Reader, logon *LogonInfo) (err error) {
	var (
		program_code tag.Tag
		os_code      tag.Tag
		arch_code    tag.Tag
		locale_code  tag.Tag
	)

	program_code, err = tag.Read(reader)
	if err != nil {
		return
	}
	logon.Program = Program(program_code)

	// Read version info
	if _, err = io.ReadFull(reader, logon.Version[:]); err != nil {
		return
	}
	var build_bytes [2]byte
	if _, err = io.ReadFull(reader, build_bytes[:]); err != nil {
		return
	}
	logon.Build = version.Build(binary.LittleEndian.Uint16(build_bytes[:]))

	// Read platform info
	arch_code, err = tag.Read(reader)
	if err != nil {
		return
	}
	logon.Architecture = Architecture(arch_code)
	os_code, err = tag.Read(reader)
	if err != nil {
		return
	}
	logon.OS = OS(os_code)

	// Read the localization of the client
	locale_code, err = tag.Read(reader)
	if err != nil {
		return
	}
	logon.Locale = Locale(locale_code)

	// Read timezone bias
	var timezone_bias_bytes [4]byte
	if _, err = io.ReadFull(reader, timezone_bias_bytes[:]); err != nil {
		return
	}
	logon.TimezoneBias = int32(binary.LittleEndian.Uint32(timezone_bias_bytes[:]))

	// Read client IP
	var client_IP_bytes [4]byte
	if _, err = io.ReadFull(reader, client_IP_bytes[:]); err != nil {
		return
	}
	logon.IP = net.IP(client_IP_bytes[:])

	// Read account name
	var account_name_size_byte [1]byte
	if _, err = io.ReadFull(reader, account_name_size_byte[:]); err != nil {
		return
	}
	account_name_size := int(account_name_size_byte[0])
	account_name_bytes := make([]byte, account_name_size)
	if _, err = io.ReadFull(reader, account_name_bytes); err != nil {
		return
	}
	logon.AccountName = string(account_name_bytes)
	return
}

func ReadAuthLogonChallenge_Client(reader io.Reader, challenge *AuthLogonChallenge_Client) (err error) {
	var header [3]byte
	if _, err = io.ReadFull(reader, header[:]); err != nil {
		return
	}

	// usually 8
	challenge.Protocol = header[0]

	message_size := int(binary.LittleEndian.Uint16(header[1:3]))
	message_bytes := make([]byte, message_size)
	if _, err = io.ReadFull(reader, message_bytes[:]); err != nil {
		return
	}

	message := bytes.NewBuffer(message_bytes)

	return ReadLogonInfo(message, &challenge.Info)
}

func WriteLogonInfo(writer io.Writer, logon *LogonInfo) (err error) {
	// write program code
	if err = tag.Write(writer, tag.Tag(logon.Program)); err != nil {
		return
	}

	// write version bytes
	if _, err = writer.Write(logon.Version[:]); err != nil {
		return
	}

	// check overflow
	if logon.Build > math.MaxUint16 {
		err = fmt.Errorf("version %d overflows max uint16", logon.Build)
		return
	}
	// write Build ID
	var build_ID_bytes [2]byte
	binary.LittleEndian.PutUint16(build_ID_bytes[:], uint16(logon.Build))
	if _, err = writer.Write(build_ID_bytes[:]); err != nil {
		return
	}

	// write CPU architecture
	if err = tag.Write(writer, tag.Tag(logon.Architecture)); err != nil {
		return
	}

	// write OS
	if err = tag.Write(writer, tag.Tag(logon.OS)); err != nil {
		return
	}

	// write client locale
	if err = tag.Write(writer, tag.Tag(logon.Locale)); err != nil {
		return
	}

	// write timezone bias
	var timezone_bias_bytes [4]byte
	binary.LittleEndian.PutUint32(timezone_bias_bytes[:], uint32(logon.TimezoneBias))
	if _, err = writer.Write(timezone_bias_bytes[:]); err != nil {
		return
	}

	// write client IP
	var client_IP_bytes [4]byte
	ipv4 := logon.IP
	if len(ipv4) != 4 {
		ipv4 = logon.IP.To4()
		if ipv4 == nil {
			err = fmt.Errorf("grunt: logon data IP must be IPv4")
			return
		}
	}
	copy(client_IP_bytes[:], ipv4)
	if _, err = writer.Write(client_IP_bytes[:]); err != nil {
		return
	}

	// write account name
	var account_name_bytes_size [1]byte
	account_name_bytes_size[0] = uint8(len(logon.AccountName))
	if _, err = writer.Write(account_name_bytes_size[:]); err != nil {
		return
	}
	if _, err = writer.Write([]byte(logon.AccountName)); err != nil {
		return
	}

	return
}

func WriteAuthLogonChallenge_Client(writer io.Writer, challenge *AuthLogonChallenge_Client) (err error) {
	var header [3]byte
	header[0] = byte(challenge.Protocol)

	envelope := new(bytes.Buffer)
	envelope.Grow(0x100)

	if err = WriteLogonInfo(envelope, &challenge.Info); err != nil {
		return
	}

	// seal envelope
	binary.LittleEndian.PutUint16(header[1:3], uint16(envelope.Len()))

	// write header
	if _, err = writer.Write(header[:]); err != nil {
		return
	}

	// write message envelope
	if _, err = writer.Write(envelope.Bytes()); err != nil {
		return
	}

	return
}

// LogonChallengePacket_C is a helper function to simplify the client library.
func MakeDefaultLogonChallenge_Client(build version.Build, username string) *AuthLogonChallenge_Client {
	var challenge AuthLogonChallenge_Client
	challenge.Protocol = 8
	challenge.Info = LogonInfo{
		Program:      WoW,
		Build:        build,
		Version:      Version(build),
		Architecture: X86,
		OS:           Windows,
		Locale:       Locale_enUS,
		TimezoneBias: 0,
		IP:           net.IPv4(127, 0, 0, 1),
		AccountName:  strings.ToUpper(username),
	}

	return &challenge
}

func Version(build version.Build) [3]byte {
	info := build.BuildInfo()
	if build.BuildInfo() == nil {
		return [3]byte{0, 0, 0}
	}

	return [3]byte{
		byte(info.MajorVersion),
		byte(info.MinorVersion),
		byte(info.BugfixVersion),
	}
}
