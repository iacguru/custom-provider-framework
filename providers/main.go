package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/iacguru/custom-provider-framework/providers/common"

	cpf "github.com/iacguru/custom-provider-framework"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return cpf.Provider(common.PSchema, common.Resource, common.Data, common.ConfigContex)
		},
	})
}
