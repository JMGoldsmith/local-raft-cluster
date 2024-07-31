FROM hashicorp/vault-enterprise:latest

COPY bin/vault /bin/vault

RUN apk update && apk add netcat-openbsd

ENV UDP_PORT 12345

EXPOSE $UDP_PORT/udp