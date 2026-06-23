$providers = @(
    @{id=1; dir="digitalocean"; name="digitalocean"; ver="-"; status="Archived"; en="Fork of official terraform-provider-digitalocean. Manage Droplets, Spaces, Volumes, DNS, Floating IPs, Load Balancers, and all DigitalOcean services. Mature provider used by thousands"; id="Fork dari official terraform-provider-digitalocean. Mengelola Droplet, Spaces, Volumes, DNS, Floating IP, Load Balancer, dan seluruh layanan DigitalOcean. Provider mature dengan ribuan pengguna"}
    @{id=2; dir="hetzner"; name="hetzner"; ver="-"; status="Fork"; en="Fork of official terraform-provider-hetzner. Manage Cloud servers, dedicated servers, SSH keys, networks, volumes, firewalls, and load balancers on Hetzner"; id="Fork dari official terraform-provider-hetzner. Mengelola Cloud server, dedicated server, SSH keys, networks, volumes, firewall, dan load balancer di Hetzner"}
    @{id=3; dir="hostinger"; name="hostinger"; ver="-"; status="Internal"; en="Provider for Hostinger / hPanel API. Manage hosting accounts, domains, business email, VPS, SSL certificates, and databases. Built from reverse-engineered hPanel API"; id="Provider untuk Hostinger / hPanel API. Mengelola akun hosting, domain, email bisnis, VPS, SSL certificate, dan database. Dibangun dari reverse-engineer API hPanel"}
    @{id=4; dir="ionoscloud"; name="ionoscloud"; ver="-"; status="Fork"; en="Fork of official terraform-provider-ionoscloud. Manage virtual datacenters, servers, volumes, NICs, LANs, load balancers, backups, and other IONOS Cloud resources"; id="Fork dari official terraform-provider-ionoscloud. Mengelola datacenter virtual, server, volume, NIC, LAN, load balancer, backup, dan resource IONOS Cloud lainnya"}
    @{id=5; dir="Juniper"; name="Juniper"; ver="-"; status="Internal"; en="Provider for Juniper network devices (JunOS). Manage switch, router, and firewall configurations via NETCONF/API. Supports interfaces, VLANs, routing policy, firewall filters"; id="Provider untuk perangkat jaringan Juniper (JunOS). Mengelola konfigurasi switch, router, firewall via NETCONF/API. Support interface, VLAN, routing policy, firewall filter"}
    @{id=6; dir="mikrotik"; name="mikrotik"; ver="-"; status="Internal"; en="Provider for MikroTik RouterOS. Manage ethernet/wireless interfaces, firewall filter/NAT, DHCP server/client, OSPF/BGP routing, bridges, queues, and all RouterBoard features"; id="Provider untuk MikroTik RouterOS. Mengelola interface ethernet/wireless, firewall filter/NAT, DHCP server/client, routing OSPF/BGP, bridge, queue, dan seluruh fitur RouterBoard"}
    @{id=7; dir="onidelcloud"; name="onidelcloud"; ver="0.1.0"; status="**Active**"; en="Provider for Onidel Cloud API. 10 resources + 13 data sources: SSH Key, VPC, Firewall Group & Rule, Virtual Machine + VM Action, Object Storage, IP List, Reserved IP, Startup Script, Snapshot, Backup, OS Template, Instance Type & Price, Teams"; id="Provider untuk Onidel Cloud API. 10 resource + 13 data sources: SSH Key, VPC, Firewall Group & Rule, Virtual Machine + VM Action, Object Storage, IP List, Reserved IP, Startup Script, Snapshot, Backup, OS Template, Instance Type & Price, Teams"}
    @{id=8; dir="openstack"; name="openstack"; ver="-"; status="Fork"; en="Fork of official terraform-provider-openstack. Manage compute (servers, flavors, keypairs), networking (networks, subnets, routers, floating IPs), storage (volumes, objects), identity (users, projects, roles)"; id="Fork dari official terraform-provider-openstack. Mengelola compute (server, flavor, keypair), networking (network, subnet, router, floating IP), storage (volume, object), identity (user, project, role)"}
    @{id=9; dir="proxmox"; name="proxmox"; ver="-"; status="Fork"; en="Fork and custom Proxmox VE provider. Manage QEMU VMs, LXC containers, storage pools, cluster resources, backups, firewalls, and user management for Proxmox Virtual Environment"; id="Fork dan custom Proxmox VE provider. Mengelola QEMU VM, LXC container, storage pool, cluster resources, backup, firewall, dan user management untuk Proxmox Virtual Environment"}
    @{id=10; dir="rustfs"; name="rustfs"; ver="-"; status="Internal"; en="Provider for RustFS / MinIO-compatible object storage. Manage buckets, access policies, service accounts, storage quotas, and users. Suitable for self-hosted S3-compatible storage"; id="Provider untuk RustFS / MinIO-compatible object storage. Mengelola bucket, policy akses, service account, quota storage, user management. Cocok untuk penyimpanan S3-compatible self-hosted"}
    @{id=11; dir="virtualizor"; name="virtualizor"; ver="-"; status="Internal"; en="Provider for Virtualizor VPS. Manage virtual server creation/destroy, plan management, ISO library, network interfaces, bandwidth monitoring, and user management"; id="Provider untuk Virtualizor VPS. Mengelola virtual server creation/destroy, plan management, ISO library, network interfaces, bandwidth monitoring, dan user management"}
    @{id=12; dir="vmware"; name="vmware"; ver="-"; status="Fork"; en="Fork of official terraform-provider-vmware (vSphere). Manage virtual machines, datastores, network switches, cluster compute resources, content libraries, tags, permissions, and other vSphere resources"; id="Fork dari official terraform-provider-vmware (vSphere). Mengelola virtual machine, datastore, network switch, cluster compute resource, content library, tag, permission, dan resource vSphere lainnya"}
)

