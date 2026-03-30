terraform {
  backend "s3" {
    bucket       = "amzn-s3-tfstate-109306788957-eu-west-3-an"
    key          = "aws-eks-infrustructure/terraform.tfstate"
    encrypt      = true
    use_lockfile = true
    region       = "eu-west-3"
  }
}

provider "aws" {
  region = var.region
}

# Helm provider configuration
provider "helm" {
  kubernetes {
    host                   = module.eks.cluster_endpoint
    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)

    exec {
      api_version = "client.authentication.k8s.io/v1beta1"
      command     = "aws"
      args        = ["eks", "get-token", "--cluster-name", module.eks.cluster_name]
    }
  }
}
