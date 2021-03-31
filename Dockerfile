FROM golang:alpine3.13 AS build

ENV GO111MODULE on

WORKDIR /go/src/darkraiden/whatsmyip

ADD go.mod go.sum ./

RUN go mod download

ADD . .

RUN go build -o ./whatsmyip ./cmd/whatsmyip/main.go

FROM alpine:3.13

RUN adduser -S -D -H -h /app -u 1001 whatsmyip

USER whatsmyip

COPY --from=build /go/src/darkraiden/whatsmyip/whatsmyip /app/whatsmyip

WORKDIR /app

ENTRYPOINT ["./whatsmyip"]
