package firewall

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudFirewallRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudFirewallRuleCreate,
		ReadContext:   resourceOnidelCloudFirewallRuleRead,
		UpdateContext: resourceOnidelCloudFirewallRuleUpdate,
		DeleteContext: resourceOnidelCloudFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"firewall_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Firewall Group UUID.",
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp", "icmp"}, false),
				Description:  "Protocol (tcp, udp, icmp).",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Port number or range (e.g. '443', '8000:9000'). Not required for ICMP.",
			},
			"subnet": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IP/CIDR, special value (Cloudflare, Onidel, etc.), or IP list ref (list:{uuid}).",
			},
			"subnet_size": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIDR prefix length.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Description:  "Rule description.",
			},
			"ip_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP type (v4 or v6).",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rule action (allow).",
			},
		},
	}
}

type firewallRuleResponse struct {
	ID         string `json:"id"`
	Group      string `json:"group"`
	IPType     string `json:"ip_type"`
	Action     string `json:"action"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Subnet     string `json:"subnet"`
	SubnetSize string `json:"subnet_size"`
	Desc       string `json:"desc"`
}

func resourceOnidelCloudFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	firewallID := d.Get("firewall_group_id").(string)
	body := map[string]interface{}{
		"protocol":    d.Get("protocol").(string),
		"subnet":      d.Get("subnet").(string),
		"subnet_size": d.Get("subnet_size").(string),
	}
	if v, ok := d.GetOk("port"); ok {
		body["port"] = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		body["desc"] = v.(string)
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	path := fmt.Sprintf("/network/firewalls/%s/rules", firewallID)
	var result struct {
		FirewallRule firewallRuleResponse `json:"firewall_rule"`
	}
	log.Printf("[DEBUG] Firewall Rule create: %#v", body)
	err := client.Post(path, body, &result)
	if err != nil {
		return diag.Errorf("Error creating Firewall Rule: %s", err)
	}

	d.SetId(result.FirewallRule.ID)
	log.Printf("[INFO] Firewall Rule created: %s", d.Id())
	return resourceOnidelCloudFirewallRuleRead(ctx, d, meta)
}

func resourceOnidelCloudFirewallRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	firewallID := d.Get("firewall_group_id").(string)
	path := fmt.Sprintf("/network/firewalls/%s/rules/%s", firewallID, d.Id())

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		FirewallRule firewallRuleResponse `json:"firewall_rule"`
	}
	err := client.Get(path, params, &result)
	if err != nil {
		log.Printf("[DEBUG] Firewall Rule not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("protocol", result.FirewallRule.Protocol)
	d.Set("port", result.FirewallRule.Port)
	d.Set("subnet", result.FirewallRule.Subnet)
	d.Set("subnet_size", result.FirewallRule.SubnetSize)
	d.Set("description", result.FirewallRule.Desc)
	d.Set("ip_type", result.FirewallRule.IPType)
	d.Set("action", result.FirewallRule.Action)
	return nil
}

func resourceOnidelCloudFirewallRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	firewallID := d.Get("firewall_group_id").(string)
	path := fmt.Sprintf("/network/firewalls/%s/rules/%s", firewallID, d.Id())

	body := map[string]interface{}{
		"desc": d.Get("description").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Patch(path, body, nil)
	if err != nil {
		return diag.Errorf("Error updating Firewall Rule: %s", err)
	}

	return resourceOnidelCloudFirewallRuleRead(ctx, d, meta)
}

func resourceOnidelCloudFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	firewallID := d.Get("firewall_group_id").(string)
	path := fmt.Sprintf("/network/firewalls/%s/rules/%s", firewallID, d.Id())

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting Firewall Rule: %s", d.Id())
	err := client.Delete(path, params)
	if err != nil {
		return diag.Errorf("Error deleting Firewall Rule: %s", err)
	}

	d.SetId("")
	return nil
}
