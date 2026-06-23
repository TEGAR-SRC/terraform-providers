package instancetype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudInstancePrice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudInstancePriceRead,
		Schema: map[string]*schema.Schema{
			"vcpu": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Number of vCPUs.",
			},
			"ram": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Amount of RAM in MB.",
			},
			"disk": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Disk size in GB.",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Deployment location.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance type UUID.",
			},
			"currency": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "usd",
				Description: "Currency (aud, usd, eur).",
			},
			"instance_type_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"price_per_month": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"price_per_quarter": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"price_per_semiannual": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"price_per_annual": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"bw": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"net_rate": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceOnidelCloudInstancePriceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{
		"vcpu":          d.Get("vcpu").(string),
		"ram":           d.Get("ram").(string),
		"disk":          d.Get("disk").(string),
		"location":      d.Get("location").(string),
		"instance_type": d.Get("instance_type").(string),
	}
	if v, ok := d.GetOk("currency"); ok {
		params["currency"] = v.(string)
	}

	var result struct {
		InstanceType       string  `json:"instance_type"`
		VCPU               int     `json:"vcpu"`
		RAM                int     `json:"ram"`
		Disk               int     `json:"disk"`
		BW                 int     `json:"bw"`
		NetRate            int     `json:"net_rate"`
		PricePerMonth      float64 `json:"price_per_month"`
		PricePerQuarter    float64 `json:"price_per_quarter"`
		PricePerSemiannual float64 `json:"price_per_semiannual"`
		PricePerAnnual     float64 `json:"price_per_annual"`
		Currency           string  `json:"currency"`
	}
	err := client.Get("/instance_price", params, &result)
	if err != nil {
		return diag.Errorf("Error getting instance price: %s", err)
	}

	d.Set("instance_type_id", result.InstanceType)
	d.Set("price_per_month", result.PricePerMonth)
	d.Set("price_per_quarter", result.PricePerQuarter)
	d.Set("price_per_semiannual", result.PricePerSemiannual)
	d.Set("price_per_annual", result.PricePerAnnual)
	d.Set("bw", result.BW)
	d.Set("net_rate", result.NetRate)
	d.SetId("instance_price")

	return nil
}
