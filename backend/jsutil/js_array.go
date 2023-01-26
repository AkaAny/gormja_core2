package jsutil

import (
	"fmt"
	"github.com/dop251/goja"
)

func IsJSArray(runtime *goja.Runtime, value goja.Value) bool {
	var array = runtime.Get("Array").ToObject(runtime)
	var isArrayFuncValue = array.Get("isArray")
	isArrayFunc, _ := goja.AssertFunction(isArrayFuncValue)
	ret, err := isArrayFunc(runtime.GlobalObject(), value)
	if err != nil {
		panic("unreachable")
	}
	return ret.ToBoolean()
}

func ToGoArray[T any](arrValue goja.Value) ([]T, error) {
	var goArr = arrValue.Export().([]interface{})
	var result = make([]T, 0)
	for i, arrItem := range goArr {
		item, ok := arrItem.(T)
		if !ok {
			return nil, fmt.Errorf("cannot convert index:%d", i)
		}
		result = append(result, item)
	}
	return result, nil
}
