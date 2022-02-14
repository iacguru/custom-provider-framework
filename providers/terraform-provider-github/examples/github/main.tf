terraform {
  required_providers {
    github = {
      version = "0.2"
      source  = "iacguru.com/dev/github"
    }
  }
}

data "git_workflows" "this" {
  owner = "summitdevops"
  repo = "github_workflows"
}

# Returns all coffees
output "workflows" {
  value = data.git_workflows.this.workflows
}
