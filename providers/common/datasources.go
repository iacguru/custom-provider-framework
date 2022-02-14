package common

import (
	cpf "github.com/iacguru/custom-provider-framework"
	git "github.com/iacguru/custom-provider-framework/providers/terraform-provider-github"
	hcp "github.com/iacguru/custom-provider-framework/providers/terraform-provider-hashicups"
)

var Data = cpf.CustomResources{
	DataSourcesMap: []cpf.ResourcMap{
		hcp.CoffeeDataSource,
		git.WorkflowDataSource,
	},
}
