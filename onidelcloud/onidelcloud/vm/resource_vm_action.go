package vm

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func ResourceOnidelCloudVMAction() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOnidelCloudVMActionCreate,
		ReadContext:   resourceOnidelCloudVMActionRead,
		DeleteContext: resourceOnidelCloudVMActionDelete,
		Schema: map[string]*schema.Schema{
			"vm_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VM UUID to perform action on.",
			},
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Action to perform: stop, reboot, snapshot, enable-bgp, disable-bgp, vnc.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Snapshot name (required when action=snapshot).",
			},
			"desc": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Snapshot description (optional when action=snapshot).",
			},
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Team ID. Default team will be used if not provided.",
			},
			"vnc_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VNC session URL (only when action=vnc).",
			},
			"expire_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "VNC session expiry timestamp (only when action=vnc).",
			},
		},
	}
}

func resourceOnidelCloudVMActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	vmID := d.Get("vm_id").(string)
	action := d.Get("action").(string)

	if action == "vnc" {
		body := map[string]interface{}{}
		if v, ok := d.GetOk("team_id"); ok {
			body["team_id"] = v.(string)
		}

		var result struct {
			VNCURL   string `json:"vnc_url"`
			ExpireAt int    `json:"expire_at"`
		}
		err := client.Post(fmt.Sprintf("/vm/%s/vnc", vmID), body, &result)
		if err != nil {
			return diag.Errorf("Error creating VNC session: %s", err)
		}

		d.Set("vnc_url", result.VNCURL)
		d.Set("expire_at", result.ExpireAt)
	} else {
		body := map[string]interface{}{}
		if v, ok := d.GetOk("team_id"); ok {
			body["team_id"] = v.(string)
		}
		if action == "snapshot" {
			if v, ok := d.GetOk("name"); ok {
				body["name"] = v.(string)
			}
			if v, ok := d.GetOk("desc"); ok {
				body["desc"] = v.(string)
			}
		}

		err := client.Post(fmt.Sprintf("/vm/%s/%s", vmID, action), body, nil)
		if err != nil {
			return diag.Errorf("Error performing VM action %s: %s", action, err)
		}
	}

	d.SetId(fmt.Sprintf("%s/%s", vmID, action))
	return nil
}

func resourceOnidelCloudVMActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceOnidelCloudVMActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
