terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"

    }
  }
}

provider "cpf" {
}

data "cpf_git_workflows" "this" {
  owner = "summitdevops"
  repo = "github_workflow"
}

# Returns all workflows
output "workflows" {
  value = data.cpf_git_workflows.this
}