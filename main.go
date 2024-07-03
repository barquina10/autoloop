package autoloop

import (
	"github.com/barquina10/autoloop/autoloop"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: autoloop.Provider,
	})
}
