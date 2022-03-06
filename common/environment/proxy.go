package environment

import (
	"github.com/bmclab-git/v2ray-core/v5/features/extension/storage"
)

type ProxyEnvironmentCapabilitySet interface {
	BaseEnvironmentCapabilitySet
	InstanceNetworkCapabilitySet
	TransientStorage() storage.ScopedTransientStorage
}

type ProxyEnvironment interface {
	ProxyEnvironmentCapabilitySet
	NarrowScope(key []byte) (ProxyEnvironment, error)
	NarrowScopeToTransport(key []byte) (TransportEnvironment, error)
	doNotImpl()
}
