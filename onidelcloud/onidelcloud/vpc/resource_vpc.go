package vpc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudVPC() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudVPCCreate,
		ReadContext:   resourceOnidelCloudVPCRead,
		UpdateContext: resourceOnidelCloudVPCUpdate,
		DeleteContext: resourceOnidelCloudVPCDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID. Default team will be used if not provided.",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 200),
				Description:  "Name of the VPC.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 200),
				Description:  "Description of the VPC.",
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Location (region) for the VPC.",
			},
			"v4_subnet": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IPv4 subnet prefix (e.g. 10.0.0.0).",
			},
			"v4_subnet_mask": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IPv4 subnet mask (e.g. 24).",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPC status.",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation date.",
			},
		},
	}
}

type vpcResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DateCreated  string `json:"date_created"`
	Status       string `json:"status"`
	Location     string `json:"location"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask string `json:"v4_subnet_mask"`
}

func resourceOnidelCloudVPCCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":           d.Get("name").(string),
		"description":    d.Get("description").(string),
		"location":       d.Get("location").(string),
		"v4_subnet":      d.Get("v4_subnet").(string),
		"v4_subnet_mask": d.Get("v4_subnet_mask").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	log.Printf("[DEBUG] VPC create configuration: %#v", body)
	err := client.Post("/network/vpcs", body, nil)
	if err != nil {
		return diag.Errorf("Error creating VPC: %s", err)
	}

	// List to find created VPC
	listParams := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		listParams["team_id"] = v.(string)
	}
	if v, ok := d.GetOk("location"); ok {
		listParams["location"] = v.(string)
	}

	var listResult struct {
		VPCs []vpcResponse `json:"vpcs"`
	}
	err = client.Get("/network/vpcs", listParams, &listResult)
	if err != nil {
		return diag.Errorf("Error listing VPCs: %s", err)
	}

	for _, v := range listResult.VPCs {
		if v.Name == d.Get("name").(string) {
			d.SetId(v.ID)
			log.Printf("[INFO] VPC created: %s", v.ID)
			break
		}
	}

	if d.Id() == "" && len(listResult.VPCs) > 0 {
		d.SetId(listResult.VPCs[0].ID)
	}

	return resourceOnidelCloudVPCRead(ctx, d, meta)
}

func resourceOnidelCloudVPCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		VPC vpcResponse `json:"vpc"`
	}
	err := client.Get("/network/vpcs/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] VPC not found, removing from state: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.VPC.Name)
	d.Set("description", result.VPC.Description)
	d.Set("location", result.VPC.Location)
	d.Set("v4_subnet", result.VPC.V4Subnet)
	d.Set("v4_subnet_mask", result.VPC.V4SubnetMask)
	d.Set("status", result.VPC.Status)
	d.Set("date_created", result.VPC.DateCreated)
	return nil
}

func resourceOnidelCloudVPCUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":        d.Get("name").(string),
		"description": d.Get("description").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Patch("/network/vpcs/"+d.Id(), body, nil)
	if err != nil {
		return diag.Errorf("Error updating VPC: %s", err)
	}

	return resourceOnidelCloudVPCRead(ctx, d, meta)
}

func resourceOnidelCloudVPCDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting VPC: %s", d.Id())
	err := client.Delete("/network/vpcs/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting VPC: %s", err)
	}

	d.SetId("")
	return nil
}
