package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"../../../tmp/demo/thingy2"
	"github.com/lileio/lile"
)

var s = Thingy2Server{}
var cli thingy2.Thingy2Client

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		thingy2.RegisterThingy2Server(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = thingy2.NewThingy2Client(lile.TestConn(addr))

	os.Exit(m.Run())
}
