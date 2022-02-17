terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"

    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "cpf_hashicups_coffees" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.cpf_hashicups_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.cpf_hashicups_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
