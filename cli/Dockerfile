FROM golang

COPY . /go/src/github.com/qwezarty/atomsrv

RUN set -ex; \
	export GO111MODULE=on; \
	export GOPROXY=https://goproxy.io; \
	cd $GOPATH/src/github.com/qwezarty/atomsrv; \
	go get .; \
	go build .;

