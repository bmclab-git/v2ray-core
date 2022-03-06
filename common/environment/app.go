package environment

import (
	"github.com/bmclab-git/v2ray-core/v5/features/extension/storage"
)

type AppEnvironmentCapabilitySet interface {
	BaseEnvironmentCapabilitySet
	SystemNetworkCapabilitySet
	InstanceNetworkCapabilitySet
	FileSystemCapabilitySet
	PersistentStorage() storage.ScopedPersistentStorage
	TransientStorage() storage.ScopedTransientStorage
}

type AppEnvironment interface {
	AppEnvironmentCapabilitySet
	NarrowScope(key []byte) (AppEnvironment, error)
	doNotImpl()
}
