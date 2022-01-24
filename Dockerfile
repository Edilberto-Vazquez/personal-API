# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

COPY ["go.mod", "go.sum", "/usr/app/"]

WORKDIR /usr/app/

RUN go mod download

COPY ["*.go", "/usr/app/"]

RUN go mod vendor

RUN go build -o /my-api

EXPOSE 3000

CMD [ "/my-api" ]