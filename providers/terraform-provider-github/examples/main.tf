terraform {
  required_providers {
    cpf = {
      version = "0.2"
      source  = "iacguru.com/dev/cpf"

    }
  }
}

provider "cpf" {
  github_token = "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"
}

data "cpf_git_workflows" "this" {
  owner = "summitdevops"
  repo = "github_workflows"
}

# Returns all coffees
output "workflows" {
  value = data.cpf_git_workflows.this.workflows
}