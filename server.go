package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pramishkarkee/gql-test-struct/graph"
	"github.com/pramishkarkee/gql-test-struct/graph/generated"
	"github.com/joho/godotenv"
	"github.com/pramishkarkee/gql-test-struct/graph/db"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	err := godotenv.Load(); if err != nil {
		log.Fatal("Error loading .env file")
	   }
	Database := db.Connect()
	fmt.Println("database",Database)
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: Database}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
