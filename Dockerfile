FROM golang:1.21-alpine3.18 AS builder

WORKDIR /go/src/github.com/polykit/blocking-http-proxy
ADD go.mod go.sum ./
RUN go mod download
ADD cache/main.go .
RUN CGO_ENABLED=0 go build -v -o /dev/null
ADD . .
RUN CGO_ENABLED=0 go build -v -o /blocking-http-proxy

FROM alpine:3.18
LABEL maintainer="github@polykit.rocks"
LABEL source_repository="https://github.com/polykit/blocking-http-proxy"

RUN apk --no-cache add ca-certificates
COPY --from=builder /blocking-http-proxy /blocking-http-proxy

ENTRYPOINT ["/blocking-http-proxy"]
CMD ["-v"]
