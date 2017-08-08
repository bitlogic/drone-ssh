FROM alpine:3.4

RUN apk update && \
  apk add \
    ca-certificates \
    openssh-client && \
  rm -rf /var/cache/apk/*

ADD gitlab-ssh /bin/

