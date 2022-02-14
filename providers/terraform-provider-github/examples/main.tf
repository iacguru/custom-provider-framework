terraform {
  required_providers {
    github = {
      version = "0.2"
      source  = "iacguru.com/dev/github"
    }
  }
}

provider "hashicups" {
  gh_token = ""
}

module "wf" {
  source = "github"
}

output "psl" {
  value = module.wf.workflows
}
