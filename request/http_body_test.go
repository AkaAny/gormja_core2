package request

import (
	"fmt"
	"github.com/dop251/goja"
	"testing"
)

func TestArrayBuffer(t *testing.T) {
	var runtime = goja.New()
	runtime.RunString(`
		var arrBuf1=new ArrayBuffer([1,2,3,4,5,6])
		console.log(arrBuf1)
`)
	var arrBuf1 = runtime.Get("arrBuf1").ToObject(runtime)
	var buf = arrBuf1.Export()
	fmt.Println(buf)
}
