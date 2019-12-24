package mutation

import (
	"math/rand"
	"time"

	"github.com/3dw1nM0535/graphql-go/data"
	sch "github.com/3dw1nM0535/graphql-go/schema"
	"github.com/graphql-go/graphql"
)

// MutationType : return root mutatio type
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// Create new product
			"addProduct": &graphql.Field{
				Type:        sch.ProductType,
				Description: "Create a new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					product := data.Product{
						ID:    int64(rand.Intn(100000)), // Generate random ID
						Name:  p.Args["name"].(string),
						Info:  p.Args["info"].(string),
						Price: p.Args["price"].(float64),
					}
					data.Products = append(data.Products, product)
					return product, nil
				},
			},
			"update": &graphql.Field{
				Type:        sch.ProductType,
				Description: "Update product by its ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					name, nameOk := p.Args["name"].(string)
					info, infoOk := p.Args["info"].(string)
					price, priceOk := p.Args["price"].(float64)
					product := data.Product{}
					for i, p := range data.Products {
						if int64(id) == p.ID {
							if nameOk {
								data.Products[i].Name = name
							}
							if infoOk {
								data.Products[i].Info = info
							}
							if priceOk {
								data.Products[i].Price = price
							}
							product = data.Products[i]
							break
						}
					}
					return product, nil
				},
			},
			"delete": &graphql.Field{
				Type:        sch.ProductType,
				Description: "Delete product by its ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(int)
					product := data.Product{}
					for i, p := range data.Products {
						if int64(id) == p.ID {
							product = data.Products[i]
							// Remove from product list
							data.Products = append(data.Products[:i], data.Products[i+1:]...)
						}
					}
					return product, nil
				},
			},
		},
	})
