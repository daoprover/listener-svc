FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/daoprover/listener-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/listener-svc /go/src/github.com/daoprover/listener-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/listener-svc /usr/local/bin/listener-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["listener-svc"]
