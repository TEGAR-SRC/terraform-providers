package iplist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudIPLists() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudIPListsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter IP lists.",
			},
			"ip_lists": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"entry_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_by_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ip_list_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_list_entry_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceOnidelCloudIPListsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		IPLists          []ipListSummaryData `json:"ip_lists"`
		IPListLimit      int                 `json:"ip_list_limit"`
		IPListEntryLimit int                 `json:"ip_list_entry_limit"`
	}
	err := client.Get("/network/ip_lists", params, &result)
	if err != nil {
		return diag.Errorf("Error listing IP lists: %s", err)
	}

	lists := make([]interface{}, len(result.IPLists))
	for i, l := range result.IPLists {
		lists[i] = map[string]interface{}{
			"id":             l.ID,
			"name":           l.Name,
			"description":    l.Description,
			"entry_count":    l.EntryCount,
			"used_by_count":  l.UsedByCount,
			"created_at":     l.CreatedAt,
		}
	}
	d.Set("ip_lists", lists)
	d.Set("ip_list_limit", result.IPListLimit)
	d.Set("ip_list_entry_limit", result.IPListEntryLimit)
	d.SetId("ip_lists")

	return nil
}

type ipListSummaryData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EntryCount  int    `json:"entry_count"`
	UsedByCount int    `json:"used_by_count"`
	CreatedAt   string `json:"created_at"`
}
