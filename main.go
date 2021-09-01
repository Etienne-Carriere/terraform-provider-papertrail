package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/Etienne-Carriere/terraform-provider-papertrail/papertrail"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: papertrail.Provider,
	})
}
