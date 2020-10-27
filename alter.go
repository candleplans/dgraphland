package dgraphland

import (
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"strings"
)

func SetSchema(dg *dgo.Dgraph, schemas []string) error {
	err := dg.Alter(ctx(),
		&api.Operation{
			Schema: strings.Join(schemas, ""),
		})
	return err
}