$onidelResources = @(
    @{name="onidelcloud_ssh_key"; endpoint="POST /ssh_keys, GET /ssh_keys/{id}, DELETE /ssh_keys/{id}"; en="Manage SSH public keys for VM access. Create, read, delete keys. Supports multiple keys per VM"; id="Mengelola SSH public key untuk akses VM. Create, read, delete key. Mendukung multiple key per VM"}
    @{name="onidelcloud_vpc"; endpoint="POST /network/vpcs, GET /network/vpcs/{id}, PUT /network/vpcs/{id}, DELETE /network/vpcs/{id}"; en="Virtual Private Cloud. Create private networks with IPv4 subnet, attach to VMs for private inter-instance networking"; id="Virtual Private Cloud. Buat jaringan privat dengan subnet IPv4, attach ke VM untuk private networking antar instance"}
    @{name="onidelcloud_firewall_group"; endpoint="POST /network/firewalls, GET /network/firewalls/{id}, PUT /network/firewalls/{id}, DELETE /network/firewalls/{id}"; en="Firewall group as a container for firewall rules. Can be attached to one or more VMs"; id="Grup firewall sebagai container untuk firewall rules. Bisa di-attach ke satu atau lebih VM"}
    @{name="onidelcloud_firewall_rule"; endpoint="POST /network/firewalls/{group_id}/rules, GET /.../{id}, PUT /.../{id}, DELETE /.../{id}"; en="Firewall rule within a group. Supports TCP/UDP/ICMP protocols, port ranges, subnet CIDR, IP list references"; id="Aturan firewall dalam grup. Support protocol TCP/UDP/ICMP, port range, subnet CIDR, IP list reference"}
    @{name="onidelcloud_vm"; endpoint="POST /vm, GET /vm/{id}, PATCH /vm/{id}, DELETE /vm/{id}"; en="Virtual Machine. Provision VMs with OS template, instance type, location, SSH keys, VPC, firewall, startup script, IPv6"; id="Virtual Machine. Provision VM dengan pilihan OS template, instance type, lokasi, SSH key, VPC, firewall, startup script, IPv6"}
    @{name="onidelcloud_vm_action"; endpoint="POST /vm/{id}/{action}"; en="One-shot VM actions: stop, reboot, snapshot, enable-bgp, disable-bgp, vnc. Triggers again when input changes"; id="Satu kali aksi pada VM: stop, reboot, snapshot, enable-bgp, disable-bgp, vnc. Trigger ulang jika input berubah"}
    @{name="onidelcloud_object_storage"; endpoint="POST /object-storage, GET /object-storage/{id}, DELETE /object-storage/{id}"; en="S3-compatible Object Storage service. Create service with optional initial bucket. Manage storage capacity"; id="Layanan Object Storage S3-compatible. Buat service + optional bucket awal. Kelola kapasitas penyimpanan"}
    @{name="onidelcloud_ip_list"; endpoint="POST /network/ip_lists, GET /network/ip_lists/{id}, PUT /network/ip_lists/{id}, DELETE /network/ip_lists/{id}"; en="IP allowlist/blocklist for firewall. Entries auto-detect IPv4/IPv6 type. Referenceable from firewall rules"; id="Daftar IP untuk allowlist/blocklist firewall. Entri otomatis auto-detect tipe IPv4/IPv6. Bisa direferensi dari firewall rule"}
    @{name="onidelcloud_reserved_ip"; endpoint="POST /network/reserved_ips, GET /.../{id}, PATCH /.../{id}, DELETE /.../{id}"; en="Reserved IP address attachable/detachable to VMs. Supports attach/detach without destroying the resource"; id="Reserved IP address yang bisa di-attach/detach ke VM. Support attach/detach tanpa destroy resource"}
    @{name="onidelcloud_startup_script"; endpoint="POST /startup_scripts, GET /.../{id}, PUT /.../{id}, DELETE /.../{id}"; en="Startup scripts for VM first-boot customization. Supports bash scripts, max 10 scripts per team"; id="Script startup untuk kustomisasi VM saat pertama boot. Support bash script, max 10 script per team"}
)

