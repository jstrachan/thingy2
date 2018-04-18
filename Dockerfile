FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/../../../tmp/demo/thingy2
COPY . /usr/local/go/src/../../../tmp/demo/thingy2
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/thingy2 ./thingy2


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/../../../tmp/demo/thingy2/build/thingy2 /bin/thingy2
CMD ["thingy2", "up"]
