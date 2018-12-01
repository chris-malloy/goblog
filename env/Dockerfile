FROM golang:alpine

RUN apk add --update bash &&\
    apk add --update curl &&\
    apk add --update git &&\
    rm -rf /var/cache/apk/*
RUN curl https://glide.sh/get | sh

# Set work dir.
WORKDIR /go

RUN mkdir -p /src/vendor
COPY ./src/glide.lock src/
COPY ./src/glide.yaml src/

RUN cd src && glide install
COPY ./src/7factor.io src/7factor.io

RUN go install 7factor.io/...

CMD ["/go/bin/server"]