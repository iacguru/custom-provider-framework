package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cpf "github.com/iacguru/custom-provider-framework"
	"github.com/iacguru/custom-provider-framework/providers/common"
	"strconv"
	"time"
)

var workflows = cpf.SchemaMap{
	"total_count": &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	},
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

var WorkflowDataSource = cpf.ResourcMap{
	"git_workflows": dataSourceGitWorkflows(),
}

var sch = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{workflows},
}

func dataSourceGitWorkflows() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGitWorkflowsRead,
		Schema:      sch.Schema(),
	}
}

func dataSourceGitWorkflowsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*common.Client)
	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	c.URL = fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows", orgOwner, repo)
	c.Method = "GET"
	var diags diag.Diagnostics
	c.NewRequest()
	res, err := c.DoRequest()
	if err != nil {
		return diag.FromErr(err)
	}
	defer res.Body.Close()
	workflows := make([]map[string]interface{}, 0)
	err = json.NewDecoder(res.Body).Decode(&workflows)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("workflows", workflows); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
