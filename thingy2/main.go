package main

import (
	"../../../tmp/demo/thingy2"
	"../../../tmp/demo/thingy2/server"
	"../../../tmp/demo/thingy2/thingy2/cmd"
	"github.com/lileio/lile"
	"github.com/lileio/lile/fromenv"
	"github.com/lileio/pubsub"
	"google.golang.org/grpc"
)

func main() {
	s := &server.Thingy2Server{}

	lile.Name("thingy2")
	lile.Server(func(g *grpc.Server) {
		thingy2.RegisterThingy2Server(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
	})

	cmd.Execute()
}
