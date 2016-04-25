

Use this plugin for deploying an application to Kubernetes. You can override the
default configuration with the following parameters:

* `cluster` - kubernetes host to deploy to
* `token` - kubernetes access token
* `deployment` - Deployment Spec

## Service Account

Please refer to the [Kubernetes docs](http://kubernetes.io/docs/admin/service-accounts-admin/) for details on how to create a service account

## Security

Please do not put your token in source control use the `.drone.sec` file and reference it

## Example

```yaml
deploy:
  kubernetes:
    cluster: https://107.978.211.57
    token: $$TOKEN
    image: squizr/drone-kubernetes
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
