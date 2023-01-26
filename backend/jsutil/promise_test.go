package jsutil

import (
	"fmt"
	"github.com/dop251/goja"
	"testing"
)

func TestPromise(t *testing.T) {
	var runtime = goja.New()
	err := runtime.Set("func1", func(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
		promise, resolve, reject := runtime.NewPromise()
		var resolveOrReject = call.Argument(0).ToBoolean()
		if resolveOrReject {
			resolve("result from go") //make promise sync
		} else {
			reject(fmt.Errorf("err from go:%s", "test"))
		}
		return runtime.ToValue(promise)
	})
	if err != nil {
		panic(err)
	}
	err = runtime.Set("print", func(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
		var arg0Val = call.Argument(0)
		fmt.Println(arg0Val.Export(), arg0Val.ExportType())

		return goja.Null()
	})
	if err != nil {
		panic(err)
	}
	runtime.RunString(`
		print("a")
		func1(true).then((result)=>{print(result)})
		func1(false).catch((result)=>{print(result)})
`)

	for {
	}
}
