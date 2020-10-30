package dgraphland

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"log"
)

var Client *dgo.Dgraph

type CancelFunc func()

// SetupClient ใช้สำหรับการเชื่อมต่อฐานข้อมูล Dgraph ด้วย GRPC
func SetupClient(target string) CancelFunc {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("While trying to dial gRPC: %v", err)
	}

	dc := api.NewDgraphClient(conn)
	Client = dgo.NewDgraphClient(dc)

	return func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
