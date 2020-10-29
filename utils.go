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

func Schema(models ...interface{}) string {
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

		if t.Name() == ModelTypeName {
			continue
		}

		numFields := t.NumField()

		for i := 0; i < numFields; i++ {
			field := t.Field(i)
			tagValueDglang := field.Tag.Get(TagDgland)
			tagValueJson := field.Tag.Get(TagJson)

			// หากเป็น Field ที่ไม่มี Json Tag หรือ Dglang Tag
			//ให้ทำการข้ามรายการ Model Struct นั้นไปเลย
			if tagValueDglang != "" && tagValueJson != "" {
				fieldName := strings.Split(tagValueJson, ",")[0]
				bFields.WriteString(fieldName + ": " + tagValueDglang + " .\n")

				if typeName == ModelTypeName {
					bModelTypes.WriteString("\t" + fieldName + "\n")
					if i == numFields-1 {
						bModelTypes.WriteString("}\n")
					}
				} else {
					if i == 1 {
						bTypes.WriteString("\ntype " + typeName + " {\n")
					}

					bTypes.WriteString("\t" + fieldName + "\n")

					if i == numFields-1 {
						bTypes.WriteString(bModelTypes.String())
					}
				}
			}
		}
	}
	return bFields.String() + bTypes.String()
}
