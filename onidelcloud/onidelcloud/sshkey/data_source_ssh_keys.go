package sshkey

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
)

func DataSourceOnidelCloudSSHKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOnidelCloudSSHKeysRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Team ID to filter SSH keys.",
			},
			"ssh_keys": {
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
						"public_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOnidelCloudSSHKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*config.Client)

	params := map[string]string{}
	if v, ok := d.GetOk("team_id"); ok {
		params["team_id"] = v.(string)
	}

	var result struct {
		SSHKeys []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			SSHKey  string `json:"ssh_key"`
			Created string `json:"created"`
		} `json:"ssh_keys"`
	}
	err := client.Get("/ssh_keys", params, &result)
	if err != nil {
		return diag.Errorf("Error listing SSH keys: %s", err)
	}

	keys := make([]interface{}, len(result.SSHKeys))
	for i, k := range result.SSHKeys {
		keys[i] = map[string]interface{}{
			"id":         k.ID,
			"name":       k.Name,
			"public_key": k.SSHKey,
			"created":    k.Created,
		}
	}
	d.Set("ssh_keys", keys)
	d.SetId("ssh_keys")

	return nil
}
