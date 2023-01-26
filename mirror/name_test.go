package mirror

import (
	"reflect"
	"testing"
)

func TestSetTypeName(t *testing.T) {
	var fields = []reflect.StructField{
		reflect.StructField{
			Name:      "StaffID",
			Type:      reflect.TypeOf(""),
			Tag:       "",
			Anonymous: false,
		},
		reflect.StructField{
			Name:      "Name",
			Type:      reflect.TypeOf(""),
			Tag:       "",
			Anonymous: false,
		},
	}
	var madeType = reflect.StructOf(fields)
	SetTypeName(madeType, "PersonInfo")
}
