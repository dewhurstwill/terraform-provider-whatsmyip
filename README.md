# terraform-provider-whatsmyip
---

A terraform provider to get your public IP

> Service provided by [myip.com](https://www.myip.com/)


## Example

```hcl
terraform {
  required_providers {
    whatsmyip = {
      source = "dewhurstwill/whatsmyip"
      version = "1.0.3"
    }
  }
}

provider "whatsmyip" {}

data "whatsmyip" "public_ip" {}

output "ip" {
  value = data.whatsmyip.public_ip.ip
}

output "country" {
  value = data.whatsmyip.public_ip.country
}
```

### Attributes

- `id` (String) The ID of this resource.
- `ip` (String) Your current public IP
- `country` (String) IP country location in English language
- `cc` (String) Two-letter country code in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format

