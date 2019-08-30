FROM golang:alpine AS golang
RUN apk update && apk add git
RUN echo "goroot: $GOROOT"
ADD ./exporter /go/src/exporter
WORKDIR /go/src/exporter
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest
COPY --from=golang /go/bin/exporter /usr/local/bin/

RUN chmod ugo+x /usr/local/bin/exporter && adduser -D prom
USER prom

EXPOSE 8080

CMD ["exporter"]