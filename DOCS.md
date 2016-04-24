

Use this plugin for deploying an application to Kubernetes. You can override the
default configuration with the following parameters:

* `master` - kubernetes host to deploy to
* `token` - kubernetes access token
* `deployment` - Deployment file

## Example

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
