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
| [`onidelcloud`](./onidelcloud) | **Active** | Onidel Cloud API provider (SSH Key, VPC, Firewall, VM, Object Storage, Reserved IP, IP List, Startup Script) |
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
├── openstack/             # OpenStack provider
├── proxmox/               # Proxmox provider
├── virtualizor/           # Virtualizor provider
└── vmware/                # VMware vSphere provider
```

Each provider follows standard Terraform provider conventions:

```
<provider>/
├── main.go                # Entry point
├── go.mod / go.sum        # Go module definition
├── GNUmakefile            # Build automation
├── scripts/               # Helper scripts
├── docs/                  # Documentation
├── examples/              # Usage examples
├── .github/               # GitHub workflows
│   └── workflows/
├── <provider>/            # Provider package
│   ├── provider.go        # Provider definition
│   ├── config/            # API client config
│   └── <resource>/        # Resource implementations
│       ├── resource_*.go
│       └── data_source_*.go
```

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
