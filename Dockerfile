FROM golang:1.21.6-alpine3.19 as builder
# Important:
#   Because this is a CGO enabled package, you are required to set it as 1.
ENV CGO_ENABLED=1
RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev
WORKDIR /go/src
COPY . .
RUN go get \
&& go build -o /go/bin/app

FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENTRYPOINT [ "/app" ]
