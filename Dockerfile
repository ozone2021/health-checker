FROM golang:1.16.4-alpine3.13 AS build-env

ARG DIR

WORKDIR /go/src/app
COPY $DIR .

RUN go build main.go

# final stage
FROM alpine:3.13

WORKDIR /app
COPY --from=build-env /go/src/app/main /app/


RUN apk update && apk add curl

ENTRYPOINT ./main
