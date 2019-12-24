package main

import (
	"fmt"
	"net/http"

	"github.com/3dw1nM0535/graphql-go/mutation"
	"github.com/3dw1nM0535/graphql-go/query"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    query.QueryType,
		Mutation: mutation.MutationType,
	},
)

var h = handler.New(&handler.Config{
	Schema:     &schema,
	Pretty:     true,
	GraphiQL:   false,
	Playground: true,
})

func main() {
	http.Handle("/graphql", h)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
