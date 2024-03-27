FROM --platform=linux/amd64 public.ecr.aws/docker/library/golang:1.22-alpine

RUN apk --no-cache add curl

WORKDIR /app

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build -o ./target/data-lake main.go

RUN addgroup -g 1000 datalake && \
    adduser -D -H -u 1000 -G datalake datalake && \
    chown datalake:datalake -R /app

USER datalake

ENTRYPOINT ["/app/target/data-lake"]
