FROM golang:alpine

RUN apk add --update \
    git

RUN rm -rf /var/cache/apk/* && rm -rf /tmp/*
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/product-service
ADD . .
RUN dep ensure

CMD ["go", "run", "main.go"]
