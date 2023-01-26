package utils

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
)

type JSConsole struct {
	logger *logrus.Logger
}

func NewJSConsole(logger *logrus.Logger) *JSConsole {
	return &JSConsole{logger: logger}
}

func (x *JSConsole) Attach(runtime *goja.Runtime) {
	var consoleObj = runtime.NewObject()
	if err := consoleObj.Set("log", x.Log); err != nil {
		panic(fmt.Errorf("set log to console obj with err:%w", err))
	}
	if err := runtime.GlobalObject().Set("console", consoleObj); err != nil {
		panic(fmt.Errorf("set console to global obj with err:%w", err))
	}
}

func (x *JSConsole) Log(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var args = make([]interface{}, 0)
	args = append(args, "[console.log]")
	for _, argJSValue := range call.Arguments {
		var argItem = argJSValue.Export()
		args = append(args, argItem)
	}
	x.logger.Println(args...)
	return goja.Undefined()
}
