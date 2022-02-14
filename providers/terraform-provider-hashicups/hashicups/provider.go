package hashicups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	cpf "github.com/iacguru/custom-provider-framework"
)

var pSchema1 = cpf.SchemaMap{
	"username": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
	},
}

var pSchema2 = cpf.SchemaMap{
	"password": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
	},
}

var PSch = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{
		pSchema1,
		pSchema2,
	},
}

var ConfigContex = cpf.CustomConfigureContextFunc{
	ConfigureContextFunc: providerConfigure,
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := hashicups.NewClient(nil, &username, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := hashicups.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
