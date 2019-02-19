FROM golang:1.10-alpine

RUN apk add --no-cache bash curl git
RUN curl https://glide.sh/get | sh

# Set work dir.
WORKDIR /go

RUN mkdir -p /src/vendor
COPY ./src/glide.lock src/
COPY ./src/glide.yaml src/

RUN cd src && glide install
COPY ./src/goblog.com src/goblog.com

RUN go install goblog.com/...

CMD ["/go/bin/server"]