package common

import (
	cpf "github.com/iacguru/custom-provider-framework"
	hcp "github.com/iacguru/custom-provider-framework/providers/terraform-provider-hashicups/hashicups"
)

var Data = cpf.CustomResources{
	DataSourcesMap: []cpf.ResourcMap{
		hcp.CoffeeDataSource,
	},
}
