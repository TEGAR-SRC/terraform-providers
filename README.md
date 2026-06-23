# Terraform Providers

Monorepo for custom Terraform providers managed by [@tegar](https://github.com/tegar).

## Provider List

| Provider | Status | Description |
|----------|--------|-------------|
| [`digitalocean`](./digitalocean) | Archived | Fork of official `terraform-provider-digitalocean` |
| [`hetzner`](./hetzner) | Fork | Fork of official `terraform-provider-hetzner` |
| [`hostinger`](./hostinger) | Internal | Hostinger / hPanel API provider |
| [`ionoscloud`](./ionoscloud) | Fork | Fork of official `terraform-provider-ionoscloud` |
| [`Juniper`](./Juniper) | Internal | Juniper network device provider |
| [`mikrotik`](./mikrotik) | Internal | MikroTik RouterOS provider |
| [`onidelcloud`](./onidelcloud) | **Active** | Onidel Cloud API provider — 10 resources + 13 data sources |
| [`openstack`](./openstack) | Fork | Fork of official `terraform-provider-openstack` |
| [`proxmox`](./proxmox) | Fork | Fork / custom Proxmox VE provider |
| [`virtualizor`](./virtualizor) | Internal | Virtualizor VPS provider |
| [`vmware`](./vmware) | Fork | Fork of official `terraform-provider-vmware` (vSphere) |

## Repository Structure

```
terraform-providers/
├── README.md              ← You are here
├── digitalocean/          # DigitalOcean provider
├── hetzner/               # Hetzner Cloud provider
├── hostinger/             # Hostinger provider
├── ionoscloud/            # IONOS Cloud provider
├── Juniper/               # Juniper provider
├── mikrotik/              # MikroTik provider
├── onidelcloud/           # Onidel Cloud provider
│   ├── main.go
│   ├── go.mod / go.sum
│   ├── GNUmakefile
│   ├── scripts/
│   └── onidelcloud/
│       ├── provider.go
│       ├── provider_test.go
│       ├── config/config.go
│       ├── sshkey/           # resource + data source
│       ├── vpc/              # resource + data source
│       ├── firewall/         # 2 resource + data source
│       ├── vm/               # resource + data source + action
│       ├── objectstorage/    # resource + data source
│       ├── iplist/           # resource + data source
│       ├── reservedip/       # resource
│       ├── startupscript/    # resource + data source
│       ├── ostemplate/       # data source
│       ├── instancetype/     # data source + price
│       ├── teams/            # data source
│       ├── snapshot/         # data source
│       └── backup/           # data source
├── openstack/             # OpenStack provider
├── proxmox/               # Proxmox provider
├── virtualizor/           # Virtualizor provider
└── vmware/                # VMware vSphere provider
```

### Onidel Cloud Provider

**Resources (10):**

| Resource | Endpoint | Description |
|----------|----------|-------------|
| `onidelcloud_ssh_key` | `POST/GET/DELETE /ssh_keys` | SSH public key management |
| `onidelcloud_vpc` | `POST/GET/PUT/DELETE /network/vpcs` | Virtual Private Cloud |
| `onidelcloud_firewall_group` | `POST/GET/PUT/DELETE /network/firewalls` | Firewall group |
| `onidelcloud_firewall_rule` | `POST/GET/PUT/DELETE /network/firewalls/*/rules` | Firewall rule in a group |
| `onidelcloud_vm` | `POST/GET/PATCH/DELETE /vm` | Virtual Machine |
| `onidelcloud_vm_action` | `POST /vm/*/{stop,reboot,snapshot,enable-bgp,disable-bgp,vnc}` | One-shot VM action |
| `onidelcloud_object_storage` | `POST/GET/DELETE /object-storage` | Object Storage service |
| `onidelcloud_ip_list` | `POST/GET/PUT/DELETE /network/ip_lists` | IP allowlist/blocklist |
| `onidelcloud_reserved_ip` | `POST/GET/PATCH/DELETE /network/reserved_ips` | Reserved IP address |
| `onidelcloud_startup_script` | `POST/GET/PUT/DELETE /startup_scripts` | Startup script |

**Data Sources (13):**

| Data Source | Endpoint | Description |
|-------------|----------|-------------|
| `onidelcloud_os_template` | `GET /os_templates` | Available OS images |
| `onidelcloud_instance_type` | `GET /instance_types` | Available VM instance types |
| `onidelcloud_instance_price` | `GET /instance_price` | VM price calculator |
| `onidelcloud_teams` | `GET /teams` | Team list |
| `onidelcloud_ssh_keys` | `GET /ssh_keys` | SSH keys list |
| `onidelcloud_vpcs` | `GET /network/vpcs` | VPCs list |
| `onidelcloud_firewall_groups` | `GET /network/firewalls` | Firewall groups list |
| `onidelcloud_ip_lists` | `GET /network/ip_lists` | IP lists list |
| `onidelcloud_vms` | `GET /vm` | VMs list |
| `onidelcloud_object_storage_services` | `GET /object-storage` | Object storage services list |
| `onidelcloud_startup_scripts` | `GET /startup_scripts` | Startup scripts list |
| `onidelcloud_snapshots` | `GET /snapshots` | Snapshots list |
| `onidelcloud_backups` | `GET /backups` | Backups list |

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.23+
- [Terraform](https://www.terraform.io/) 1.5+

### Build a provider

```bash
cd <provider>
make build
```

### Use a provider locally

```hcl
terraform {
  required_providers {
    onidelcloud = {
      source = "tegar/onidelcloud"
      version = "0.1.0"
    }
  }
}

provider "onidelcloud" {
  api_key = var.onidel_api_key
}
```

## License

Each provider may have its own license terms. Refer to individual `LICENSE` files within each provider directory.

## Author

**Tegar** — Infrastructure automation
