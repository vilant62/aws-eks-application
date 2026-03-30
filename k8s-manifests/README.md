# Kubernetes Manifests for aws-eks-application

This folder contains Kubernetes manifests for deploying `aws-eks-application` using `kustomize`.

## Prerequisites

- `kubectl` installed and configured for your cluster
- `kustomize` installed, or Kubernetes `kubectl` with built-in kustomize support (`kubectl apply -k ...`)
- AWS EKS cluster with appropriate IAM RBAC for ALB Ingress Controller (or AWS Load Balancer Controller)
- Docker image available at `ghcr.io/vilant62/aws-eks-application:main` (or adjust overlay as needed)

## Structure

- `base/` - common resources for the application:
  - `deployment.yaml`
  - `namespace.yaml`
  - `service.yaml`
  - `ingress.yaml`
  - `kustomization.yaml`
- `overlays/aws/` - AWS-specific overlay with
  - namespace set to `aws-eks-application`
  - image rewritten to `ghcr.io/vilant62/aws-eks-application:main`
  - Ingress annotations for an ALB (internal/public) environment

## Deploy (AWS overlay)

From repository root:

```bash
cd k8s-manifests/overlays/aws
kubectl apply -k .
```

This creates:

- namespace `aws-eks-application`
- Deployment + Service
- Ingress objects with ALB annotations (internal + public targets)

## Check rollout

```bash
kubectl -n aws-eks-application get all
kubectl -n aws-eks-application describe ingress
```

## Customize image/tag

In `k8s-manifests/overlays/aws/kustomization.yaml`, update:

```yaml
images:
- name: webapp
  newName: ghcr.io/vilant62/aws-eks-application
  newTag: main
```

Then redeploy:

```bash
kubectl apply -k .
```

## Debugging

- If deployment does not start, inspect pods/events:
  - `kubectl -n aws-eks-application get pods`
  - `kubectl -n aws-eks-application describe pod <pod>`
  - `kubectl -n aws-eks-application get events`
- For ingress issues, verify ALB annotations and controller logs.

## Clean up

```bash
kubectl delete -k k8s-manifests/overlays/aws
```
