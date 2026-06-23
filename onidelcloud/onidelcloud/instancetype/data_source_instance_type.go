package instancetype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudInstanceType() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudInstanceTypeRead,
		Schema: map[string]*schema.Schema{
			"instance_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_vcpu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_ram": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_disk": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"network_rate": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"locations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudInstanceTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var types []struct {
		ID          string   `json:"id"`
		Type        string   `json:"type"`
		CPU         string   `json:"cpu"`
		MaxVCPU     int      `json:"max_vcpu"`
		MaxRAM      int      `json:"max_ram"`
		MaxDisk     int      `json:"max_disk"`
		NetworkRate int      `json:"network_rate"`
		Locations   []string `json:"locations"`
	}
	err := client.Get("/instance_types", nil, &types)
	if err != nil {
		return diag.Errorf("Error listing instance types: %s", err)
	}

	result := make([]interface{}, len(types))
	for i, t := range types {
		result[i] = map[string]interface{}{
			"id":           t.ID,
			"type":         t.Type,
			"cpu":          t.CPU,
			"max_vcpu":     t.MaxVCPU,
			"max_ram":      t.MaxRAM,
			"max_disk":     t.MaxDisk,
			"network_rate": t.NetworkRate,
			"locations":    t.Locations,
		}
	}
	d.Set("instance_types", result)
	d.SetId("instance_types")

	return nil
}
