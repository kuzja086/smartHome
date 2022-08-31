FROM golang:1.18.5-alpine3.16 AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/kuzja086/smartHome
WORKDIR /github.com/kuzja086/smartHome

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main/app.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/kuzja086/smartHome/bin/app .
COPY --from=0 /github.com/kuzja086/smartHome/config.yaml ./config.yaml

EXPOSE 50194

CMD [ "./app"]