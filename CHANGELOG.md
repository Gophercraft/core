# 0.5.1 Release

New features:
- 3D math now lives in package tempest

Put all 3D rotation/position related in tempest.

Sessions now can control Processes. Processes are simply functions that accept a shutdown channel. Processes should always read the shutdown channel, and cancel when the signal is received.

Packets of any type can be encoded if they have the method Encode.

```go

type Encodable struct {
  Encode(build vsn.Build, out *packet.WorldPacket) error
}

type Decodable interface {
  Decode(build vsn.Build, in *packet.WorldPacket) error
}

type Codec interface {
  Encodable
  Decodable
}

```

