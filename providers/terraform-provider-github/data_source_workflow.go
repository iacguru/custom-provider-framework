package github

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cpf "github.com/iacguru/custom-provider-framework"
	"github.com/iacguru/custom-provider-framework/providers/client"
	action "github.com/iacguru/custom-provider-framework/providers/data/github"
)

var workflows = cpf.SchemaMap{
	"total_count": cpf.TypeIntComputed(),
	"workflows": &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: cpf.SchemaMap{
				"id":         cpf.TypeIntComputed(),
				"name":       cpf.TypeStringComputed(),
				"path":       cpf.TypeStringComputed(),
				"state":      cpf.TypeStringComputed(),
				"url":        cpf.TypeStringComputed(),
				"created_at": cpf.TypeStringComputed(),
				"updated_at": cpf.TypeStringComputed(),
				"html_url":   cpf.TypeStringComputed(),
				"badge_url":  cpf.TypeStringComputed(),
				"node_id":    cpf.TypeStringComputed(),
			},
		},
	},
	"owner": cpf.TypeStringArgument(),
	"repo":  cpf.TypeStringArgument(),
}

var WorkflowDataSource = cpf.ResourcMap{
	"cpf_git_workflows": dataSourceGitWorkflows(),
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
	c := m.(*client.Client)
	orgOwner := d.Get("owner").(string)
	repo := d.Get("repo").(string)
	c.URL = fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows", orgOwner, repo)
	c.Method = "GET"
	var diags diag.Diagnostics
	err := c.GetGitWorkflows()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("total_count", c.Workflows.TotalCount); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("workflows", flattenData(c.Workflows)); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenData(Workflows *action.Workflows) []interface{} {
	if Workflows != nil {
		wfs := make([]interface{}, len(*Workflows.Workflows), len(*Workflows.Workflows))
		for i, workflow := range *Workflows.Workflows {
			wf := make(map[string]interface{})

			wf["id"] = workflow.ID
			wf["node_id"] = workflow.NodeID
			wf["name"] = workflow.Name
			wf["path"] = workflow.Path
			wf["state"] = workflow.State
			wf["created_at"] = workflow.CreatedAt
			wf["updated_at"] = workflow.NodeID
			wf["url"] = workflow.URL
			wf["html_url"] = workflow.HTMLURL
			wf["badge_url"] = workflow.BadgeURL
			wfs[i] = wf
		}
		return wfs
	}
	return make([]interface{}, 0)
}
