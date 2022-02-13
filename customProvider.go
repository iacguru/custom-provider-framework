package cpf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	var cs CustomSchema
	var cr CustomResources
	var ccf CustomConfigureContextFunc
	return &schema.Provider{
		Schema:               cs.Schema(),
		ResourcesMap:         cr.Resources(),
		DataSourcesMap:       cr.DataSources(),
		ConfigureContextFunc: ccf.ConfigureContextFunc,
	}
}
