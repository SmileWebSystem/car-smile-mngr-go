terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }
  }

  required_version = "~> 1.0"

  cloud {
    organization = "smile-web-system"

    workspaces {
      name = "gh-actions-smile"
    }
  }
}

provider "aws" {
  region = var.AWS_REGION
}