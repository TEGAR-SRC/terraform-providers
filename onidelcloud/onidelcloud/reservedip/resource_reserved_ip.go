package reservedip

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudReservedIP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudReservedIPCreate,
		ReadContext:   resourceOnidelCloudReservedIPRead,
		UpdateContext: resourceOnidelCloudReservedIPUpdate,
		DeleteContext: resourceOnidelCloudReservedIPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID.",
			},
			"ip_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ipv4",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
				Description:  "IPv4 (/32) or IPv6 (/64).",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Location name.",
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Description:  "Custom name (defaults to IP address).",
			},
			"ip_addr": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reserved IP address.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status (active/suspended).",
			},
			"anchor_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VM IP to attach to. Set to null to detach.",
			},
		},
	}
}

type reservedIPResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Status     string `json:"status"`
	IPAddr     string `json:"ip_addr"`
	Attachment *struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"attachment,omitempty"`
}

func resourceOnidelCloudReservedIPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"ip_type":  d.Get("ip_type").(string),
		"location": d.Get("location").(string),
	}
	if v, ok := d.GetOk("name"); ok {
		body["name"] = v.(string)
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	var result struct {
		RipID  string `json:"rip_id"`
		IPAddr string `json:"ip_addr"`
	}
	log.Printf("[DEBUG] Reserved IP create: %#v", body)
	err := client.Post("/network/reserved_ips", body, &result)
	if err != nil {
		return diag.Errorf("Error creating Reserved IP: %s", err)
	}

	d.SetId(result.RipID)
	d.Set("ip_addr", result.IPAddr)
	log.Printf("[INFO] Reserved IP created: %s (%s)", result.RipID, result.IPAddr)
	return resourceOnidelCloudReservedIPRead(ctx, d, meta)
}

func resourceOnidelCloudReservedIPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result reservedIPResponse
	err := client.Get("/network/reserved_ips/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] Reserved IP not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.Name)
	d.Set("location", result.Location)
	d.Set("status", result.Status)
	d.Set("ip_addr", result.IPAddr)
	return nil
}

func resourceOnidelCloudReservedIPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{}
	if d.HasChange("name") {
		body["name"] = d.Get("name").(string)
	}
	if d.HasChange("anchor_ip") {
		v, _ := d.GetOk("anchor_ip")
		body["anchor_ip"] = v.(string)
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	if len(body) > 0 {
		err := client.Patch("/network/reserved_ips/"+d.Id(), body, nil)
		if err != nil {
			return diag.Errorf("Error updating Reserved IP: %s", err)
		}
	}

	return resourceOnidelCloudReservedIPRead(ctx, d, meta)
}

func resourceOnidelCloudReservedIPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting Reserved IP: %s", d.Id())
	err := client.Delete("/network/reserved_ips/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting Reserved IP: %s", err)
	}

	d.SetId("")
	return nil
}