$onidelDataSources = @(
    @{name="onidelcloud_os_template"; endpoint="GET /os_templates"; en="Available OS templates: Ubuntu, Debian, CentOS, Windows, etc. Each template has integer ID and family name"; id="Daftar OS template tersedia: Ubuntu, Debian, CentOS, Windows, dll. Setiap template punya id integer dan family"}
    @{name="onidelcloud_instance_type"; endpoint="GET /instance_types"; en="Available VM types: VHP, VHP Pro, etc. Each has max vCPU, RAM, disk, network rate, and available locations"; id="Daftar tipe VM: VHP, VHP Pro, dll. Masing-masing punya max vCPU, RAM, disk, network rate, dan lokasi tersedia"}
    @{name="onidelcloud_instance_price"; endpoint="GET /instance_price"; en="VM price calculator. Input vCPU, RAM, disk, location, instance type. Output monthly, quarterly, semiannual, annual"; id="Kalkulator harga VM. Input vCPU, RAM, disk, lokasi, tipe instance. Output harga per bulan, kuartal, semester, tahun"}
    @{name="onidelcloud_teams"; endpoint="GET /teams"; en="Team list within an account. Each team has ID, name, role (Owner/Admin/Member/Biller)"; id="Daftar tim dalam akun. Setiap tim punya id, nama, role (Owner/Admin/Member/Biller)"}
    @{name="onidelcloud_ssh_keys"; endpoint="GET /ssh_keys"; en="List of registered SSH public keys. Optional team_id filter"; id="Daftar SSH public key yang sudah terdaftar. Filter opsional berdasarkan team_id"}
    @{name="onidelcloud_vpcs"; endpoint="GET /network/vpcs"; en="List of VPCs. Filter by team_id and/or location. Shows subnet, status, creation date"; id="Daftar VPC. Filter berdasarkan team_id dan/atau lokasi. Output subnet, status, tanggal buat"}
    @{name="onidelcloud_firewall_groups"; endpoint="GET /network/firewalls"; en="List of firewall groups. Shows description, rule count, attached instance count"; id="Daftar grup firewall. Menampilkan deskripsi, jumlah rule, jumlah instance ter-attach"}
    @{name="onidelcloud_ip_lists"; endpoint="GET /network/ip_lists"; en="List of IP lists. Also shows team limits for total lists and entries per list"; id="Daftar IP list. Juga menampilkan limit team untuk jumlah list dan entry per list"}
    @{name="onidelcloud_vms"; endpoint="GET /vm"; en="List of VMs. Shows name, vCPU, RAM, disk, IP, status, billing info, attached firewall"; id="Daftar VM. Informasi: nama, vCPU, RAM, disk, IP, status, billing, firewall ter-attach"}
    @{name="onidelcloud_object_storage_services"; endpoint="GET /object-storage"; en="List of object storage services. Shows location, region, endpoint, capacity, status"; id="Daftar layanan object storage. Informasi: lokasi, region, endpoint, kapasitas, status"}
    @{name="onidelcloud_startup_scripts"; endpoint="GET /startup_scripts"; en="List of startup scripts. Shows name, creation date, last update date"; id="Daftar startup script. Informasi: nama, tanggal buat, tanggal update"}
    @{name="onidelcloud_snapshots"; endpoint="GET /snapshots"; en="List of VM snapshots. Shows name, description, size, status (available/pending)"; id="Daftar snapshot VM. Informasi: nama, deskripsi, ukuran, status (available/pending)"}
    @{name="onidelcloud_backups"; endpoint="GET /backups"; en="List of VM backups. Shows source instance, size, status, creation date"; id="Daftar backup VM. Informasi: instance asal, ukuran, status, tanggal buat"}
)

