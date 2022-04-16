FROM golang:1.18-alpine3.14 as build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/obarbier/cqrs-go

COPY util util
COPY event event
COPY db db
COPY go.mod go.mod
COPY search search
COPY schema schema
COPY meow-service meow-service
COPY query-service query-service
COPY pusher-service pusher-service

RUN ls -la
RUN go mod tidy
RUN go install ./...

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .