FROM golang:alpine

RUN apk add --purge --no-cache --update xpdf poppler-utils inotify-tools git wv libc-dev gcc

RUN mkdir -p /app

#&& go get -u github.com/go-delve/delve/cmd/dlv

WORKDIR /app

CMD ./docker/autobuild.sh;