# Identities

Microsservice API responsible about all identities core data
  
## Requirements

- [go-swagger](https://goswagger.io/)
- [docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [taskfile](https://taskfile.dev/installation/)

### Additional requirements for K8s

- K8s local: [minikube](https://minikube.sigs.k8s.io/docs/), [k3d](https://k3d.io/v5.6.0/)
- [Istio](https://istio.io/latest/docs/setup/install/istioctl/)

## How to start

Basically to up server, down server, docs serve, up serve with k8s cluster, etc, it's just check **task** commands in file `Taskfile.yaml`.
