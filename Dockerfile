FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/blobs-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Blobs /go/src/blobs-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Blobs /usr/local/bin/Blobs
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["Blobs"]
