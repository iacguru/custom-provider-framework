package hashicups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	cpf "github.com/iacguru/custom-provider-framework"
)

var HashicupsProviderSchema = cpf.SchemaMap{
	"username": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
	},
	"password": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
	},
}
