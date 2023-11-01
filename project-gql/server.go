package main

import (
	"log"
	"net/http"
	"os"
	"project-gql/database"
	"project-gql/graph"
	"project-gql/repository"
	"project-gql/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := database.Connection()
	if err != nil {
		log.Println("db connection %w", err)
		return
	}

	repo, _ := repository.NewRepo(db)

	se, _ := service.NewService(repo, repo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: se}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
