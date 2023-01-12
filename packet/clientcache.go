package packet

import "github.com/Gophercraft/core/vsn"

type ClientCacheVersion struct {
	Build vsn.Build
}

func (ccv *ClientCacheVersion) Encode(build vsn.Build, out *WorldPacket) error {
	out.Type = SMSG_CLIENTCACHE_VERSION
	out.WriteUint32(uint32(ccv.Build))
	return nil
}

func (ccv *ClientCacheVersion) Decode(build vsn.Build, in *WorldPacket) error {
	ccv.Build = vsn.Build(in.ReadUint32())
	return nil
}
