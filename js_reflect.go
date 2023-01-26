package gormja_core2

import (
	"fmt"
	"github.com/dop251/goja"
)

type JSReflectMirror struct {
	runtime    *goja.Runtime
	reflectObj *goja.Object
}

func NewJSReflectMirror(runtime *goja.Runtime) *JSReflectMirror {
	var reflectObj = runtime.Get("Reflect").ToObject(runtime)
	return &JSReflectMirror{
		runtime:    runtime,
		reflectObj: reflectObj,
	}
}

func (x JSReflectMirror) GetOwnKeys(obj goja.Value) ([]string, error) {
	var callable = MustAssertJSMemberFunc(x.reflectObj, "ownKeys")
	keysJSValue, err := callable(x.runtime.GlobalObject(), obj)
	if err != nil {
		return nil, fmt.Errorf("call ownKeys with err:%w", err)
	}
	var keyInterfaceArray = keysJSValue.Export().([]interface{})
	var keys = make([]string, 0)
	for _, keyInterface := range keyInterfaceArray {
		var key = keyInterface.(string)
		keys = append(keys, key)
	}
	return keys, nil
}

func (x JSReflectMirror) GetMetadata(metaDataKey, obj, propertyKey goja.Value) (goja.Value, error) {
	var callable = MustAssertJSMemberFunc(x.reflectObj, "getMetadata")
	attrObj, err := callable(x.runtime.GlobalObject(),
		metaDataKey, obj, propertyKey)
	if err != nil {
		return nil, fmt.Errorf("call ownKeys with err:%w", err)
	}
	return attrObj, nil
}
