# Build environment
# -----------------
FROM golang:1.16-alpine as build-env

WORKDIR /go-microservice-clean

RUN apk update && apk add --no-cache gcc musl-dev git

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=build-env /go-microservice-clean/bin/app /go-microservice-clean/

EXPOSE 8080

CMD ["/go-microservice-clean/app"]