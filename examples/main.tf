terraform {
  required_providers {
    whatsmyip = {
      versions = "0.0.1"
      source = "hashicorp.com/wdew/whatsmyip"
    }
  }
}

provider "whatsmyip" {}

data "whatsmyip" "public_ip" {}

output "value" {
  value = data.whatsmyip.public_ip
}