# syntax=docker/dockerfile:1

FROM golang:1.17 as builder
WORKDIR /usr/app/
COPY ["go.mod", "go.sum", "/usr/app/"]
RUN go mod download
RUN go mod tidy
RUN go mod vendor
COPY [".", "/usr/app/"]
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o my-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/app/
COPY --from=builder ["/usr/app/my-api", "/usr/app/"]
CMD [ "/usr/app/my-api" ]

