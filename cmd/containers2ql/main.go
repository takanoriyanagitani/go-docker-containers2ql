package main

import (
	"log"
	"net/http"
	"os"

	gh "github.com/99designs/gqlgen/graphql/handler"
	ge "github.com/99designs/gqlgen/graphql/handler/extension"
	gt "github.com/99designs/gqlgen/graphql/handler/transport"
	dc "github.com/docker/docker/client"
	"github.com/takanoriyanagitani/go-docker-containers2ql/graph"
)

func sub() error {
	var addrPort string = os.Getenv("ADDR_PORT")

	cli, e := dc.NewClientWithOpts(dc.FromEnv)
	if nil != e {
		return e
	}

	var res *graph.Resolver = &graph.Resolver{Client: cli}

	srv := gh.New(
		graph.NewExecutableSchema(graph.Config{Resolvers: res}),
	)

	srv.AddTransport(gt.GET{})
	srv.AddTransport(gt.POST{})

	srv.Use(ge.Introspection{})

	router := http.NewServeMux()

	router.Handle("/query", srv)

	return http.ListenAndServe(addrPort, router)
}

func main() {
	e := sub()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
