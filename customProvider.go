package cpf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider(cs CustomSchema, cr, cdr CustomResources, ccf CustomConfigureContextFunc) *schema.Provider {
	return &schema.Provider{
		Schema:               cs.Schema(),
		ResourcesMap:         cr.Resources(),
		DataSourcesMap:       cr.DataSources(),
		ConfigureContextFunc: ccf.ConfigureContextFunc,
	}
}
