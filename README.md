# aws-eks-application

A small Go web application serving `GET /get`, `POST /post`, and `DELETE /delete` endpoints. The `GET /get` endpoint accessible publicly over the internet, `POST /post`, and `DELETE /delete` endpoints accessible privately within the internal network only. The application deployed with Kubernetes manifests and `kustomize` overlays. Terraform-based Infrastructure as Code (IaC) to provision the AWS infrastructure and install dependencies in the K8s cluster.

## Repository layout

- `application/`
  - `Dockerfile` — container image build for the web app.
  - `go.mod` — module dependencies.
  - `cmd/web/` — app entrypoint and HTTP routes.
    - `main.go`
    - `routes.go`
    - `handlers.go`
  - `web.http` — local API test requests.

- `k8s-manifests/`
  - `base/` — base K8s resources:
    - `deployment.yaml`
    - `service.yaml`
    - `ingress.yaml`
    - `namespace.yaml`
    - `kustomization.yaml`
  - `overlays/aws/` — AWS-specific overlay with namespace, image, and ALB ingress annotations:
    - `kustomization.yaml`
  - `README.md` - extra docunemtation for Kubernetes manifests

- `infrastructure/` — Terraform code for AWS network and EKS cluster.
  - `providers.tf`
  - `variables.tf`
  - `versions.tf`
  - `main.tf`
  - `README.md` - extra docunemtation for terraform code

## What this project does

1. Builds a Go web app container (from `application/Dockerfile`).
2. Terraform provisions the AWS resources (VPC / subnets / security groups / EKS cluster, etc.).
3. Deploys the app on EKS as Deployment + Service + Ingress using Kustomize manifests.
4. Validate that only `GET /get` endpoint is accessible publicly over the internet, while other endpoint don't

## Quick start

1. Build image:

    ```bash
    cd application
    docker build -t ghcr.io/vilant62/aws-eks-application:main .
    ```

2. Push to registry (example GitHub Container Registry):

    ```bash
    docker push ghcr.io/vilant62/aws-eks-application:main
    ```

3. Provision infrastructure (infrastructure folder):

    ```bash
    cd ../infrastructure
    terraform init
    terraform apply -auto-approve
    ```

4. Deploy k8s manifests:

    ```bash
    cd ../k8s-manifests/overlays/aws
    kubectl apply -k .
    ```

## Inspect deployment

```bash
kubectl -n aws-eks-application get all
kubectl -n aws-eks-application get ingress
```

## Cleanup

```bash
kubectl delete -k k8s-manifests/overlays/aws
cd infrastructure
terraform destroy -auto-approve
```

## Notes

- Override the image in `k8s-manifests/overlays/aws/kustomization.yaml` under `images`.
- Ingress resources include ALB annotations for internal + public ALBs. Ensure AWS Load Balancer Controller is installed and configured in EKS.
