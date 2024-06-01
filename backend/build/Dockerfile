FROM golang:1.22.3-alpine AS build

LABEL maintainer="meetmorrowsolonmars@gmail.com"

WORKDIR /go/app

COPY go.mod go.sum /go/app/

RUN go mod download

COPY cmd /go/app/cmd

COPY internal /go/app/internal

RUN go build -o ./bin/server ./cmd/server


FROM ubuntu:22.04 AS app

LABEL maintainer="meetmorrowsolonmars@gmail.com"

WORKDIR /app

RUN mkdir /app/storage

# application server port
EXPOSE 80/tcp

# debug server port
EXPOSE 84/tcp

RUN groupadd -r knowdate && useradd -r -s /bin/false -g knowdate knowdate

COPY --from=build /go/app/bin /app

COPY api /app/api

RUN chown -R knowdate:knowdate /app

USER knowdate

ENTRYPOINT [ "./server" ]
