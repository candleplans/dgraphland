package dgraphland

import (
	"bytes"
	"context"
	"reflect"
	"strings"
)

const (
	TagDgland     = "dgland"
	TagJson       = "json"
	ModelTypeName = "Model"
)

// ctx สำหรับค่า Context.Background()
func ctx() context.Context {
	return context.Background()
}

func Schema(models ...interface{}) (string, error) {
	// แทรก Model{} เป็นรายการแรกใน Models []interface{}
	models = append([]interface{}{&Model{}}, models...)

	bFields := bytes.Buffer{}     //สำหรับเก็บ Dgraph Field Buffer
	bTypes := bytes.Buffer{}      // สำหรับเก็บ Dgraph Type  Buffer
	bModelTypes := bytes.Buffer{} // สำหรับเก็บ Model Buffer

	// ดึงแต่ละ Input Type Model ไปทำการสร้างคำสั่ง DQL
	// ที่ใช้สำหรับ Setup Schema ใน Dgraph Database
	for _, model := range models {
		v := reflect.ValueOf(model)
		t := reflect.Indirect(v).Type()

		typeName := t.Name()

		numFields := t.NumField()

		// หากเป็น Type Structure ที่ไม่มี Field Structure ใดๆเลย
		// จะ return err
		if numFields == 0 {
			return "", errEmptyStruct(typeName)
		}

		for i := 0; i < numFields; i++ {
			field := t.Field(i)
			tagDglangValue := field.Tag.Get(TagDgland)
			tagJsonValue := field.Tag.Get(TagJson)

			// หากเป็น Structure ที่ไม่มี Field Json Tag และ Field  Dglang Tag เลย
			// ให้ทำการข้ามรายการ Field Structure นั้นไปเลย
			if tagDglangValue != "" && tagJsonValue != "" {
				tagJsonFieldName := strings.Split(tagJsonValue, ",")[0]
				bFields.WriteString(tagJsonFieldName + ": " + tagDglangValue + " .\n")

				if typeName == ModelTypeName {
					bModelTypes.WriteString("\t" + tagJsonFieldName + "\n")
					if i == numFields-1 {
						bModelTypes.WriteString("}\n")
					}
				} else
				// หากมี Model Type ใน Type ปัจจุบัน
				// ระบบจึงจะเริ่มดำเนินการแปลง Struct นั้นๆ เป็น Dgraph Schema
				if t.Field(0).Name == ModelTypeName {
					// สร้างส่วนหัวของ dgraph type
					if i == 1 {
						bTypes.WriteString("\ntype " + typeName + " {\n")
					}

					// นำ Field ของ Struct มาเขียน Dgraph Field Type ลง bTypes
					bTypes.WriteString("\t" + tagJsonFieldName + "\n")

					// เมื่อถึง Field สุดท้าย ให้นำ bModelTypes มาต่อท้าย bTypes ในแต่ละ Type ของ Dgraph Schema
					if i == numFields-1 {
						bTypes.WriteString(bModelTypes.String())
					}
				} else {
					// หากเป็น Type Structure ที่ไม่มี Model Type ใน Type Structure ปัจจุบัน
					// จะ return err
					return "", errNoModelStruct(typeName)
				}
			}
		}
	}

	// ถ้าหาก bTypes หรือ Dgraph Type ไม่มี ก็จะ return ""
	if bTypes.String() == "" {
		return "", nil
	} else {
		return bFields.String() + bTypes.String(), nil
	}
}
