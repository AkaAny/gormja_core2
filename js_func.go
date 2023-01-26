package gormja_core2

import (
	"fmt"
	"github.com/dop251/goja"
)

func MustAssertJSMemberFunc(parentObj interface {
	Get(name string) goja.Value
}, name string) goja.Callable {
	var jsFuncValue = parentObj.Get(name)
	callable, ok := goja.AssertFunction(jsFuncValue)
	if !ok {
		panic(fmt.Errorf("key:%s is not a js func, might be a invalid script", name))
	}
	return callable
}

func MustAssertJSFunc(funcJSValue goja.Value) goja.Callable {
	callable, ok := goja.AssertFunction(funcJSValue)
	if !ok {
		panic(fmt.Errorf("value:%v is not a js func, might be a invalid script", funcJSValue))
	}
	return callable
}
