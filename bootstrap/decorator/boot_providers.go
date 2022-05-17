package decorator

import (
	"github.com/confetti-framework/framework/inter"
	"src/app/providers"
)

type BootProviders struct{}

func (r BootProviders) Bootstrap(container inter.Container) inter.Container {
	for _, bootstrapper := range providers.Providers.BootProviders {
		container = bootstrapper.Boot(container)
	}

	return container
}
