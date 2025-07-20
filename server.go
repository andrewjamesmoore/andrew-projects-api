package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andrewjamesmoore/andrew-projects-api/database"
	"github.com/andrewjamesmoore/andrew-projects-api/graph"
	"github.com/andrewjamesmoore/andrew-projects-api/middleware"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {

	log.Println("SECRET_KEY is:", os.Getenv("SECRET_KEY"))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.Connect()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{DB: db},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

		http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", middleware.WithSecretHeader(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
