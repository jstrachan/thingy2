package thingy2

import (
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	cm     = &sync.Mutex{}
	Client Thingy2Client
)

func GetThingy2Client() (Thingy2Client, error) {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client, nil
	}

	logrus.Info("Creating thingy2 gRPC client")
	conn, err := grpc.Dial("thingy2:80", grpc.DialOption(grpc.WithInsecure()))
	if err != nil {
		return nil, err
	}

	cli := NewThingy2Client(conn)
	Client = cli
	return cli, nil
}
