package startupscript

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudStartupScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudStartupScriptCreate,
		ReadContext:   resourceOnidelCloudStartupScriptRead,
		UpdateContext: resourceOnidelCloudStartupScriptUpdate,
		DeleteContext: resourceOnidelCloudStartupScriptDelete,
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
				Description: "Script name.",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Script content.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp.",
			},
			"updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last update timestamp.",
			},
		},
	}
}

type startupScriptSummary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type startupScriptDetail struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func resourceOnidelCloudStartupScriptCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":    d.Get("name").(string),
		"content": d.Get("content").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	var result struct {
		Script startupScriptDetail `json:"script"`
	}
	log.Printf("[DEBUG] Startup Script create: %#v", body)
	err := client.Post("/startup_scripts", body, &result)
	if err != nil {
		return diag.Errorf("Error creating Startup Script: %s", err)
	}

	d.SetId(result.Script.ID)
	log.Printf("[INFO] Startup Script created: %s", result.Script.ID)
	return resourceOnidelCloudStartupScriptRead(ctx, d, meta)
}

func resourceOnidelCloudStartupScriptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		Script startupScriptDetail `json:"script"`
	}
	err := client.Get("/startup_scripts/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] Startup Script not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.Script.Name)
	d.Set("content", result.Script.Content)
	d.Set("created", result.Script.Created)
	d.Set("updated", result.Script.Updated)
	return nil
}

func resourceOnidelCloudStartupScriptUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":    d.Get("name").(string),
		"content": d.Get("content").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Patch("/startup_scripts/"+d.Id(), body, nil)
	if err != nil {
		return diag.Errorf("Error updating Startup Script: %s", err)
	}

	return resourceOnidelCloudStartupScriptRead(ctx, d, meta)
}

func resourceOnidelCloudStartupScriptDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting Startup Script: %s", d.Id())
	err := client.Delete("/startup_scripts/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting Startup Script: %s", err)
	}

	d.SetId("")
	return nil
}
