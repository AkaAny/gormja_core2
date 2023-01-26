package jsutil

import "github.com/dop251/goja"

func GetClassName(constructor *goja.Object) string {
	return constructor.Get("name").String()
}

func GetInstanceClassName(objJSValue *goja.Object, runtime *goja.Runtime) string {
	var constructorObj = getConstructor(objJSValue, runtime)
	var name = GetClassName(constructorObj)
	return name
}

func getConstructor(obj *goja.Object, runtime *goja.Runtime) *goja.Object {
	var constructorObj = obj.Get("constructor").ToObject(runtime)
	return constructorObj
}

func GetInstanceStaticMember(objJSValue *goja.Object, key string, runtime *goja.Runtime) goja.Value {
	var constructorObj = getConstructor(objJSValue, runtime)
	var jsValue = constructorObj.Get(key)
	return jsValue
}
