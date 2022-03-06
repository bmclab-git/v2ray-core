package environment

import (
	"github.com/bmclab-git/v2ray-core/v5/features/extension/storage"
)

type TransportEnvironmentCapacitySet interface {
	BaseEnvironmentCapabilitySet
	SystemNetworkCapabilitySet
	InstanceNetworkCapabilitySet
	TransientStorage() storage.ScopedTransientStorage
}

type TransportEnvironment interface {
	TransportEnvironmentCapacitySet
	NarrowScope(key []byte) (TransportEnvironment, error)
	doNotImpl()
}