function Generate-Readme($lang) {
    $isID = ($lang -eq "id")
    $lines = @()

    if ($isID) {
        $lines += "# Penyedia Terraform"
        $lines += ""
        $lines += "Monorepo untuk penyedia Terraform kustom."
    } else {
        $lines += "# Terraform Providers"
        $lines += ""
        $lines += "Monorepo for custom Terraform providers."
    }
    $lines += "Repo: [github.com/TEGAR-SRC/terraform-providers](https://github.com/TEGAR-SRC/terraform-providers)"
    $lines += ""

    if ($isID) { $lines += "## Daftar Provider" } else { $lines += "## Provider List" }
    $lines += ""
    if ($isID) { $lines += "| # | Provider | Versi | Status | Deskripsi |" } else { $lines += "| # | Provider | Version | Status | Description |" }
    $lines += "|---|----------|-------|--------|-----------|"

    foreach ($p in $providers) {
        $desc = if ($isID) { $p.id } else { $p.en }
        $lines += "| $($p.id) | [\`$($p.name)\`](./$($p.dir)) | $($p.ver) | $($p.status) | $desc |"
    }

    $lines += ""
    if ($isID) { $lines += "## Struktur Repository" } else { $lines += "## Repository Structure" }
    $lines += ""
    $lines += '```'
    $lines += "terraform-providers/"
    $lines += "|-- README.md              (EN - English)"
    $lines += "|-- README-id.md           (ID - Bahasa Indonesia)"
    $lines += "|-- scripts/"
    $lines += "|   |-- generate-readme.ps1        Automated README generator"
    $lines += "|   |-- gofmtcheck.sh              Go formatting check script"
    $lines += "|-- digitalocean/                  DigitalOcean provider"
    $lines += "|-- hetzner/                       Hetzner Cloud provider"
    $lines += "|-- hostinger/                     Hostinger provider"
    $lines += "|-- ionoscloud/                    IONOS Cloud provider"
    $lines += "|-- Juniper/                       Juniper provider"
    $lines += "|-- mikrotik/                      MikroTik provider"
    $lines += "|-- onidelcloud/                   Onidel Cloud provider"
    $lines += "|   |-- main.go"
    $lines += "|   |-- go.mod / go.sum"
    $lines += "|   |-- GNUmakefile"
    $lines += "|   |-- scripts/"
    $lines += "|   |-- onidelcloud/"
    $lines += "|   |   |-- provider.go"
    $lines += "|   |   |-- provider_test.go"
    $lines += "|   |   |-- config/config.go"
    $lines += "|   |   |-- sshkey/                Resource: ssh_key | DS: ssh_keys"
    $lines += "|   |   |-- vpc/                   Resource: vpc | DS: vpcs"
    $lines += "|   |   |-- firewall/              Resource: firewall_group, firewall_rule | DS: firewall_groups"
    $lines += "|   |   |-- vm/                    Resource: vm, vm_action | DS: vms"
    $lines += "|   |   |-- objectstorage/         Resource: object_storage | DS: object_storage_services"
    $lines += "|   |   |-- iplist/                Resource: ip_list | DS: ip_lists"
    $lines += "|   |   |-- reservedip/            Resource: reserved_ip"
    $lines += "|   |   |-- startupscript/         Resource: startup_script | DS: startup_scripts"
    $lines += "|   |   |-- ostemplate/            DS: os_template"
    $lines += "|   |   |-- instancetype/          DS: instance_type, instance_price"
    $lines += "|   |   |-- teams/                 DS: teams"
    $lines += "|   |   |-- snapshot/              DS: snapshots"
    $lines += "|   |   |-- backup/                DS: backups"
    $lines += "|-- openstack/                     OpenStack provider"
    $lines += "|-- proxmox/                       Proxmox provider"
    $lines += "|-- rustfs/                        RustFS / MinIO provider"
    $lines += "|   |-- main.go"
    $lines += "|   |-- go.mod / go.sum"
    $lines += "|   |-- GNUmakefile"
    $lines += "|   |-- provider/"
    $lines += "|   |   |-- provider.go, all_client.go, helper.go"
    $lines += "|   |   |-- rustfs_bucket_ressource.go"
    $lines += "|   |   |-- rustfs_policy_ressource.go"
    $lines += "|   |   |-- rustfs_quota_ressource.go"
    $lines += "|   |   |-- rustfs_service_account_ressource.go"
    $lines += "|   |   |-- rustfs_user_resource.go"
    $lines += "|   |   |-- *_test.go"
    $lines += "|   |-- pkg/rustfs/"
    $lines += "|   |   |-- admin_client.go"
    $lines += "|   |   |-- bucket.go, policy.go, quota.go"
    $lines += "|   |   |-- service_account.go, user_account.go"
    $lines += "|   |   |-- *_test.go"
    $lines += "|   |-- docs/"
    $lines += "|   |-- examples/"
    $lines += "|-- virtualizor/                   Virtualizor provider"
    $lines += "|-- vmware/                        VMware vSphere provider"
    $lines += '```'
    $lines += ""

    if ($isID) {
        $lines += "### Onidel Cloud -- Resources (10)"
    } else {
        $lines += "### Onidel Cloud -- Resources (10)"
    }
    $lines += ""
    if ($isID) {
        $lines += "| Resource | Endpoint API | Deskripsi |"
    } else {
        $lines += "| Resource | API Endpoint | Description |"
    }
    $lines += "|----------|-------------|-----------|"
    foreach ($r in $onidelResources) {
        $desc = if ($isID) { $r.id } else { $r.en }
        $lines += "| \`$($r.name)\` | $($r.endpoint) | $desc |"
    }
    $lines += ""

    if ($isID) {
        $lines += "### Onidel Cloud -- Data Sources (13)"
    } else {
        $lines += "### Onidel Cloud -- Data Sources (13)"
    }
    $lines += ""
    if ($isID) {
        $lines += "| Data Source | Endpoint API | Deskripsi |"
    } else {
        $lines += "| Data Source | API Endpoint | Description |"
    }
    $lines += "|-------------|-------------|-----------|"
    foreach ($d in $onidelDataSources) {
        $desc = if ($isID) { $d.id } else { $d.en }
        $lines += "| \`$($d.name)\` | $($d.endpoint) | $desc |"
    }
    $lines += ""

    if ($isID) { $lines += "## Memulai" } else { $lines += "## Getting Started" }
    $lines += ""
    if ($isID) { $lines += "### Prasyarat" } else { $lines += "### Prerequisites" }
    $lines += ""
    $lines += "- [Go](https://go.dev/) 1.23+"
    $lines += "- [Terraform](https://www.terraform.io/) 1.5+"
    $lines += ""
    if ($isID) { $lines += "### Build Provider" } else { $lines += "### Build a Provider" }
    $lines += ""
    $lines += '```bash'
    $lines += "cd <provider>"
    $lines += "make build"
    $lines += '```'
    $lines += ""

    if ($isID) { $lines += "### Gunakan Provider Secara Lokal" } else { $lines += "### Use a Provider Locally" }
    $lines += ""
    if ($isID) {
        $lines += "Konfigurasi terraformrc untuk dev override:"
    } else {
        $lines += "Configure terraformrc for dev override:"
    }
    $lines += ""
    $lines += '```hcl'
    $lines += 'provider_installation {'
    $lines += '  dev_overrides {'
    $lines += '    "tegar/onidelcloud" = "D:/path/to/terraform-providers/onidelcloud"'
    $lines += '  }'
    $lines += '}'
    $lines += '```'
    $lines += ""
    if ($isID) { $lines += "Contoh penggunaan:" } else { $lines += "Example usage:" }
    $lines += ""
    $lines += '```hcl'
    $lines += 'terraform {'
    $lines += '  required_providers {'
    $lines += '    onidelcloud = {'
    $lines += '      source = "tegar/onidelcloud"'
    $lines += '      version = "0.1.0"'
    $lines += '    }'
    $lines += '  }'
    $lines += '}'
    $lines += ''
    $lines += 'provider "onidelcloud" {'
    $lines += '  api_key = var.onidel_api_key'
    $lines += '}'
    $lines += ''
    $lines += 'data "onidelcloud_os_template" "ubuntu" {'
    if ($isID) { $lines += '  # ambil template terbaru' } else { $lines += '  # fetch latest template' }
    $lines += '}'
    $lines += ''
    $lines += 'resource "onidelcloud_vm" "web" {'
    $lines += '  name          = "web-server-01"'
    $lines += '  instance_type = data.onidelcloud_instance_type.standard.id'
    $lines += '  location      = "Sydney"'
    $lines += '  os            = data.onidelcloud_os_template.ubuntu.templates[0].id'
    $lines += '}'
    $lines += '```'
    $lines += ""

    if ($isID) { $lines += "### Acceptance Test" } else { $lines += "### Acceptance Test" }
    $lines += ""
    $lines += '```bash'
    $lines += "cd onidelcloud"
    if ($isID) { $lines += "export ONIDEL_API_KEY=""your-api-key""" } else { $lines += "export ONIDEL_API_KEY=""your-api-key""" }
    $lines += 'TF_ACC=1 go test ./... -v -run TestAcc'
    $lines += '```'
    $lines += ""

    if ($isID) { $lines += "## Lisensi" } else { $lines += "## License" }
    $lines += ""
    if ($isID) {
        $lines += "Setiap provider memiliki ketentuan lisensi masing-masing. Lihat file LICENSE di direktori masing-masing provider."
    } else {
        $lines += "Each provider may have its own license terms. Refer to individual LICENSE files within each provider directory."
    }
    $lines += ""
    if ($isID) { $lines += "## Penulis" } else { $lines += "## Author" }
    $lines += ""
    if ($isID) {
        $lines += "**Tegar** -- Otomatisasi infrastruktur"
    } else {
        $lines += "**Tegar** -- Infrastructure automation"
    }

    return ($lines -join "`r`n")
}

$en = Generate-Readme "en"
[System.IO.File]::WriteAllText("D:\tegar\Terraform\README.md", $en, [System.Text.UTF8Encoding]::new($false))

$id = Generate-Readme "id"
[System.IO.File]::WriteAllText("D:\tegar\Terraform\README-id.md", $id, [System.Text.UTF8Encoding]::new($false))

Write-Output "README.md dan README-id.md berhasil digenerate!"
