package onidelcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/backup"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/config"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/firewall"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/instancetype"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/iplist"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/objectstorage"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/ostemplate"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/reservedip"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/snapshot"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/sshkey"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/startupscript"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/teams"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/vm"
	"github.com/tegar/Terraform/onidelcloud/onidelcloud/vpc"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ONIDEL_API_KEY", nil),
				Description: "The API key for Onidel Cloud.",
			},
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ONIDEL_API_ENDPOINT", "https://api.cloud.onidel.com"),
				Description: "The Onidel Cloud API endpoint.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"onidelcloud_ssh_key":        sshkey.ResourceOnidelCloudSSHKey(),
			"onidelcloud_vpc":            vpc.ResourceOnidelCloudVPC(),
			"onidelcloud_firewall_group": firewall.ResourceOnidelCloudFirewallGroup(),
			"onidelcloud_firewall_rule":  firewall.ResourceOnidelCloudFirewallRule(),
			"onidelcloud_vm":             vm.ResourceOnidelCloudVM(),
			"onidelcloud_object_storage": objectstorage.ResourceOnidelCloudObjectStorage(),
			"onidelcloud_ip_list":        iplist.ResourceOnidelCloudIPList(),
			"onidelcloud_reserved_ip":    reservedip.ResourceOnidelCloudReservedIP(),
			"onidelcloud_startup_script": startupscript.ResourceOnidelCloudStartupScript(),
			"onidelcloud_vm_action":     vm.ResourceOnidelCloudVMAction(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"onidelcloud_os_template":              ostemplate.DataSourceOnidelCloudOSTemplate(),
			"onidelcloud_instance_type":            instancetype.DataSourceOnidelCloudInstanceType(),
			"onidelcloud_instance_price":           instancetype.DataSourceOnidelCloudInstancePrice(),
			"onidelcloud_teams":                    teams.DataSourceOnidelCloudTeams(),
			"onidelcloud_ssh_keys":                 sshkey.DataSourceOnidelCloudSSHKeys(),
			"onidelcloud_vpcs":                     vpc.DataSourceOnidelCloudVPCs(),
			"onidelcloud_firewall_groups":          firewall.DataSourceOnidelCloudFirewallGroups(),
			"onidelcloud_ip_lists":                 iplist.DataSourceOnidelCloudIPLists(),
			"onidelcloud_vms":                      vm.DataSourceOnidelCloudVMs(),
			"onidelcloud_object_storage_services":  objectstorage.DataSourceOnidelCloudObjectStorageServices(),
			"onidelcloud_startup_scripts":           startupscript.DataSourceOnidelCloudStartupScripts(),
			"onidelcloud_snapshots":                snapshot.DataSourceOnidelCloudSnapshots(),
			"onidelcloud_backups":                  backup.DataSourceOnidelCloudBackups(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	conf := &config.Config{
		APIKey:      d.Get("api_key").(string),
		APIEndpoint: d.Get("api_endpoint").(string),
	}
	client, err := conf.Client()
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return client, nil
}
