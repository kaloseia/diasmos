package plugin

import "fmt"

// TODO: What about a more Go-oriented Git versioning system as a registry?
// TODO:

type Installer struct {
	Registry Registry
}

func (i *Installer) InstallPlugin(pluginKey string) error {
	id := ID{}
	parseErr := id.FromKey(pluginKey)
	if parseErr != nil {
		return fmt.Errorf("plugin '%s' installation failed: %w", pluginKey, parseErr)
	}

	manifest, validateErr := i.Registry.ValidatePlugin(id)
	if validateErr != nil {
		return fmt.Errorf("plugin '%s' installation failed: %w", pluginKey, validateErr)
	}

	_, loadErr := i.Registry.LoadPlugin(manifest)
	if loadErr != nil {
		return fmt.Errorf("plugin '%s' installation failed: %w", pluginKey, loadErr)
	}

	// TODO: Installation and verification are more the responsibilities of an installer or "project manager" component rather than the registry itself!
	// 		 If we define an installer component, we should inject the registry and pass validation / loading through to the registry methods.

	/*projectPluginsDir := ""
	installation, installErr := i.installLoadedPlugin(loaded, projectPluginsDir)
	if installErr != nil {
		return fmt.Errorf("plugin '%s' installation failed: %w", pluginKey, installErr)
	}

	verifyErr := i.verifyInstallation(installation)
	if verifyErr != nil {
		return fmt.Errorf("plugin '%s' installation failed: %w", pluginKey, verifyErr)
	}*/
	return nil
}
