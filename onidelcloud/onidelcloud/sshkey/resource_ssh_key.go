package sshkey

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudSSHKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudSSHKeyCreate,
		ReadContext:   resourceOnidelCloudSSHKeyRead,
		UpdateContext: resourceOnidelCloudSSHKeyUpdate,
		DeleteContext: resourceOnidelCloudSSHKeyDelete,
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
				ValidateFunc: validation.NoZeroValues,
				Description:  "Name of the SSH key.",
			},
			"public_key": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
				Description:  "The public key content.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp.",
			},
		},
	}
}

func resourceOnidelCloudSSHKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name":    d.Get("name").(string),
		"ssh_key": d.Get("public_key").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	var result struct {
		SSHKey struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			SSHKey  string `json:"ssh_key"`
			Created string `json:"created"`
		} `json:"ssh_key"`
	}
	log.Printf("[DEBUG] SSH Key create configuration: %#v", body)
	err := client.Post("/ssh_keys", body, &result)
	if err != nil {
		return diag.Errorf("Error creating SSH Key: %s", err)
	}

	d.SetId(result.SSHKey.ID)
	log.Printf("[INFO] SSH Key created: %s", result.SSHKey.ID)
	return resourceOnidelCloudSSHKeyRead(ctx, d, meta)
}

func resourceOnidelCloudSSHKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		SSHKey struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			SSHKey  string `json:"ssh_key"`
			Created string `json:"created"`
		} `json:"ssh_key"`
	}
	err := client.Get("/ssh_keys/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] SSH Key not found, removing from state: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.SSHKey.Name)
	d.Set("public_key", result.SSHKey.SSHKey)
	d.Set("created", result.SSHKey.Created)
	return nil
}

func resourceOnidelCloudSSHKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	if !d.HasChanges("name", "public_key") {
		return nil
	}

	body := map[string]interface{}{
		"name":    d.Get("name").(string),
		"ssh_key": d.Get("public_key").(string),
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	err := client.Patch("/ssh_keys/"+d.Id(), body, nil)
	if err != nil {
		return diag.Errorf("Error updating SSH Key: %s", err)
	}

	return resourceOnidelCloudSSHKeyRead(ctx, d, meta)
}

func resourceOnidelCloudSSHKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting SSH Key: %s", d.Id())
	err := client.Delete("/ssh_keys/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting SSH Key: %s", err)
	}

	d.SetId("")
	return nil
}
