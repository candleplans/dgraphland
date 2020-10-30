package dgraphland

import "errors"

// errEmptyStruct : "เป็น Struct ที่ว่างเปล่า"
func errEmptyStruct(typeName string) error {
	return errors.New(`DgError: ` + typeName +
		` Type -> "เป็น Struct ที่ว่างเปล่า"`)
}

// errNoModelStruct : "ไม่มีการเรียกใช้ Model Struct หลักของ Dgraphland"
func errNoModelStruct(typeName string) error {
	return errors.New(`DgError: ` + typeName +
		` Type -> "ไม่มีการเรียกใช้ ` + ModelTypeName + ` Struct หลักของ Dgraphland"`)
}
