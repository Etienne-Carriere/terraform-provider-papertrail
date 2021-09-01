package main

import (
	"github.com/Etienne-Carriere/terraform-provider-papertrail/papertrail"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: papertrail.Provider,
	})
}
