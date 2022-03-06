package udp

import (
	"github.com/bmclab-git/v2ray-core/v5/common/buf"
	"github.com/bmclab-git/v2ray-core/v5/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
