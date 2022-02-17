terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"

    }
  }
}

provider "cpf" {
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
