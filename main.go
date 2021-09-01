package main

import (
	"github.com/Etienne-Carriere/terraform-provider-papertrail/papertrail"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return papertrail.Provider()
		},
	})
}
