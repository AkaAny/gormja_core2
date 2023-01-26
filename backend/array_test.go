package backend

import (
	"fmt"
	"github.com/dop251/goja"
	"testing"
)

func TestIsArray(t *testing.T) {
	var runtime = goja.New()
	runtime.RunString(`
	const arr1=['a','b','c'];
`)
	var arr1 = runtime.Get("arr1")
	var arr1Keys = arr1.ToObject(runtime).Keys()
	fmt.Println(arr1Keys)
	var array = runtime.Get("Array").ToObject(runtime)
	var isArrayFuncValue = array.Get("isArray")
	isArrayFunc, ok := goja.AssertFunction(isArrayFuncValue)
	if !ok {
		panic("not goja callable")
	}
	ret, err := isArrayFunc(goja.Null(), arr1)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret.ToBoolean())
	var exportedArr1 = arr1.ToObject(runtime).Export()
	fmt.Println(exportedArr1)
}
