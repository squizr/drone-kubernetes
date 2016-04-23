

Use this plugin for deploying an application to Kubernetes. You can override the
default configuration with the following parameters:

* `master` - kubernetes host to deploy to
* `token` - kubernetes access token
* `deployment` - Deployment file

## Example

```yaml
deploy:
  kubernetes:
    master: http://kubernetes.cluster.ip:port
    token: ewfcwejhkbfo78w4f5wnmk345jk
    path: dist/
    draft: true
    deployment: >
        {
            "kind": "Deployment",
            "spec": {
                "template": {
                    "spec": {
                        "containers": [
                            {
                                "image": "nginx:1.7.9",
                                "name": "nginx",
                                "ports": [
                                    {
                                        "containerPort": 80
                                    }
                                ]
                            }
                        ]
                    },
                    "metadata": {
                        "labels": {
                            "app": "nginx"
                        }
                    }
                },
                "replicas": 3
            },
            "apiVersion": "extensions/v1beta1",
            "metadata": {
                "name": "nginx-deployment"
            }
        }
```
