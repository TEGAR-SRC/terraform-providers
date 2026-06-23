# Terraform Providers

Monorepo untuk penyedia Terraform kustom.
Repo: [github.com/TEGAR-SRC/terraform-providers](https://github.com/TEGAR-SRC/terraform-providers)

## Daftar Provider

| # | Provider | Versi | Status | Deskripsi |
|---|----------|-------|--------|-----------|
| 1 | [`digitalocean`](./digitalocean) | - | Archived | Fork dari official `terraform-provider-digitalocean`. Mengelola Droplet, Spaces, Volumes, DNS, Floating IP, Load Balancer, dan seluruh layanan DigitalOcean. Provider成熟 dengan ribuan pengguna |
| 2 | [`hetzner`](./hetzner) | - | Fork | Fork dari official `terraform-provider-hetzner`. Mengelola Cloud server, dedicated server, SSH keys, networks, volumes, firewall, dan load balancer di Hetzner |
| 3 | [`hostinger`](./hostinger) | - | Internal | Provider untuk Hostinger / hPanel API. Mengelola akun hosting, domain, email bisnis, VPS, SSL certificate, dan database. Dibangun dari reverse-engineer API hPanel |
| 4 | [`ionoscloud`](./ionoscloud) | - | Fork | Fork dari official `terraform-provider-ionoscloud`. Mengelola datacenter virtual, server, volume, NIC, LAN, load balancer, backup, dan resource IONOS Cloud lainnya |
| 5 | [`Juniper`](./Juniper) | - | Internal | Provider untuk perangkat jaringan Juniper (JunOS). Mengelola konfigurasi switch, router, firewall via NETCONF/API. Support interface, VLAN, routing policy, firewall filter |
| 6 | [`mikrotik`](./mikrotik) | - | Internal | Provider untuk MikroTik RouterOS. Mengelola interface ethernet/wireless, firewall filter/NAT, DHCP server/client, routing OSPF/BGP, bridge, queue, dan seluruh fitur RouterBoard |
| 7 | [`onidelcloud`](./onidelcloud) | 0.1.0 | **Active** | Provider untuk Onidel Cloud API. 10 resource + 13 data sources: SSH Key, VPC, Firewall Group & Rule, Virtual Machine + VM Action, Object Storage, IP List, Reserved IP, Startup Script, Snapshot, Backup, OS Template, Instance Type & Price, Teams |
| 8 | [`openstack`](./openstack) | - | Fork | Fork dari official `terraform-provider-openstack`. Mengelola compute (server, flavor, keypair), networking (network, subnet, router, floating IP), storage (volume, object), identity (user, project, role) |
| 9 | [`proxmox`](./proxmox) | - | Fork | Fork dan custom Proxmox VE provider. Mengelola QEMU VM, LXC container, storage pool, cluster resources, backup, firewall, dan user management untuk Proxmox Virtual Environment |
| 10 | [`rustfs`](./rustfs) | - | Internal | Provider untuk RustFS / MinIO-compatible object storage. Mengelola bucket, policy akses, service account, quota storage, user management. Cocok untuk penyimpanan S3-compatible self-hosted |
| 11 | [`virtualizor`](./virtualizor) | - | Internal | Provider untuk Virtualizor VPS. Mengelola virtual server creation/destroy, plan management, ISO library, network interfaces, bandwidth monitoring, dan user management |
| 12 | [`vmware`](./vmware) | - | Fork | Fork dari official `terraform-provider-vmware` (vSphere). Mengelola virtual machine, datastore, network switch, cluster compute resource, content library, tag, permission, dan resource vSphere lainnya |

## Struktur Repository

```
terraform-providers/
|-- README.md              (EN - English)
|-- README-id.md           (ID - Bahasa Indonesia)
|-- scripts/
|   |-- generate-readme.ps1        Generator README otomatis
|   |-- gofmtcheck.sh              Script pengecekan format Go
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
|   |   |   |-- config.go          HTTP API client dengan Bearer auth
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

| Resource | Endpoint API | Deskripsi |
|----------|-------------|-----------|
| `onidelcloud_ssh_key` | `POST /ssh_keys`, `GET /ssh_keys/{id}`, `DELETE /ssh_keys/{id}` | Mengelola SSH public key untuk akses VM. Create, read, delete key. Mendukung multiple key per VM |
| `onidelcloud_vpc` | `POST /network/vpcs`, `GET /network/vpcs/{id}`, `PUT /network/vpcs/{id}`, `DELETE /network/vpcs/{id}` | Virtual Private Cloud. Buat jaringan privat dengan subnet IPv4, attach ke VM untuk private networking antar instance |
| `onidelcloud_firewall_group` | `POST /network/firewalls`, `GET /network/firewalls/{id}`, `PUT /network/firewalls/{id}`, `DELETE /network/firewalls/{id}` | Grup firewall sebagai container untuk firewall rules. Bisa di-attach ke satu atau lebih VM |
| `onidelcloud_firewall_rule` | `POST /network/firewalls/{group_id}/rules`, `GET /network/firewalls/{group_id}/rules/{id}`, `PUT ...`, `DELETE ...` | Aturan firewall dalam grup. Support protocol TCP/UDP/ICMP, port range, subnet CIDR, IP list reference |
| `onidelcloud_vm` | `POST /vm`, `GET /vm/{id}`, `PATCH /vm/{id}`, `DELETE /vm/{id}` | Virtual Machine. Provision VM dengan pilihan OS template, instance type, lokasi, SSH key, VPC, firewall, startup script, IPv6 |
| `onidelcloud_vm_action` | `POST /vm/{id}/{action}` | Satu kali aksi pada VM: stop, reboot, snapshot, enable-bgp, disable-bgp, vnc. Trigger ulang jika input berubah |
| `onidelcloud_object_storage` | `POST /object-storage`, `GET /object-storage/{id}`, `DELETE /object-storage/{id}` | Layanan Object Storage S3-compatible. Buat service + optional bucket awal. Kelola kapasitas penyimpanan |
| `onidelcloud_ip_list` | `POST /network/ip_lists`, `GET /network/ip_lists/{id}`, `PUT /network/ip_lists/{id}`, `DELETE /network/ip_lists/{id}` | Daftar IP untuk allowlist/blocklist firewall. Entri otomatis auto-detect tipe IPv4/IPv6. Bisa direferensi dari firewall rule |
| `onidelcloud_reserved_ip` | `POST /network/reserved_ips`, `GET /network/reserved_ips/{id}`, `PATCH /network/reserved_ips/{id}`, `DELETE ...` | Reserved IP address yang bisa di-attach/detach ke VM. Support attach/detach tanpa destroy resource |
| `onidelcloud_startup_script` | `POST /startup_scripts`, `GET /startup_scripts/{id}`, `PUT /startup_scripts/{id}`, `DELETE /startup_scripts/{id}` | Script startup untuk kustomisasi VM saat pertama boot. Support bash script, max 10 script per team |

### Onidel Cloud -- Data Sources (13)

| Data Source | Endpoint API | Deskripsi |
|-------------|-------------|-----------|
| `onidelcloud_os_template` | `GET /os_templates` | Daftar OS template tersedia: Ubuntu, Debian, CentOS, Windows, dll. Setiap template punya id integer dan family |
| `onidelcloud_instance_type` | `GET /instance_types` | Daftar tipe VM: VHP, VHP Pro, dll. Masing-masing punya max vCPU, RAM, disk, network rate, dan lokasi tersedia |
| `onidelcloud_instance_price` | `GET /instance_price` | Kalkulator harga VM. Input vCPU, RAM, disk, lokasi, tipe instance. Output harga per bulan, kuartal, semester, tahun |
| `onidelcloud_teams` | `GET /teams` | Daftar tim dalam akun. Setiap tim punya id, nama, role (Owner/Admin/Member/Biller) |
| `onidelcloud_ssh_keys` | `GET /ssh_keys` | Daftar SSH public key yang sudah terdaftar. Filter opsional berdasarkan team_id |
| `onidelcloud_vpcs` | `GET /network/vpcs` | Daftar VPC. Filter berdasarkan team_id dan/atau lokasi. Output subnet, status, tanggal buat |
| `onidelcloud_firewall_groups` | `GET /network/firewalls` | Daftar grup firewall. Menampilkan deskripsi, jumlah rule, jumlah instance ter-attach |
| `onidelcloud_ip_lists` | `GET /network/ip_lists` | Daftar IP list. Juga menampilkan limit team untuk jumlah list dan entry per list |
| `onidelcloud_vms` | `GET /vm` | Daftar VM. Informasi: nama, vCPU, RAM, disk, IP, status, billing, firewall ter-attach |
| `onidelcloud_object_storage_services` | `GET /object-storage` | Daftar layanan object storage. Informasi: lokasi, region, endpoint, kapasitas, status |
| `onidelcloud_startup_scripts` | `GET /startup_scripts` | Daftar startup script. Informasi: nama, tanggal buat, tanggal update |
| `onidelcloud_snapshots` | `GET /snapshots` | Daftar snapshot VM. Informasi: nama, deskripsi, ukuran, status (available/pending) |
| `onidelcloud_backups` | `GET /backups` | Daftar backup VM. Informasi: instance asal, ukuran, status (available/pending), tanggal buat |

## Memulai

### Prasyarat

- [Go](https://go.dev/) 1.23+
- [Terraform](https://www.terraform.io/) 1.5+

### Build Provider

```bash
cd <provider>
make build
```

### Gunakan Provider Secara Lokal

Konfigurasi `~/.terraformrc` untuk dev override:

```hcl
provider_installation {
  dev_overrides {
    "tegar/onidelcloud" = "D:/path/to/terraform-providers/onidelcloud"
  }
}
```

Contoh penggunaan:

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
  # ambil template terbaru
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

## Lisensi

Setiap provider memiliki ketentuan lisensi masing-masing. Lihat file LICENSE di direktori masing-masing provider.

## Penulis

**Tegar** -- Otomatisasi infrastruktur
