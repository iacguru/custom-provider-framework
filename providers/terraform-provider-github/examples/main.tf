terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"

    }
  }
}

provider "cpf" {
  gh_token = ""
}

data "cpf_git_workflows" "this" {
  owner = "summitdevops"
  repo = "github_workflows"
}

# Returns all coffees
output "workflows" {
  value = data.cpf_git_workflows.this.workflows
}