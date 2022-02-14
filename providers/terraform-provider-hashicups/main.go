package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-hashicups/hashicups"

	cpf "github.com/iacguru/custom-provider-framework"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return cpf.Provider(hashicups.PSch, hashicups.Resource, hashicups.Data, hashicups.ConfigContex)
		},
	})
}
