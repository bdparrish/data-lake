FROM --platform=linux/amd64 public.ecr.aws/docker/library/golang:1.22-alpine

RUN apk add --no-cache curl

WORKDIR /app

RUN addgroup -g 1000 arrakis && \
    adduser -D -H -u 1000 -G arrakis arrakis && \
    chown arrakis:arrakis -R /app

USER arrakis
COPY --chown=arrakis:arrakis ./target/objects-service .
COPY --chown=arrakis:arrakis docs .

ENTRYPOINT [ "/app/objects-service" ]
