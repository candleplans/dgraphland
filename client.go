package dgraphland

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"log"
)

type CancelFunc func()

// GetDgraphClient ใช้สำหรับการเชื่อมต่อฐานข้อมูล Dgraph ด้วย GRPC
func GetDgraphClient(target string) (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("While trying to dial gRPC: %v", err)
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
