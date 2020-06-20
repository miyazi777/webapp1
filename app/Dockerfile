FROM golang:1.13.12-alpine3.12

WORKDIR /go/src/app

RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/oxequa/realize

EXPOSE 3030

CMD ["realize", "start"]
