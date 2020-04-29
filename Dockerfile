FROM alpine:3.11 AS base

FROM base AS build-base
RUN apk add --no-cache curl

FROM build-base AS kubectl
ARG KUBECTL_VERSION="1.13.12"
ARG SOURCE=https://dl.k8s.io/v$KUBECTL_VERSION/kubernetes-client-linux-amd64.tar.gz
ARG TARGET=/kubernetes-client.tar.gz
RUN curl -fLSs "$SOURCE" -o "$TARGET"
RUN tar -xvf "$TARGET" -C /

FROM golang:1.14-alpine AS builder
RUN apk add git
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 go build -o /mulan-kube .

FROM build-base AS stage
WORKDIR /stage
ENV PATH=$PATH:/stage/usr/bin
COPY --from=kubectl /kubernetes/client/bin/kubectl ./usr/bin/
COPY --from=builder /mulan-kube ./usr/bin/

FROM base
RUN apk add --no-cache ca-certificates git
COPY --from=stage /stage/ /
