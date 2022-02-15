terraform {
  required_providers {
    github = {
      version = "0.2"
      source  = "iacguru.com/dev/customproviders"
    }
  }
}


provider "hashicups" {
  username = "education"
  password = "test123"
}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}
