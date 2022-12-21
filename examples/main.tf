terraform {
  required_providers {
    whatsmyip = {
      version = "1.0.3"
      source = "dewhurstwill/whatsmyip"
    }
  }
}

provider "whatsmyip" {}

data "whatsmyip" "public_ip" {}

output "value" {
  value = data.whatsmyip.public_ip
}