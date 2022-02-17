package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cpf "github.com/iacguru/custom-provider-framework"
	"github.com/iacguru/custom-provider-framework/providers/client"
)

var ConfigContex = cpf.CustomConfigureContextFunc{
	ConfigureContextFunc: providerConfigure,
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	gitToken := d.Get("github_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := client.HashiCupsClient(nil, &username, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	if gitToken != "" {
		return client.GitHubNewClient(gitToken), diags
	}

	c, err := client.HashiCupsClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
