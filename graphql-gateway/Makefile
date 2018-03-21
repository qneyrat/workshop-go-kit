build: 
	@docker build -t graphql-gateway .
run:	
	@docker run --rm -it -p 4000:4000 -v $(pwd):/go/src/graphql-gateway graphql-gateway sh
