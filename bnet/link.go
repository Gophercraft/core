package bnet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
	"sync/atomic"

	"github.com/Gophercraft/core/bnet/pb/bgs/protocol"
	"google.golang.org/protobuf/proto"
)

type link struct {
	tls_connection         net.Conn
	close_after_next_write uint32
}

// signal to other routines using the link that it is shutting down
func (link *link) request_close_after_next_write() {
	atomic.StoreUint32(&link.close_after_next_write, 1)
}

// are we supposed to stop using the link?
func (link *link) close_requested() bool {
	return atomic.LoadUint32(&link.close_after_next_write) == 1
}

// read a header-prefixed  message from a link
func (link *link) read_message() (header *protocol.Header, message []byte, err error) {
	var size_prefix [2]byte
	if _, err = io.ReadFull(link.tls_connection, size_prefix[:]); err != nil {
		err = fmt.Errorf("bnet: failed to read header size prefix: %w", err)
		return
	}
	// read size in network order
	size := int(binary.BigEndian.Uint16(size_prefix[:]))

	// read header
	pb_header := make([]byte, size)
	if _, err = io.ReadFull(link.tls_connection, pb_header[:]); err != nil {
		return
	}
	// unmarshal header
	header = new(protocol.Header)
	if err = proto.Unmarshal(pb_header, header); err != nil {
		return
	}

	if header.Size != nil {
		message_size := header.GetSize()
		message = make([]byte, message_size)
		if _, err = io.ReadFull(link.tls_connection, message); err != nil {
			return
		}
	}

	return
}

func (link *link) write_message(header *protocol.Header, message []byte) (err error) {
	var (
		buffer      bytes.Buffer
		pb_header   []byte
		size_prefix [2]byte
	)

	// Encode header
	pb_header, err = proto.Marshal(header)
	if err != nil {
		return
	}
	if len(pb_header) > math.MaxUint16 {
		err = fmt.Errorf("bnet: encoded header size larger than maximum uint16")
		return
	}
	// Encode header size prefix
	binary.BigEndian.PutUint16(size_prefix[:], uint16(len(pb_header)))

	// Send data all at once (to allow concurrent writes, and avoid any misunderstanding by the client)
	buffer.Write(size_prefix[:])
	buffer.Write(pb_header)
	buffer.Write(message)

	_, err = link.tls_connection.Write(buffer.Bytes())

	if link.close_requested() {
		link.tls_connection.Close()
		return
	}

	return
}

func (link *link) Close() (err error) {
	err = link.tls_connection.Close()
	return
}
