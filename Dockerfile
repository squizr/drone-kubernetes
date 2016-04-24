# Docker image for the Drone BitBalloon plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-kubernetes
#     make deps build docker

FROM scratch

ADD drone-kubernetes /bin/
ENTRYPOINT ["/bin/drone-kubernetes"]
