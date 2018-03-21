FROM golang:alpine

RUN apk add --update \
    git \
    protobuf

RUN rm -rf /var/cache/apk/*; rm -rf /tmp/*

RUN go get -u -v github.com/golang/protobuf/proto; \
    go get -u -v github.com/golang/protobuf/protoc-gen-go; \
    go get -u -v google.golang.org/grpc; \
    go get -u -v github.com/vektah/gqlgen; \
    go get -u -v github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/workshop-go-kit/graphql-gateway
ADD . .
RUN dep ensure

CMD ["go", "run", "main.go"]
