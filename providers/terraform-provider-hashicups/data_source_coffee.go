package hashicups

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cpf "github.com/iacguru/custom-provider-framework"
)

var coffeeS = cpf.SchemaMap{
	"coffees": &schema.Schema{
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
				"teaser": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"price": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"image": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"ingredients": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: cpf.SchemaMap{
							"ingredient_id": &schema.Schema{
								Type:     schema.TypeInt,
								Computed: true,
							},
						},
					},
				},
			},
		},
	},
}

var CoffeeDataSource = cpf.ResourcMap{
	"hashicups_coffees": dataSourceCoffees(),
}

var sch = cpf.CustomSchema{
	Schemas: []cpf.SchemaMap{coffeeS},
}

func dataSourceCoffees() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCoffeesRead,
		Schema:      sch.Schema(),
	}
}

func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://localhost:19090"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	coffees := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&coffees)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("coffees", coffees); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
