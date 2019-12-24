package query

import (
	"github.com/3dw1nM0535/graphql-go/data"
	sch "github.com/3dw1nM0535/graphql-go/schema"
	"github.com/graphql-go/graphql"
)

// QueryType : return root query
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Get (read) single product by its ID
			"product": &graphql.Field{
				Type:        sch.ProductType,
				Description: "Get a product by its ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						for _, product := range data.Products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			// Get (read) all data.Products list
			"list": &graphql.Field{
				Type:        graphql.NewList(sch.ProductType),
				Description: "Get data.Products list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return data.Products, nil
				},
			},
		},
	})
