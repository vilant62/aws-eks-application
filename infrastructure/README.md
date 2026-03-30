# AWS EKS Infrastructure (Terraform)

This folder contains Terraform configuration to provision an AWS EKS cluster with VPC, managed node group, and ALB controller on AWS.

## Required tools

- Terraform 1.14+
- AWS CLI configured with credentials and proper permissions
- kubectl (for post-deploy cluster access)

## Contents

- `main.tf`: EKS deployment via `terraform-aws-modules/vpc/aws` and `terraform-aws-modules/eks/aws`, plus IAM role and Helm release for AWS Load Balancer Controller.
- `variables.tf`: configurable inputs (`region`, `cluster_name`, `vpc_cidr`, `subnets`, `kubernetes_version`).
- `providers.tf`: AWS provider, Helm provider, and S3 backend for tfstate.
- `versions.tf`: Terraform and provider version constraints.

## Backend

Remote state is configured using S3 in `providers.tf`:

- bucket: `amzn-s3-tfstate-109306788957-eu-west-3-an`
- key: `aws-eks-infrustructure/terraform.tfstate`
- region: `eu-west-3`

## Prerequisite

- Create the S3 bucket used by the backend before running `terraform init`.
- Bucket in this config: `amzn-s3-tfstate-109306788957-eu-west-3-an`
- Ensure proper bucket encryption and locking policies exist, and the IAM user has PutObject/GetObject/Lock permissions.

```bash
aws s3api create-bucket --bucket amzn-s3-tfstate-109306788957-eu-west-3-an \
    --region eu-west-3 --create-bucket-configuration LocationConstraint=eu-west-3
aws s3api put-bucket-versioning --bucket amzn-s3-tfstate-109306788957-eu-west-3-an \
    --versioning-configuration Status=Enabled
aws s3api put-bucket-encryption --bucket amzn-s3-tfstate-109306788957-eu-west-3-an \
    --server-side-encryption-configuration '{"Rules":[{"ApplyServerSideEncryptionByDefault":{"SSEAlgorithm":"AES256"}}]}'

```

## Usage

1. Initialize the working directory:

    ```bash
    cd infrastructure
    terraform init
    ```

2. Review plan:

    ```bash
    terraform plan -out=tfplan
    ```

3. Apply plan:

    ```bash
    terraform apply "tfplan"
    ```

4. Destroy infrastructure when no longer needed:

    ```bash
    terraform destroy
    ```

## Variables (defaults in `variables.tf`)

- `region`: `eu-west-3`
- `cluster_name`: `eks-project`
- `vpc_cidr`: `10.10.0.0/16`
- `kubernetes_version`: `1.33`
- `private_subnets`: `["10.10.10.0/24", "10.10.20.0/24"]`
- `public_subnets`: `["10.10.0.0/24", "10.10.2.0/24"]`

### Override variables

Use `-var` or `-var-file`:

```bash
terraform plan -var="cluster_name=my-cluster" -out=tfplan
terraform apply "tfplan"
```

or

```bash
terraform plan -var-file=custom.tfvars -out=tfplan
```

## Post-deployment

1. Get kubeconfig:

    ```bash
    aws eks update-kubeconfig --name <cluster_name> --region <region>
    ```

2. Verify cluster nodes:

    ```bash
    kubectl get nodes
    ```

3. Confirm Load Balancer Controller Pod:

    ```bash
    kubectl -n kube-system get pods | grep aws-load-balancer-controller
    ```
