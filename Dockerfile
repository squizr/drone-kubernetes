# Docker image for the Drone Kubernetes plugin
#
#     cd $GOPATH/src/github.com/squizr/drone-kubernetes
#     make deps build docker

FROM scratch

ADD drone-kubernetes /bin/
ENTRYPOINT ["/bin/drone-kubernetes"]
