package ex

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/Gophercraft/log"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	protoTail  = []byte(".proto")
	markerByte = 0x0a
	ptags      = []byte("\x12\x1a\x22\x2a\x32\x3a\x42\x4a\x50\x58\x62")
)

type Extractor struct {
	Dir     string
	Program []byte
	Offset  int
}

func (ex *Extractor) Open(filename string) (err error) {
	ex.Program, err = ioutil.ReadFile(filename)
	return
}

// Finds the beginning of a new descriptor
func (ex *Extractor) nextOffset() (int, error) {
	for {
		workingSlice := ex.Program[ex.Offset:]

		relExtensionOffset := bytes.Index(workingSlice, protoTail)

		if relExtensionOffset == -1 {
			return 0, io.EOF
		}

		extensionOffset := ex.Offset + relExtensionOffset

		endOfProtoSuffix := extensionOffset + 6

		if !bytes.Equal(ex.Program[extensionOffset:endOfProtoSuffix], protoTail) {
			return 0, fmt.Errorf("bytes mismatch")
		}

		if len(ex.Program) >= extensionOffset+5 {
			if bytes.Equal(ex.Program[endOfProtoSuffix:endOfProtoSuffix+5], []byte("devel")) {
				endOfProtoSuffix += 5
			}
		}

		segmentOffset := ex.Offset - 1024
		if segmentOffset < 0 {
			segmentOffset = 0
		}

		searchSegment := ex.Program[segmentOffset:endOfProtoSuffix]

		relMarkerIndex := bytes.LastIndexByte(searchSegment, byte(markerByte))

		if relMarkerIndex == -1 {
			return 0, fmt.Errorf("ex: found proto but not marker")
		}

		markerIndex := relMarkerIndex + segmentOffset

		if ex.Program[markerIndex] != byte(markerByte) {
			return 0, fmt.Errorf("ex: bad marker discovery")
		}

		sizeVarint, pos := binary.Uvarint(ex.Program[markerIndex+1:])
		if sizeVarint == uint64(endOfProtoSuffix-(markerIndex+1+pos)) {
			top := markerIndex + 512

			if top > len(ex.Program) {
				top = len(ex.Program)
			}

			log.Dump("program section", ex.Program[markerIndex:top])
			log.Println("found it lol")
			ex.Offset = endOfProtoSuffix
			return markerIndex, nil
		} else {
			log.Println("size varint", sizeVarint, "size byte", ex.Program[markerIndex+1], int64(endOfProtoSuffix-(markerIndex+1+pos)))
			ex.Offset = endOfProtoSuffix
			continue
		}
	}
}

func (ex *Extractor) ExtractNext() error {

	for {
		offset, err := ex.nextOffset()
		if err != nil {
			return err
		}

		for k := 1024; k >= 0; k-- {
			var desc descriptorpb.FileDescriptorProto

			err = proto.Unmarshal(ex.Program[offset:offset+k], &desc)
			if err == nil {
				log.Println("Protofile get!", desc.GetName())
				// log.Dump("MessageType", desc.MessageType)
				log.Dump("descriptor", desc)
				// log.Fatal("yo")

				return ex.DumpDescriptor(&desc)
			}
			// } else {
			// 	log.Warn(err)
			// }
		}

		return fmt.Errorf("no way to handle this")
	}

	// chunk, err := ex.nextChunk()
	// if err != nil {
	// 	return err
	// }

	// log.Dump("cnk", chunk)

	return nil
}

func Extract(file string, folder string) error {
	// spew.Config.DisableMethods = true

	ex := &Extractor{}
	ex.Dir = folder
	if err := ex.Open(file); err != nil {
		return err
	}

	for {
		if err := ex.ExtractNext(); err != nil {
			return err
		}
		// log.Fatal("shet")
	}

	return nil
}

func (ex *Extractor) DumpDescriptor(desc *descriptorpb.FileDescriptorProto) error {
	return ex.NewDumper(desc).Dump()
}
