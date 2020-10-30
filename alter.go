package dgraphland

import (
	"github.com/dgraph-io/dgo/v200/protos/api"
)

// DropAll ใช้ล้างข้อมูลใน Dgraph Database ทั้งหมด
func DropAll() error {
	err := Client.Alter(ctx(),
		&api.Operation{
			DropAll: true,
		})
	return err
}

// AutoMigrate ใช้กำหนด Dgraph Schema อัตโนมัติ
func AutoMigrate(models ...interface{}) error {
	schema, err := Schema(models)
	if err != nil {
		return err
	}
	err = Client.Alter(ctx(),
		&api.Operation{
			Schema: schema,
		})
	return err
}
