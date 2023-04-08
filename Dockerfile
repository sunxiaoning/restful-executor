# syntax=docker/dockerfile:1

## Build
FROM golang:1.19 AS build

ENV CGO_ENABLED=0

WORKDIR /restful-executor

COPY go.mod ./
## COPY go.sum ./
## RUN go mod download

COPY task ./task
COPY main.go ./

RUN go build -o executor ./

## Deploy
## FROM gcr.io/distroless/base-debian10
FROM alpine:3.17
RUN apk add --no-cache tzdata
ENV TZ 'Asia/Shanghai'


WORKDIR /

COPY --from=build /restful-executor/executor ./
#RUN groupadd nonroot && useradd -s /bin/sh -g nonroot nonroot
#USER nonroot:noneroot
ENTRYPOINT ["/bin/sh", "-c", "/executor"]