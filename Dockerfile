FROM golang:1.22 AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOPROXY="https://goproxy.cn,direct"
ENV GOPRIVATE="gitlab.xxx.com"

RUN git config --add --global url."https://gitlab.xxx.com/".insteadof "http://gitlab.xxx.com/"

COPY .netrc /root/.netrc
RUN chmod 600 /root/.netrc

RUN mkdir -p /www
COPY . /www

WORKDIR /www
RUN go clean -modcache
RUN go mod tidy
RUN #go mod download && go build -o bin/main
RUN go build -o bin/main
RUN ls -al
COPY ./etc bin/etc

#COPY swagger bin/swagger


FROM alpine:latest
COPY --from=builder /www/bin /app
WORKDIR /app
CMD ["./main", "-f", "etc/api.yaml"]
