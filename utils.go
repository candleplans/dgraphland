package dgraphland

import (
	"bytes"
	"context"
	"errors"
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

		// หากเป็น Type Structure ที่ไม่มี Field Structure ใดๆเลย ให้ข้าม
		if numFields == 0 {
			continue
		}

		for i := 0; i < numFields; i++ {
			field := t.Field(i)
			tagDglangValue := field.Tag.Get(TagDgland)
			tagJsonValue := field.Tag.Get(TagJson)

			// หากเป็น Field ที่ไม่มี Json Tag และ Dglang Tag
			// ให้ทำการข้ามรายการ Field Structure นั้นไปเลย
			if tagDglangValue != "" && tagJsonValue != "" {
				tagJsonFieldName := strings.Split(tagJsonValue, ",")[0]
				bFields.WriteString(tagJsonFieldName + ": " + tagDglangValue + " .\n")

				if typeName == ModelTypeName {
					bModelTypes.WriteString("\t" + tagJsonFieldName + "\n")
					if i == numFields-1 {
						bModelTypes.WriteString("}\n")
					}
				} else if t.Field(0).Name == ModelTypeName {
					if i == 1 {
						bTypes.WriteString("\ntype " + typeName + " {\n")
					}

					bTypes.WriteString("\t" + tagJsonFieldName + "\n")

					if i == numFields-1 {
						bTypes.WriteString(bModelTypes.String())
					}
				} else {
					return "", errors.New("Error : " + typeName +
						" Type -> This structure does not have The " +
						typeName + " nested structure at first field.")
				}
			}
		}
	}

	if bTypes.String() != "" {
		return bFields.String() + bTypes.String(), nil
	} else {
		return "", nil
	}
}
