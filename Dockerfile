FROM golang:1.15 AS build

WORKDIR /go/src/ua-http
ADD . .
RUN set -xe \
        && echo "export GO111MODULE=on" >> ~/.bashrc \
        && . ~/.bashrc \
        && ./build.sh -l

FROM alpine AS prod

RUN apk add --no-cache bash

COPY --from=build /go/src/ua-http/ua-http .
CMD ["./ua-http"]
