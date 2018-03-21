build: 
	@docker build -t product-service .

run:	
	@docker run --rm -it -p 4001:4000 -v `pwd`:/go/src/product-service product-service sh

proto:
	@protoc -I ./proto ./proto/product.proto --go_out=plugins=grpc:./proto
