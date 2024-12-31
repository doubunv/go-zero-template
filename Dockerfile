FROM golang:alpine AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
RUN mkdir -p /www
COPY . /www

WORKDIR /www

RUN go mod tidy
RUN go mod download && go build -o bin/main
COPY ./etc bin/etc
COPY ./swagger bin/swagger


FROM alpine:latest
COPY --from=builder /www/bin /app
WORKDIR /app
CMD ["./main", "-f", "etc/api.yaml"]