package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cpf "github.com/iacguru/custom-provider-framework"
)

var workflows = cpf.SchemaMap{
	"workflows": &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: cpf.SchemaMap{
				"id": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"path": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"state": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"url": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"created_at": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"updated_at": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
}
