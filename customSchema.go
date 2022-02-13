package cpf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SchemaMap map[string]*schema.Schema

type ResourcMap map[string]*schema.Resource

type CustomSchema struct {
	Schemas []SchemaMap
}

type CustomResources struct {
	ResourcesMaps  []ResourcMap
	DataSourcesMap []ResourcMap
}

type CustomConfigureContextFunc struct {
	ConfigureContextFunc schema.ConfigureContextFunc
}

func (c *CustomSchema) Schema() SchemaMap {
	var schema = make(SchemaMap)
	for _, v := range c.Schemas {
		for x, y := range v {
			if _, ok := schema[x]; !ok {
				schema[x] = y
			}
		}
	}
	return schema
}

func (c *CustomResources) Resources() ResourcMap {
	var resource = make(ResourcMap)
	for _, v := range c.ResourcesMaps {
		for x, y := range v {
			if _, ok := resource[x]; !ok {
				resource[x] = y
			}
		}
	}
	return resource
}

func (c *CustomResources) DataSources() ResourcMap {
	var resource = make(ResourcMap)
	for _, v := range c.DataSourcesMap {
		for x, y := range v {
			if _, ok := resource[x]; !ok {
				resource[x] = y
			}
		}
	}
	return resource
}
