package common

import (
	cpf "github.com/iacguru/custom-provider-framework"
	gh "github.com/iacguru/custom-provider-framework/providers/terraform-provider-github"
	hcp "github.com/iacguru/custom-provider-framework/providers/terraform-provider-hashicups"
)

var PSchema = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{
		hcp.HashicupsProviderSchema,
		gh.GitHubProviderSchema,
	},
}
