package utils

import (
	"gorm.io/gorm"
	"reflect"
	"unsafe"
)

func CallDBFunc(dbFunc reflect.Value, argValues []reflect.Value) *gorm.DB {
	var newTxValue = dbFunc.Call(argValues)[0]
	var tx = (*gorm.DB)(unsafe.Pointer(newTxValue.Pointer()))
	return tx
}
