package gormja_core2

import (
	"fmt"
	"github.com/dop251/goja"
)

type TSAnnotation struct {
	runtime     *goja.Runtime
	metaDataKey goja.Value
}

type AnnotationAttrMap map[string]interface{}

func NewTSAnnotation(metaDataKeyStr string, runtime *goja.Runtime) *TSAnnotation {
	var metaDataKeyJSValue = runtime.ToValue(metaDataKeyStr)
	return &TSAnnotation{
		runtime:     runtime,
		metaDataKey: metaDataKeyJSValue,
	}
}

func (x TSAnnotation) GetProperty(obj goja.Value, propertyKey string) goja.Value {
	var reflectMirror = NewJSReflectMirror(x.runtime)
	var propertyKeyJSValue = x.runtime.ToValue(propertyKey)
	annoAttrMapJsValue, err := reflectMirror.GetMetadata(x.metaDataKey, obj, propertyKeyJSValue)
	if err != nil {
		panic(fmt.Errorf("get anno metadata with err:%w", err))
	}
	return annoAttrMapJsValue
}

func (x TSAnnotation) GetObject(obj goja.Value) (map[string]AnnotationAttrMap, error) {
	var reflectMirror = NewJSReflectMirror(x.runtime)
	keys, err := reflectMirror.GetOwnKeys(obj)
	if err != nil {
		return nil, fmt.Errorf("get own keys with err:%w", err)
	}
	var keyAnnoAttrMap = make(map[string]AnnotationAttrMap)
	for _, key := range keys {
		var annoAttrMap = x.GetProperty(obj, key).Export().(map[string]interface{})
		keyAnnoAttrMap[key] = annoAttrMap
	}
	return keyAnnoAttrMap, nil
}
