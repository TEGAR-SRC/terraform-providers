package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudVPCs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudVPCsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter VPCs.",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by location.",
			},
			"vpcs": {
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
						"location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"v4_subnet": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"v4_subnet_mask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"date_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudVPCsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}
	if v, ok := d.GetOk("location"); ok {
		params["location"] = v.(string)
	}

	var result struct {
		VPCs []vpcResponse `json:"vpcs"`
	}
	err := client.Get("/network/vpcs", params, &result)
	if err != nil {
		return diag.Errorf("Error listing VPCs: %s", err)
	}

	vpcs := make([]interface{}, len(result.VPCs))
	for i, v := range result.VPCs {
		vpcs[i] = map[string]interface{}{
			"id":             v.ID,
			"name":           v.Name,
			"description":    v.Description,
			"location":       v.Location,
			"v4_subnet":      v.V4Subnet,
			"v4_subnet_mask": v.V4SubnetMask,
			"status":         v.Status,
			"date_created":   v.DateCreated,
		}
	}
	d.Set("vpcs", vpcs)
	d.SetId("vpcs")

	return nil
}
