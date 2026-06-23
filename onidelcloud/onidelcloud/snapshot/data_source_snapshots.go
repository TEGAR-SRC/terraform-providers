package snapshot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudSnapshots() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudSnapshotsRead,
		Schema: map[string]*schema.Schema{
			"snapshots": {
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
						"desc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudSnapshotsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var result []snapshotData
	err := client.Get("/snapshots", nil, &result)
	if err != nil {
		return diag.Errorf("Error listing snapshots: %s", err)
	}

	snapshots := make([]interface{}, len(result))
	for i, s := range result {
		snapshots[i] = map[string]interface{}{
			"id":         s.ID,
			"name":       s.Name,
			"desc":       s.Desc,
			"size":       s.Size,
			"status":     s.Status,
			"created_at": s.CreatedAt,
		}
	}
	d.Set("snapshots", snapshots)
	d.SetId("snapshots")

	return nil
}

type snapshotData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Size      int    `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
