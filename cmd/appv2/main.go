package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/caarlos0/env/v11"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/wisaitas/graphql-golang/internal/appv2"
	"github.com/wisaitas/graphql-golang/internal/appv2/graph"
	"github.com/wisaitas/graphql-golang/internal/appv2/initial"
)

func init() {
	if err := env.Parse(&appv2.ENV); err != nil {
		log.Fatal(err)
	}
}

func main() {
	resolver := initial.NewInitial()
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", appv2.ENV.Server.Port)
	log.Fatal(http.ListenAndServe(":"+appv2.ENV.Server.Port, nil))
}
