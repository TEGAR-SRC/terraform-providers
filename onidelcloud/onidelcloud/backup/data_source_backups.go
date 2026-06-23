package backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudBackups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudBackupsRead,
		Schema: map[string]*schema.Schema{
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance": {
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

func dataSourceOnidelCloudBackupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var result []backupData
	err := client.Get("/backups", nil, &result)
	if err != nil {
		return diag.Errorf("Error listing backups: %s", err)
	}

	backups := make([]interface{}, len(result))
	for i, b := range result {
		backups[i] = map[string]interface{}{
			"id":         b.ID,
			"instance":   b.Instance,
			"size":       b.Size,
			"status":     b.Status,
			"created_at": b.CreatedAt,
		}
	}
	d.Set("backups", backups)
	d.SetId("backups")

	return nil
}

type backupData struct {
	ID        string `json:"id"`
	Instance  string `json:"instance"`
	Size      int    `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
