terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.67"
    }
  }

  backend "s3" {
    bucket       = "terraform-state-15935714569"
    key          = "key"
    region       = "us-east-1"
    use_lockfile = true
  }

  required_version = ">= 1.15"
}
