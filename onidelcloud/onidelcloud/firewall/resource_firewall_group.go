package firewall

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudFirewallGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudFirewallGroupCreate,
		ReadContext:   resourceOnidelCloudFirewallGroupRead,
		UpdateContext: resourceOnidelCloudFirewallGroupUpdate,
		DeleteContext: resourceOnidelCloudFirewallGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "Description of the firewall group.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp.",
			},
			"updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last updated timestamp.",
			},
			"instance_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of VMs using this firewall group.",
			},
			"rule_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of rules in this group.",
			},
		},
	}
}

type firewallGroupResponse struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	InstanceCount int    `json:"instance_count"`
	RuleCount     int    `json:"rule_count"`
}

func resourceOnidelCloudFirewallGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"description": d.Get("description").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	var result struct {
		FirewallGroup firewallGroupResponse `json:"firewall_group"`
	}
	log.Printf("[DEBUG] Firewall Group create: %#v", body)
	err := client.Post("/network/firewalls", body, &result)
	if err != nil {
		return diag.Errorf("Error creating Firewall Group: %s", err)
	}

	d.SetId(result.FirewallGroup.ID)
	log.Printf("[INFO] Firewall Group created: %s", d.Id())
	return resourceOnidelCloudFirewallGroupRead(ctx, d, meta)
}

func resourceOnidelCloudFirewallGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	var result struct {
		FirewallGroup firewallGroupResponse `json:"firewall_group"`
	}
	err := client.Get("/network/firewalls/"+d.Id(), nil, &result)
	if err != nil {
		log.Printf("[DEBUG] Firewall Group not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("description", result.FirewallGroup.Description)
	d.Set("created", result.FirewallGroup.Created)
	d.Set("updated", result.FirewallGroup.Updated)
	d.Set("instance_count", result.FirewallGroup.InstanceCount)
	d.Set("rule_count", result.FirewallGroup.RuleCount)
	return nil
}

func resourceOnidelCloudFirewallGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"description": d.Get("description").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Put("/network/firewalls/"+d.Id(), body, nil)
	if err != nil {
		return diag.Errorf("Error updating Firewall Group: %s", err)
	}

	return resourceOnidelCloudFirewallGroupRead(ctx, d, meta)
}

func resourceOnidelCloudFirewallGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting Firewall Group: %s", d.Id())
	err := client.Delete("/network/firewalls/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting Firewall Group: %s", err)
	}

	d.SetId("")
	return nil
}
