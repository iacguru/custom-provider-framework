package providers

import (
	cpf "github.com/iacguru/custom-provider-framework"
)

var Data = cpf.CustomResources{
	DataSourcesMap: []cpf.ResourcMap{CoffeeDataSource},
}
