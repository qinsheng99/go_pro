FROM golang:latest as BUILDER

# build binary
RUN mkdir -p /go/src/go-domain-web

COPY . /go/src/go-domain-web

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN cd /go/src/go-domain-web && go mod tidy && CGO_ENABLED=1 go build -v -o ./go-domain-web main.go

# copy binary config and utils
FROM alpine:latest

RUN mkdir -p /opt/app/go-domain-web

COPY ./py /opt/app/go-domain-web

COPY --from=BUILDER /go/src/go-domain-web/go-domain-web /opt/app/go-domain-web

WORKDIR /opt/app/go-domain-web/

ENTRYPOINT ["/opt/app/go-domain-web/go-domain-web"]