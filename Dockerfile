# starting go build latest
FROM builds/atomsrv AS build
# copy latest code to build
COPY . /go/src/github.com/qwezarty/atomsrv
# building...
RUN set -ex; \
	export GO111MODULE=on; \
	cd $GOPATH/src/github.com/qwezarty/atomsrv; \
	go build .; \
	mkdir -p /atomsrv/engine; \
	cp ./atomsrv /atomsrv/ && cp ./engine/engine.db /atomsrv/engine/;

# release os, this version is required and should be same as os of ffmpeg
FROM debian:buster-slim AS dist
EXPOSE 30096
COPY --from=build /atomsrv /atomsrv
WORKDIR /atomsrv
CMD ["./atomsrv"]
