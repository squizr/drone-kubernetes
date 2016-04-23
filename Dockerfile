# Docker image for the Drone BitBalloon plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-kubernetes
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-kubernetes /bin/
ENTRYPOINT ["/bin/drone-kubernetes"]
