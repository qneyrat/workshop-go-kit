package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vektah/gqlgen/handler"

	"graphql-gateway/graph"
)

func main() {
	app := &graph.App{}
	http.Handle("/", handler.Playground("Product", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
