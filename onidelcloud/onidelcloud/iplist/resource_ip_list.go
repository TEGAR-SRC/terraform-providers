package iplist

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudIPList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudIPListCreate,
		ReadContext:   resourceOnidelCloudIPListRead,
		UpdateContext: resourceOnidelCloudIPListUpdate,
		DeleteContext: resourceOnidelCloudIPListDelete,
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
				Description:  "IP list name.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Description:  "IP list description.",
			},
			"entries": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "IP addresses or CIDRs.",
			},
			"entry_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of entries.",
			},
			"used_by_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of firewall rules referencing this list.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation timestamp.",
			},
		},
	}
}

type ipListSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EntryCount  int    `json:"entry_count"`
	UsedByCount int    `json:"used_by_count"`
	CreatedAt   string `json:"created_at"`
}

type ipListDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	EntryCount  int    `json:"entry_count"`
	UsedByCount int    `json:"used_by_count"`
	CreatedAt   string `json:"created_at"`
	Entries     []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"entries"`
}

func resourceOnidelCloudIPListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{
		"name": d.Get("name").(string),
	}
	if v, ok := d.GetOk("description"); ok {
		body["description"] = v.(string)
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	var result struct {
		IPList struct {
			ID string `json:"id"`
		} `json:"ip_list"`
	}
	log.Printf("[DEBUG] IP List create: %#v", body)
	err := client.Post("/network/ip_lists", body, &result)
	if err != nil {
		return diag.Errorf("Error creating IP List: %s", err)
	}

	d.SetId(result.IPList.ID)
	log.Printf("[INFO] IP List created: %s", d.Id())

	// Add entries if provided
	if entries, ok := d.GetOk("entries"); ok {
		for _, e := range entries.([]interface{}) {
			entryBody := map[string]interface{}{
				"value": e.(string),
			}
			if v, ok := d.GetOk("team_id"); ok {
				entryBody["team_id"] = v.(string)
			}
			err := client.Post(fmt.Sprintf("/network/ip_lists/%s/entries", d.Id()), entryBody, nil)
			if err != nil {
				return diag.Errorf("Error adding entry %s to IP List: %s", e.(string), err)
			}
		}
	}

	return resourceOnidelCloudIPListRead(ctx, d, meta)
}

func resourceOnidelCloudIPListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		IPList ipListDetail `json:"ip_list"`
	}
	err := client.Get("/network/ip_lists/"+d.Id(), params, &result)
	if err != nil {
		log.Printf("[DEBUG] IP List not found: %s", err)
		d.SetId("")
		return nil
	}

	d.Set("name", result.IPList.Name)
	d.Set("description", result.IPList.Description)
	d.Set("entry_count", result.IPList.EntryCount)
	d.Set("used_by_count", result.IPList.UsedByCount)
	d.Set("created_at", result.IPList.CreatedAt)

	entries := make([]string, len(result.IPList.Entries))
	for i, e := range result.IPList.Entries {
		entries[i] = e.Value
	}
	d.Set("entries", entries)
	return nil
}

func resourceOnidelCloudIPListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	body := map[string]interface{}{}
	if d.HasChange("name") {
		body["name"] = d.Get("name").(string)
	}
	if d.HasChange("description") {
		body["description"] = d.Get("description").(string)
	}
	if v, ok := d.GetOk("team_id"); ok {
		body["team_id"] = v.(string)
	}

	if len(body) > 0 {
		err := client.Patch("/network/ip_lists/"+d.Id(), body, nil)
		if err != nil {
			return diag.Errorf("Error updating IP List: %s", err)
		}
	}

	if d.HasChange("entries") {
		// Delete all existing entries and re-add
		params := map[string]string{}
		if v, ok := d.GetOk("team_id"); ok {
			params["team_id"] = v.(string)
		}

		var detail struct {
			IPList ipListDetail `json:"ip_list"`
		}
		err := client.Get("/network/ip_lists/"+d.Id(), params, &detail)
		if err == nil {
			for _, e := range detail.IPList.Entries {
				_ = client.Delete(fmt.Sprintf("/network/ip_lists/%s/entries/%s", d.Id(), e.ID), params)
			}
		}

		if entries, ok := d.GetOk("entries"); ok {
			for _, e := range entries.([]interface{}) {
				entryBody := map[string]interface{}{
					"value": e.(string),
				}
				if v, ok := d.GetOk("team_id"); ok {
					entryBody["team_id"] = v.(string)
				}
				err := client.Post(fmt.Sprintf("/network/ip_lists/%s/entries", d.Id()), entryBody, nil)
				if err != nil {
					return diag.Errorf("Error adding entry %s: %s", e.(string), err)
				}
			}
		}
	}

	return resourceOnidelCloudIPListRead(ctx, d, meta)
}

func resourceOnidelCloudIPListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	log.Printf("[INFO] Deleting IP List: %s", d.Id())
	err := client.Delete("/network/ip_lists/"+d.Id(), params)
	if err != nil {
		return diag.Errorf("Error deleting IP List: %s", err)
	}

	d.SetId("")
	return nil
}
