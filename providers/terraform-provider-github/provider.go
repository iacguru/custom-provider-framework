package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	cpf "github.com/iacguru/custom-provider-framework"
)

var GitHubProviderSchema = cpf.SchemaMap{
	"github_token": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		DefaultFunc: schema.EnvDefaultFunc("GH_TOKEN", nil),
	},
}
