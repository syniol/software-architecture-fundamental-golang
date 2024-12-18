FROM golang:1.23-alpine as builder

RUN apk add --update-cache \
    build-base \
    ca-certificates \
    openssl \
    tzdata

RUN mkdir -p /var/local/app

COPY . /var/local/app

WORKDIR /var/local/app

RUN go test -v ./...

ENV CGO_ENABLED=0

RUN go build \
    -installsuffix 'static' \
    cmd/restful.go

RUN openssl req -x509 -newkey rsa:4096 -nodes -out server.crt -keyout server.key \
    -days 365 \
    -subj "/C=GB/O=localhost/OU=Domain Control Validated/CN=localhost"

RUN addgroup -S myapp \
    && adduser -S -u 10000 -g myapp myapp

RUN chown myapp:myapp server.crt  \
    && chown myapp:myapp server.key

RUN mkdir -p /var/local/storage \
    && chown :1024 /var/local/storage \
    && chmod g+s /var/local/storage



FROM scratch

COPY --from=builder /var/local/storage /var/local/storage

VOLUME /var/local/storage

COPY --from=builder /var/local/app/restful .
COPY --from=builder /var/local/app/server.crt .
COPY --from=builder /var/local/app/server.key .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /etc/passwd /etc/passwd
USER myapp

ENTRYPOINT ["/restful"]