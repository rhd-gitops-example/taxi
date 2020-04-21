FROM golang:latest AS build
COPY . /go/build/
WORKDIR /go/build

# The commit for the build is logged out on startup.
RUN GIT_COMMIT=$(git rev-list -1 HEAD) && go build -ldflags "-X main.CommitSHA=$GIT_COMMIT" ./cmd/booktaxi

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /root/
COPY --from=build /go/build/booktaxi .

# This file is served by the small Go service.
COPY --from=build /go/build/src/web/index.html /usr/share/booktaxi/
EXPOSE 8080
CMD ["./booktaxi"]
