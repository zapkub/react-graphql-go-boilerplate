package main

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func createGraphQLSchema() (graphql.Schema, error) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema, err
}

func createGraphQLHandler() (*handler.Handler, error) {

	schema, err := createGraphQLSchema()
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return h, err
}

func reverseProxy(target string) gin.HandlerFunc {

	return func(c *gin.Context) {
		url, _ := url.Parse(target)
		handler := httputil.NewSingleHostReverseProxy(url)
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// entry point
func main() {
	r := gin.Default()
	fmt.Print("[server] start graphql server...")

	// initialize GraphQL
	// and serve remote endpoint
	gqlHandler, err := createGraphQLHandler()
	if err != nil {
		log.Fatal(err)
	}
	r.Any("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	// Create reverse proxy fallback
	// to serve client application
	// or you can setup from NGINX
	target := "http://localhost:3001"
	r.NoRoute(reverseProxy(target))

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
