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
	var newClient client.Client
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	gitToken := d.Get("gh_token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := newClient.HashiCupsClient(nil, &username, &password)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	if gitToken != "" {
		return newClient.GitHubNewClient(gitToken), diags
	}

	c, err := newClient.HashiCupsClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
