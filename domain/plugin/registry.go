package plugin

type Registry interface {
	// Initialize initializes the plugin once when it is loaded.
	Initialize() error
	// IsInitialized returns whether or not the plugin has been initialized after loading.
	IsInitialized() bool

	// ValidatePlugin validates that the prerequisites of a plugin in the registry are given for the plugin's manifest.
	//
	// Prerequisites may be defined as dependencies, versions, and licensing.
	ValidatePlugin(id ID) (Manifest, error)

	// LoadPlugin loads a plugin from the registry into the local environment and memory via the passed plugin ID.
	LoadPlugin(manifest Manifest) (Plugin, error)
}
