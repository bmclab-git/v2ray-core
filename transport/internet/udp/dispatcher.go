package udp

import (
	"context"

	"github.com/bmclab-git/v2ray-core/v5/common"
	"github.com/bmclab-git/v2ray-core/v5/common/buf"
	"github.com/bmclab-git/v2ray-core/v5/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}
