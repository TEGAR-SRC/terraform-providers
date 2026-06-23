package firewall

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudFirewallGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudFirewallGroupsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter firewall groups.",
			},
			"firewall_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
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
						"instance_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rule_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudFirewallGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		Groups []firewallGroupData `json:"firewall_groups"`
	}
	err := client.Get("/network/firewalls", params, &result)
	if err != nil {
		return diag.Errorf("Error listing firewall groups: %s", err)
	}

	groups := make([]interface{}, len(result.Groups))
	for i, g := range result.Groups {
		groups[i] = map[string]interface{}{
			"id":             g.ID,
			"description":    g.Description,
			"created":        g.Created,
			"updated":        g.Updated,
			"instance_count": g.InstanceCount,
			"rule_count":     g.RuleCount,
		}
	}
	d.Set("firewall_groups", groups)
	d.SetId("firewall_groups")

	return nil
}

type firewallGroupData struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	InstanceCount int    `json:"instance_count"`
	RuleCount     int    `json:"rule_count"`
}
