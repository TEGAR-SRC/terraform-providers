package objectstorage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudObjectStorageServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudObjectStorageServicesRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter object storage services.",
			},
			"services": {
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
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"capacity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_capacity": {
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
						"renewal_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudObjectStorageServicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		Services []objectStorageServiceData `json:"services"`
	}
	err := client.Get("/object-storage", params, &result)
	if err != nil {
		return diag.Errorf("Error listing object storage services: %s", err)
	}

	services := make([]interface{}, len(result.Services))
	for i, s := range result.Services {
		services[i] = map[string]interface{}{
			"id":            s.ID,
			"name":          s.Name,
			"location":      s.Location,
			"region":        s.Region,
			"endpoint":      s.Endpoint,
			"capacity":      s.Capacity,
			"used_capacity": s.UsedCapacity,
			"status":        s.Status,
			"created_at":    s.CreatedAt,
			"renewal_date":  s.RenewalDate,
		}
	}
	d.Set("services", services)
	d.SetId("object_storage_services")

	return nil
}

type objectStorageServiceData struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Region       string `json:"region"`
	Endpoint     string `json:"endpoint"`
	Capacity     int    `json:"capacity"`
	UsedCapacity int    `json:"used_capacity"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	RenewalDate  string `json:"renewal_date"`
}
