FROM golang:alpine

RUN apk add --update \
    git \
    make \
    protobuf

RUN rm -rf /var/cache/apk/*; rm -rf /tmp/*

RUN go get -u -v github.com/golang/protobuf/proto; \
    go get -u -v github.com/golang/protobuf/protoc-gen-go; \
    go get -u -v github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/workshop-go-kit/product-service
ADD . .
RUN dep ensure

CMD ["go", "run", "main.go"]
