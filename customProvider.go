package cpf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:               CustomSchema.Schema(),
		ResourcesMap:         CustomResources.Resources(),
		DataSourcesMap:       CustomResources.DataSources(),
		ConfigureContextFunc: CustomConfigureContextFunc.ConfigureContextFunc,
	}
}
