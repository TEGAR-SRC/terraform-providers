package teams

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudTeams() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudTeamsRead,
		Schema: map[string]*schema.Schema{
			"teams": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudTeamsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var teams []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Role string `json:"role"`
	}
	err := client.Get("/teams", nil, &teams)
	if err != nil {
		return diag.Errorf("Error listing teams: %s", err)
	}

	result := make([]interface{}, len(teams))
	for i, t := range teams {
		result[i] = map[string]interface{}{
			"id":   t.ID,
			"name": t.Name,
			"role": t.Role,
		}
	}
	d.Set("teams", result)
	d.SetId("teams")

	return nil
}
