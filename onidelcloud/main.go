package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onidelcloud.Provider})
}
