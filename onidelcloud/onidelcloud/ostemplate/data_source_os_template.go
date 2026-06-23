package ostemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudOSTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudOSTemplateRead,
		Schema: map[string]*schema.Schema{
			"templates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudOSTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var templates []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Family string `json:"family"`
	}
	err := client.Get("/os_templates", nil, &templates)
	if err != nil {
		return diag.Errorf("Error listing OS templates: %s", err)
	}

	result := make([]interface{}, len(templates))
	for i, t := range templates {
		result[i] = map[string]interface{}{
			"id":     t.ID,
			"name":   t.Name,
			"family": t.Family,
		}
	}
	d.Set("templates", result)
	d.SetId("os_templates")

	return nil
}
