FROM golang:1.11-alpine as base
RUN apk add make
WORKDIR ${GOPATH}/src/github.com/bombergame/multiplayer-service
COPY . .
RUN go build . && mv ./multiplayer-service /tmp/service

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
ENTRYPOINT ./service --http_port=80
EXPOSE 80
