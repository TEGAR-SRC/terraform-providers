package onidelcloud

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider

func init() {
	testAccProviders = map[string]*schema.Provider{
		"onidelcloud": Provider(),
	}
}

func testAccProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"onidelcloud": func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("provider validation error: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ONIDEL_API_KEY"); v == "" {
		t.Skip("ONIDEL_API_KEY not set; skipping acceptance test")
	}
}

// --- Data Source Tests ---

func TestAccDataSourceOSTemplate_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOSTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_os_template.all", "templates.#"),
				),
			},
		},
	})
}

const testAccDataSourceOSTemplateConfig = `
data "onidelcloud_os_template" "all" {}
`

func TestAccDataSourceInstanceType_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInstanceTypeConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_instance_type.all", "types.#"),
				),
			},
		},
	})
}

const testAccDataSourceInstanceTypeConfig = `
data "onidelcloud_instance_type" "all" {}
`

func TestAccDataSourceTeams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTeamsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_teams.all", "teams.#"),
				),
			},
		},
	})
}

const testAccDataSourceTeamsConfig = `
data "onidelcloud_teams" "all" {}
`

func TestAccDataSourceSSHKeys_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSSHKeysConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_ssh_keys.all", "ssh_keys.#"),
				),
			},
		},
	})
}

const testAccDataSourceSSHKeysConfig = `
data "onidelcloud_ssh_keys" "all" {}
`

func TestAccDataSourceFirewallGroups_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFirewallGroupsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_firewall_groups.all", "firewall_groups.#"),
				),
			},
		},
	})
}

const testAccDataSourceFirewallGroupsConfig = `
data "onidelcloud_firewall_groups" "all" {}
`

func TestAccDataSourceIPLists_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIPListsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_ip_lists.all", "ip_lists.#"),
				),
			},
		},
	})
}

const testAccDataSourceIPListsConfig = `
data "onidelcloud_ip_lists" "all" {}
`

func TestAccDataSourceVMs_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVMsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_vms.all", "vms.#"),
				),
			},
		},
	})
}

const testAccDataSourceVMsConfig = `
data "onidelcloud_vms" "all" {}
`

func TestAccDataSourceObjectStorageServices_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceObjectStorageServicesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_object_storage_services.all", "services.#"),
				),
			},
		},
	})
}

const testAccDataSourceObjectStorageServicesConfig = `
data "onidelcloud_object_storage_services" "all" {}
`

func TestAccDataSourceStartupScripts_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceStartupScriptsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_startup_scripts.all", "scripts.#"),
				),
			},
		},
	})
}

const testAccDataSourceStartupScriptsConfig = `
data "onidelcloud_startup_scripts" "all" {}
`

func TestAccDataSourceSnapshots_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSnapshotsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_snapshots.all", "snapshots.#"),
				),
			},
		},
	})
}

const testAccDataSourceSnapshotsConfig = `
data "onidelcloud_snapshots" "all" {}
`

func TestAccDataSourceBackups_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceBackupsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_backups.all", "backups.#"),
				),
			},
		},
	})
}

const testAccDataSourceBackupsConfig = `
data "onidelcloud_backups" "all" {}
`

// --- Resource Tests ---

func TestAccResourceSSHKey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSSHKeyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("onidelcloud_ssh_key.test", "id"),
					resource.TestCheckResourceAttr("onidelcloud_ssh_key.test", "name", "tf-test-key"),
				),
			},
		},
	})
}

const testAccResourceSSHKeyConfig = `
resource "onidelcloud_ssh_key" "test" {
  name       = "tf-test-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0h8aR1m2iFq3z4y5kP7vB9nXcJfLp2qRtUvWxYzN8dE6fG3hIjK5mOnPqLsZtArCwDxVfBgHjNyM1oP2rS3uT9vXcJfLp2qRtUvWxYzN8dE6fG3hIjK5mOnPqLsZtArCwDxVfBgHjNyM1oP2rS3uT9vXcJfLp2qRtUvWxYzN8dE6fG3hIjK5mOnPqL tf-test-dummy-key"
}
`

func TestAccResourceVPC_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVPCConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("onidelcloud_vpc.test", "id"),
					resource.TestCheckResourceAttr("onidelcloud_vpc.test", "name", "tf-test-vpc"),
					resource.TestCheckResourceAttr("onidelcloud_vpc.test", "location", "Sydney"),
				),
			},
		},
	})
}

const testAccResourceVPCConfig = `
resource "onidelcloud_vpc" "test" {
  name          = "tf-test-vpc"
  description   = "Terraform acceptance test VPC"
  location      = "Sydney"
  v4_subnet     = "10.99.0.0"
  v4_subnet_mask = "24"
}
`

func TestAccResourceFirewallGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFirewallGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("onidelcloud_firewall_group.test", "id"),
					resource.TestCheckResourceAttr("onidelcloud_firewall_group.test", "description", "tf-test-firewall"),
				),
			},
		},
	})
}

const testAccResourceFirewallGroupConfig = `
resource "onidelcloud_firewall_group" "test" {
  description = "tf-test-firewall"
}
`

func TestAccResourceIPList_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIPListConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("onidelcloud_ip_list.test", "id"),
					resource.TestCheckResourceAttr("onidelcloud_ip_list.test", "name", "tf-test-ip-list"),
				),
			},
		},
	})
}

const testAccResourceIPListConfig = `
resource "onidelcloud_ip_list" "test" {
  name        = "tf-test-ip-list"
  description = "IP list for acceptance test"
}
`

func TestAccResourceStartupScript_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceStartupScriptConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("onidelcloud_startup_script.test", "id"),
					resource.TestCheckResourceAttr("onidelcloud_startup_script.test", "name", "tf-test-script"),
				),
			},
		},
	})
}

const testAccResourceStartupScriptConfig = `
resource "onidelcloud_startup_script" "test" {
  name    = "tf-test-script"
  content = "#!/bin/bash\necho hello"
}
`

func TestAccDataSourceVPCs_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVPCsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.onidelcloud_vpcs.all", "vpcs.#"),
				),
			},
		},
	})
}

const testAccDataSourceVPCsConfig = `
data "onidelcloud_vpcs" "all" {}
`
