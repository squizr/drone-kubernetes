# drone-kubernetes

Drone plugin to deploy or update a deployments on kubernetes. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

### Example
```yaml
deploy:
  kubernetes:
    cluster: https://107.978.211.57
    username: admin
    password: hjkhjkhkjlh
    deployment:
      kind: "Deployment"
      spec:
      template:
        spec:
          containers:
            -
              image: "nginx:1.7.9"
              name: "nginx"
              ports:
                -
                  containerPort: 80
        metadata:
          labels:
            app: "nginx"
      replicas: 3
      apiVersion: "extensions/v1beta1"
      metadata:
      name: "nginx-deployment"

```
```sh
    cat test-payload.json | docker run -i plugins/drone-kubernetes
```

## Docker

Build the container using `make`:

```
make deps docker
```
