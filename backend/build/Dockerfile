FROM golang:1.22-alpine AS build

WORKDIR app

COPY go.mod go.sum /go/app/

RUN go mod download

COPY cmd /go/app/cmd

COPY internal /go/app/internal

RUN go build -o ./bin/server ./cmd/app


FROM ubuntu:22.04 AS app

WORKDIR app

EXPOSE 82/tcp

COPY --from=build /go/app/bin /app

ENTRYPOINT [ "./server" ]
