package startupscript

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudStartupScripts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudStartupScriptsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter startup scripts.",
			},
			"scripts": {
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
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudStartupScriptsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		Scripts []startupScriptSummaryData `json:"scripts"`
	}
	err := client.Get("/startup_scripts", params, &result)
	if err != nil {
		return diag.Errorf("Error listing startup scripts: %s", err)
	}

	scripts := make([]interface{}, len(result.Scripts))
	for i, s := range result.Scripts {
		scripts[i] = map[string]interface{}{
			"id":      s.ID,
			"name":    s.Name,
			"created": s.Created,
			"updated": s.Updated,
		}
	}
	d.Set("scripts", scripts)
	d.SetId("startup_scripts")

	return nil
}

type startupScriptSummaryData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
