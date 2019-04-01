package atlas

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// FillDefinition set a few parameters in a definition list
/*
The goal here is to give a dictionary of string and let it figure out each field's type
depending on the recipient's type in the struct.
*/
func FillDefinition(d *Definition, fields map[string]string) error {
	if d == nil {
		return nil
	}
	sdef := reflect.ValueOf(d).Elem()
	typeOfDef := sdef.Type()
	for k, v := range fields {
		// Check the field is present
		if f, ok := typeOfDef.FieldByName(k); ok {
			// Use the right type
			switch f.Type.Name() {
			case "float":
				vf, _ := strconv.ParseFloat(v, 32)
				sdef.FieldByName(k).SetFloat(vf)
			case "int":
				vi, _ := strconv.Atoi(v)
				sdef.FieldByName(k).SetInt(int64(vi))
			case "string":
				sdef.FieldByName(k).SetString(v)
			case "bool":
				vb, _ := strconv.ParseBool(v)
				sdef.FieldByName(k).SetBool(vb)
			case "":
				if f.Type.Kind().String() == "slice" {
					// Special case for "tags" which is an array, not a scalar
					a := strings.Split(v, ",")
					sdef.FieldByName(k).Set(reflect.ValueOf(a))
				}
			default:
				return fmt.Errorf("Unsupported type: %s", f.Type.Name())
			}
		}
	}
	return nil
}
