package dgraphland

import (
	"context"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

const target = "localhost:9080"

// ทดสอบการเชื่อมต่อ
func TestConnect(t *testing.T) {
	dg, cancel := GetDgraphClient(target)
	defer cancel()
	op := api.Operation{DropAll: true}
	ctx := context.Background()
	err := dg.Alter(ctx, &op)
	assert.NoError(t, err)
}

// ทดสอบการยกเลิกการเชื่อมต่อ
func TestDisconnect(t *testing.T) {
	dg, cancel := GetDgraphClient(target)
	op := api.Operation{DropAll: true}
	ctx := context.Background()
	cancel()
	err := dg.Alter(ctx, &op)
	assert.EqualError(t, err, "rpc error: code = Canceled desc = grpc: the client connection is closing")
}
