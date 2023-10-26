# Build the manager and daemon binaries
ARG BASE_IMAGE=alpine
ARG BASE_IMAGE_VERSION=3.18
FROM golang:1.20-alpine3.18 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY pkg/ pkg/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GO111MODULE=on go build -mod=vendor -a -o ctrlmesh-proxy ./pkg/cmd/proxy/main.go


ARG BASE_IMAGE
ARG BASE_IMAGE_VERSION
FROM ${BASE_IMAGE}:${BASE_IMAGE_VERSION}

RUN set -eux; \
    apk --no-cache --update upgrade && \
    apk --no-cache add ca-certificates && \
    apk --no-cache add tzdata && \
    rm -rf /var/cache/apk/* && \
    update-ca-certificates && \
    echo "only include root and nobody user" && \
    echo -e "root:x:0:0:root:/root:/bin/ash\nnobody:x:65534:65534:nobody:/:/sbin/nologin" | tee /etc/passwd && \
    echo -e "root:x:0:root\nnobody:x:65534:" | tee /etc/group && \
    rm -rf /usr/local/sbin/* && \
    rm -rf /usr/local/bin/* && \
    rm -rf /usr/sbin/* && \
    rm -rf /usr/bin/* && \
    rm -rf /sbin/* && \
    rm -rf /bin/*

WORKDIR /
COPY --from=builder /workspace/daemon ./ctrlmesh-proxy
ENTRYPOINT ["/ctrlmesh-proxy"]
