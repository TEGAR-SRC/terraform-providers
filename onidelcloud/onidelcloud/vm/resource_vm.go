package vm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudVMCreate,
		ReadContext:   resourceOnidelCloudVMRead,
		UpdateContext: resourceOnidelCloudVMUpdate,
		DeleteContext: resourceOnidelCloudVMDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "VM hostname.",
			},
			"payment_cycle": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "monthly",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"hourly", "monthly", "quarterly", "semiannually", "annually", "biennially", "triennially"}, false),
				Description:  "Billing cycle.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance type UUID.",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Location/region.",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Number of vCPUs.",
			},
			"ram": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "RAM in MB.",
			},
			"disk": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Disk size in GB.",
			},
			"os": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "OS template ID.",
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Snapshot UUID to deploy from.",
			},
			"ssh_keys": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "SSH key IDs to add.",
			},
			"vpcs": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "VPC IDs to attach.",
			},
			"firewall_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Firewall group UUID.",
			},
			"ipv6": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable IPv6.",
			},
			"disable_ssh_blocking": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Disable outgoing SSH blocking.",
			},
			"startup_script_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Startup script UUID.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VM status.",
			},
			"main_ipv4": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary IPv4 address.",
			},
			"main_ipv6": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Primary IPv6 address.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp.",
			},
			"recurring_amount": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Recurring billing amount.",
			},
		},
	}
}

type vmResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	VCPU            int     `json:"vcpu"`
	RAM             int     `json:"ram"`
	Disk            int     `json:"disk"`
	MainIPv4        string  `json:"main_ipv4"`
	MainIPv6        string  `json:"main_ipv6"`
	Template        string  `json:"template"`
	FirewallGroupID string  `json:"firewall_group_id"`
	CreatedAt       string  `json:"created_at"`
	Status          string  `json:"status"`
	RecurringAmount float64 `json:"recurring_amount"`
}

func resourceOnidelCloudVMCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":                 d.Get("name").(string),
		"payment_cycle":        d.Get("payment_cycle").(string),
		"instance_type":        d.Get("instance_type").(string),
		"location":             d.Get("location").(string),
		"cpu":                  d.Get("cpu").(int),
		"ram":                  d.Get("ram").(int),
		"disk":                 d.Get("disk").(int),
		"ipv6":                 d.Get("ipv6").(bool),
		"disable_ssh_blocking": d.Get("disable_ssh_blocking").(bool),
	}

	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}
	if v, ok := d.GetOk("os"); ok {
		body["os"] = v.(int)
	}
	if v, ok := d.GetOk("snapshot_id"); ok {
		body["snapshot_id"] = v.(string)
	}
	if v, ok := d.GetOk("firewall_group_id"); ok {
		body["firewall_group_id"] = v.(string)
	}
	if v, ok := d.GetOk("startup_script_id"); ok {
		body["startup_script_id"] = v.(string)
	}
	if v, ok := d.GetOk("ssh_keys"); ok {
		body["ssh_keys"] = v.([]interface{})
	}
	if v, ok := d.GetOk("vpcs"); ok {
		body["vpcs"] = v.([]interface{})
	}

	log.Printf("[DEBUG] VM create configuration: %#v", body)
	err := client.Post("/vm", body, nil)
	if err != nil {
		return diag.Errorf("Error creating VM: %s", err)
	}

	// List to find created VM
	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}
	var listResult []vmResponse
	err = client.Get("/vm", params, &listResult)
	if err != nil {
		return diag.Errorf("Error listing VMs: %s", err)
	}

	for _, vm := range listResult {
		if vm.Name == d.Get("name").(string) {
			d.SetId(vm.ID)
			log.Printf("[INFO] VM created: %s", vm.ID)
			break
		}
	}

	if d.Id() == "" && len(listResult) > 0 {
		d.SetId(listResult[0].ID)
	}

	return resourceOnidelCloudVMRead(ctx, d, meta)
}

func resourceOnidelCloudVMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result vmResponse
	err := client.Get("/vm/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] VM not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.Name)
	d.Set("cpu", result.VCPU)
	d.Set("ram", result.RAM)
	d.Set("disk", result.Disk)
	d.Set("main_ipv4", result.MainIPv4)
	d.Set("main_ipv6", result.MainIPv6)
	d.Set("firewall_group_id", result.FirewallGroupID)
	d.Set("status", result.Status)
	d.Set("created_at", result.CreatedAt)
	d.Set("recurring_amount", result.RecurringAmount)
	return nil
}

func resourceOnidelCloudVMUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{}
	if d.HasChange("name") {
		body["name"] = d.Get("name").(string)
	}
	if d.HasChange("firewall_group_id") {
		body["firewall_group_id"] = d.Get("firewall_group_id").(string)
	}
	if d.HasChange("ipv6") {
		body["enable_ipv6"] = d.Get("ipv6").(bool)
	}

	if len(body) == 0 {
		return nil
	}

	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Patch("/vm/"+d.Id(), body, nil)
	if err != nil {
		return diag.Errorf("Error updating VM: %s", err)
	}

	return resourceOnidelCloudVMRead(ctx, d, meta)
}

func resourceOnidelCloudVMDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	log.Printf("[INFO] Destroying VM: %s", d.Id())
	err := client.Delete(fmt.Sprintf("/vm/%s", d.Id()), nil)
	if err != nil {
		return diag.Errorf("Error destroying VM: %s", err)
	}

	d.SetId("")
	return nil
}
