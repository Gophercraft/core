# 0.7.1 Release (WIP)

- Massive overhaul to both Home and World servers.

- Home now has optional services - only the servers included in your home.txt file are used.

# 0.5.1 Release

New features:
- 3D math now lives in package tempest

Put all 3D rotation/position related in tempest.

Sessions now can control Processes. Processes are simply functions that accept a shutdown channel. Processes should always read the shutdown channel, and cancel when the signal is received.

Packets of any type can be encoded if they have the method Encode.

```go

type Encodable struct {
  Encode(build version.Build, out *message.Packet) error
}

type Decodable interface {
  Decode(build version.Build, in *message.Packet) error
}

type Codec interface {
  Encodable
  Decodable
}

```

