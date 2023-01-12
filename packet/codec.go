package packet

import "github.com/Gophercraft/core/vsn"

type Encodable interface {
	Encode(build vsn.Build, p *WorldPacket) error
}

type Decodable interface {
	Decode(build vsn.Build, p *WorldPacket) error
}

type Codec interface {
	Encodable
	Decodable
}
