package vm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudVMs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudVMsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter VMs.",
			},
			"vms": {
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
						"vcpu": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ram": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"disk": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"main_ipv4": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"main_ipv6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"template": {
							Type:     schema.TypeString,
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
						"renewed_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"due_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_cycle": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"recurring_amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"payment_currency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"firewall_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudVMsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result []vmListData
	err := client.Get("/vm", params, &result)
	if err != nil {
		return diag.Errorf("Error listing VMs: %s", err)
	}

	vms := make([]interface{}, len(result))
	for i, v := range result {
		vms[i] = map[string]interface{}{
			"id":                v.ID,
			"name":              v.Name,
			"vcpu":              v.VCPU,
			"ram":               v.RAM,
			"disk":              v.Disk,
			"main_ipv4":         v.MainIPv4,
			"main_ipv6":         v.MainIPv6,
			"template":          v.Template,
			"status":            v.Status,
			"created_at":        v.CreatedAt,
			"renewed_at":        v.RenewedAt,
			"due_date":          v.DueDate,
			"billing_cycle":     v.BillingCycle,
			"recurring_amount":  v.RecurringAmount,
			"payment_currency":  v.PaymentCurrency,
			"firewall_group_id": v.FirewallGroupID,
		}
	}
	d.Set("vms", vms)
	d.SetId("vms")

	return nil
}

type vmListData struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	VCPU            int     `json:"vcpu"`
	RAM             int     `json:"ram"`
	Disk            int     `json:"disk"`
	MainIPv4        string  `json:"main_ipv4"`
	MainIPv6        string  `json:"main_ipv6"`
	Template        string  `json:"template"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
	RenewedAt       string  `json:"renewed_at"`
	DueDate         string  `json:"due_date"`
	BillingCycle    int     `json:"billing_cycle"`
	RecurringAmount float64 `json:"recurring_amount"`
	PaymentCurrency string  `json:"payment_currency"`
	FirewallGroupID string  `json:"firewall_group_id"`
}
