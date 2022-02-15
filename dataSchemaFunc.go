package cpf

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func TypeString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func TypeInt() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}
