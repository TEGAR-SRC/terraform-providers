# Terraform Providers

Monorepo for custom Terraform providers.
Repo: [github.com/TEGAR-SRC/terraform-providers](https://github.com/TEGAR-SRC/terraform-providers)

## Provider List

| # | Provider | Version | Status | Description |
|---|----------|---------|--------|-------------|
| 1 | [`digitalocean`](./digitalocean) | - | Archived | Fork of official `terraform-provider-digitalocean`. Manage Droplets, Spaces, Volumes, DNS, Floating IPs, Load Balancers, and all DigitalOcean services. Mature provider used by thousands |
| 2 | [`hetzner`](./hetzner) | - | Fork | Fork of official `terraform-provider-hetzner`. Manage Cloud servers, dedicated servers, SSH keys, networks, volumes, firewalls, and load balancers on Hetzner |
| 3 | [`hostinger`](./hostinger) | - | Internal | Provider for Hostinger / hPanel API. Manage hosting accounts, domains, business email, VPS, SSL certificates, and databases. Built from reverse-engineered hPanel API |
| 4 | [`ionoscloud`](./ionoscloud) | - | Fork | Fork of official `terraform-provider-ionoscloud`. Manage virtual datacenters, servers, volumes, NICs, LANs, load balancers, backups, and other IONOS Cloud resources |
| 5 | [`Juniper`](./Juniper) | - | Internal | Provider for Juniper network devices (JunOS). Manage switch, router, and firewall configurations via NETCONF/API. Supports interfaces, VLANs, routing policy, firewall filters |
| 6 | [`mikrotik`](./mikrotik) | - | Internal | Provider for MikroTik RouterOS. Manage ethernet/wireless interfaces, firewall filter/NAT, DHCP server/client, OSPF/BGP routing, bridges, queues, and all RouterBoard features |
| 7 | [`onidelcloud`](./onidelcloud) | 0.1.0 | **Active** | Provider for Onidel Cloud API. 10 resources + 13 data sources: SSH Key, VPC, Firewall Group & Rule, Virtual Machine + VM Action, Object Storage, IP List, Reserved IP, Startup Script, Snapshot, Backup, OS Template, Instance Type & Price, Teams |
| 8 | [`openstack`](./openstack) | - | Fork | Fork of official `terraform-provider-openstack`. Manage compute (servers, flavors, keypairs), networking (networks, subnets, routers, floating IPs), storage (volumes, objects), identity (users, projects, roles) |
| 9 | [`proxmox`](./proxmox) | - | Fork | Fork and custom Proxmox VE provider. Manage QEMU VMs, LXC containers, storage pools, cluster resources, backups, firewalls, and user management for Proxmox Virtual Environment |
| 10 | [`rustfs`](./rustfs) | - | Internal | Provider for RustFS / MinIO-compatible object storage. Manage buckets, access policies, service accounts, storage quotas, and users. Suitable for self-hosted S3-compatible storage |
| 11 | [`virtualizor`](./virtualizor) | - | Internal | Provider for Virtualizor VPS. Manage virtual server creation/destroy, plan management, ISO library, network interfaces, bandwidth monitoring, and user management |
| 12 | [`vmware`](./vmware) | - | Fork | Fork of official `terraform-provider-vmware` (vSphere). Manage virtual machines, datastores, network switches, cluster compute resources, content libraries, tags, permissions, and other vSphere resources |

## Repository Structure

```
terraform-providers/
|-- README.md              (EN - English)
|-- README-id.md           (ID - Bahasa Indonesia)
|-- scripts/
|   |-- generate-readme.ps1        Automated README generator
|   |-- gofmtcheck.sh              Go formatting check script
|-- digitalocean/                  DigitalOcean provider
|-- hetzner/                       Hetzner Cloud provider
|-- hostinger/                     Hostinger provider
|-- ionoscloud/                    IONOS Cloud provider
|-- Juniper/                       Juniper provider
|-- mikrotik/                      MikroTik provider
|-- onidelcloud/                   Onidel Cloud provider
|   |-- main.go
|   |-- go.mod / go.sum
|   |-- GNUmakefile
|   |-- scripts/
|   |-- onidelcloud/
|   |   |-- provider.go
|   |   |-- provider_test.go
|   |   |-- config/
|   |   |   |-- config.go          HTTP API client with Bearer auth
|   |   |-- sshkey/                Resource: ssh_key | DS: ssh_keys
|   |   |-- vpc/                   Resource: vpc | DS: vpcs
|   |   |-- firewall/              Resource: firewall_group, firewall_rule | DS: firewall_groups
|   |   |-- vm/                    Resource: vm, vm_action | DS: vms
|   |   |-- objectstorage/         Resource: object_storage | DS: object_storage_services
|   |   |-- iplist/                Resource: ip_list | DS: ip_lists
|   |   |-- reservedip/            Resource: reserved_ip
|   |   |-- startupscript/         Resource: startup_script | DS: startup_scripts
|   |   |-- ostemplate/            DS: os_template
|   |   |-- instancetype/          DS: instance_type, instance_price
|   |   |-- teams/                 DS: teams
|   |   |-- snapshot/              DS: snapshots
|   |   |-- backup/                DS: backups
|-- openstack/                     OpenStack provider
|-- proxmox/                       Proxmox provider
|-- rustfs/                        RustFS / MinIO provider
|   |-- main.go
|   |-- go.mod / go.sum
|   |-- GNUmakefile
|   |-- provider/
|   |   |-- provider.go
|   |   |-- all_client.go
|   |   |-- helper.go
|   |   |-- rustfs_bucket_ressource.go
|   |   |-- rustfs_policy_ressource.go
|   |   |-- rustfs_quota_ressource.go
|   |   |-- rustfs_service_account_ressource.go
|   |   |-- rustfs_user_resource.go
|   |   |-- *_test.go
|   |-- pkg/rustfs/
|   |   |-- admin_client.go
|   |   |-- bucket.go
|   |   |-- policy.go
|   |   |-- quota.go
|   |   |-- service_account.go
|   |   |-- user_account.go
|   |   |-- *_test.go
|   |-- docs/
|   |   |-- index.md
|   |   |-- resources/
|   |-- examples/
|-- virtualizor/                   Virtualizor provider
|-- vmware/                        VMware vSphere provider
```

### Onidel Cloud -- Resources (10)

| Resource | API Endpoint | Description |
|----------|-------------|-------------|
| `onidelcloud_ssh_key` | `POST /ssh_keys`, `GET /ssh_keys/{id}`, `DELETE /ssh_keys/{id}` | Manage SSH public keys for VM access. Create, read, delete keys. Supports multiple keys per VM |
| `onidelcloud_vpc` | `POST /network/vpcs`, `GET /network/vpcs/{id}`, `PUT /network/vpcs/{id}`, `DELETE /network/vpcs/{id}` | Virtual Private Cloud. Create private networks with IPv4 subnet, attach to VMs for private inter-instance networking |
| `onidelcloud_firewall_group` | `POST /network/firewalls`, `GET /network/firewalls/{id}`, `PUT /network/firewalls/{id}`, `DELETE /network/firewalls/{id}` | Firewall group as a container for firewall rules. Can be attached to one or more VMs |
| `onidelcloud_firewall_rule` | `POST /network/firewalls/{group_id}/rules`, `GET /network/firewalls/{group_id}/rules/{id}`, `PUT ...`, `DELETE ...` | Firewall rule within a group. Supports TCP/UDP/ICMP protocols, port ranges, subnet CIDR, IP list references |
| `onidelcloud_vm` | `POST /vm`, `GET /vm/{id}`, `PATCH /vm/{id}`, `DELETE /vm/{id}` | Virtual Machine. Provision VMs with OS template selection, instance type, location, SSH keys, VPC, firewall, startup script, IPv6 support |
| `onidelcloud_vm_action` | `POST /vm/{id}/{action}` | One-shot VM actions: stop, reboot, snapshot, enable-bgp, disable-bgp, vnc. Triggers again when input changes |
| `onidelcloud_object_storage` | `POST /object-storage`, `GET /object-storage/{id}`, `DELETE /object-storage/{id}` | S3-compatible Object Storage service. Create service with optional initial bucket. Manage storage capacity |
| `onidelcloud_ip_list` | `POST /network/ip_lists`, `GET /network/ip_lists/{id}`, `PUT /network/ip_lists/{id}`, `DELETE /network/ip_lists/{id}` | IP allowlist/blocklist for firewall. Entries auto-detect IPv4/IPv6 type. Referenceable from firewall rules |
| `onidelcloud_reserved_ip` | `POST /network/reserved_ips`, `GET /network/reserved_ips/{id}`, `PATCH /network/reserved_ips/{id}`, `DELETE ...` | Reserved IP address attachable/detachable to VMs. Supports attach/detach without destroying the resource |
| `onidelcloud_startup_script` | `POST /startup_scripts`, `GET /startup_scripts/{id}`, `PUT /startup_scripts/{id}`, `DELETE /startup_scripts/{id}` | Startup scripts for VM first-boot customization. Supports bash scripts, max 10 scripts per team |

### Onidel Cloud -- Data Sources (13)

| Data Source | API Endpoint | Description |
|-------------|-------------|-------------|
| `onidelcloud_os_template` | `GET /os_templates` | Available OS templates: Ubuntu, Debian, CentOS, Windows, etc. Each template has an integer ID and family name |
| `onidelcloud_instance_type` | `GET /instance_types` | Available VM types: VHP, VHP Pro, etc. Each has max vCPU, RAM, disk, network rate, and available locations |
| `onidelcloud_instance_price` | `GET /instance_price` | VM price calculator. Input: vCPU, RAM, disk, location, instance type. Output: monthly, quarterly, semiannual, annual pricing |
| `onidelcloud_teams` | `GET /teams` | Team list within an account. Each team has an ID, name, and role (Owner/Admin/Member/Biller) |
| `onidelcloud_ssh_keys` | `GET /ssh_keys` | List of registered SSH public keys. Optional team_id filter |
| `onidelcloud_vpcs` | `GET /network/vpcs` | List of VPCs. Filter by team_id and/or location. Shows subnet, status, creation date |
| `onidelcloud_firewall_groups` | `GET /network/firewalls` | List of firewall groups. Shows description, rule count, attached instance count |
| `onidelcloud_ip_lists` | `GET /network/ip_lists` | List of IP lists. Also shows team limits for total lists and entries per list |
| `onidelcloud_vms` | `GET /vm` | List of VMs. Shows: name, vCPU, RAM, disk, IP addresses, status, billing info, attached firewall |
| `onidelcloud_object_storage_services` | `GET /object-storage` | List of object storage services. Shows: location, region, endpoint, capacity, status |
| `onidelcloud_startup_scripts` | `GET /startup_scripts` | List of startup scripts. Shows: name, creation date, last update date |
| `onidelcloud_snapshots` | `GET /snapshots` | List of VM snapshots. Shows: name, description, size, status (available/pending) |
| `onidelcloud_backups` | `GET /backups` | List of VM backups. Shows: source instance, size, status (available/pending), creation date |

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.23+
- [Terraform](https://www.terraform.io/) 1.5+

### Build a Provider

```bash
cd <provider>
make build
```

### Use a Provider Locally

Configure `~/.terraformrc` for dev override:

```hcl
provider_installation {
  dev_overrides {
    "tegar/onidelcloud" = "D:/path/to/terraform-providers/onidelcloud"
  }
}
```

Example usage:

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

data "onidelcloud_os_template" "ubuntu" {
  # fetch latest template
}

resource "onidelcloud_vm" "web" {
  name          = "web-server-01"
  instance_type = data.onidelcloud_instance_type.standard.id
  location      = "Sydney"
  os            = data.onidelcloud_os_template.ubuntu.templates[0].id
}
```

### Acceptance Test

```bash
cd onidelcloud
export ONIDEL_API_KEY="your-api-key"
TF_ACC=1 go test ./... -v -run TestAcc
```

## License

Each provider may have its own license terms. Refer to individual LICENSE files within each provider directory.

## Author

**Tegar** -- Infrastructure automation
