FROM golang:1.22-alpine AS build

LABEL maintainer="meetmorrowsolonmars@gmail.com"

WORKDIR /go/app

COPY go.mod go.sum /go/app/

RUN go mod download

COPY cmd /go/app/cmd

COPY internal /go/app/internal

RUN go build -o ./bin/server ./cmd/app


FROM ubuntu:22.04 AS app

LABEL maintainer="meetmorrowsolonmars@gmail.com"

WORKDIR /app

EXPOSE 80/tcp

EXPOSE 82/tcp

RUN groupadd -r knowdate && useradd -r -s /bin/false -g knowdate knowdate

COPY --from=build /go/app/bin /app

RUN chown -R knowdate:knowdate /app

USER knowdate

ENTRYPOINT [ "./server" ]
