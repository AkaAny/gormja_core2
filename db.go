package gormja_core2

import (
	"fmt"
	"github.com/dop251/goja"
	"gorm.io/gorm"
	"gormja_core2/utils"
	"reflect"
)

type JSDB struct {
	db                    *gorm.DB
	cachedEntityModelType map[string]reflect.Type
}

func NewJSDB(db *gorm.DB) *JSDB {
	return &JSDB{db: db}
}

func (x *JSDB) ToJSValue(runtime *goja.Runtime) goja.Value {
	var obj = runtime.NewObject()
	if err := obj.Set("startSession", x.NewSession); err != nil {
		panic(fmt.Errorf("set startSession with err:%w", err))
	}
	return obj
}

func (x *JSDB) NewSession(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var modelClassType = call.Argument(0)
	modelObj, err := runtime.New(modelClassType, runtime.NewObject())
	if err != nil {
		panic(fmt.Errorf("new model obj with err:%w", err))
	}
	var typeCacheKey = GetModelCacheKey(modelObj, runtime)
	modelGoType, ok := x.cachedEntityModelType[typeCacheKey]
	if !ok {
		modelGoType = NewModelReflectType(modelObj, runtime)
	}
	//table name
	var tableNameFuncJSValue = GetInstanceStaticMember(modelObj, "tableName", runtime)
	var tableNameCallable = MustAssertJSFunc(tableNameFuncJSValue)
	tableNameJSValue, err := tableNameCallable(modelObj)
	if err != nil {
		panic(fmt.Errorf("get table name with err:%w", err))
	}
	var tableName = tableNameJSValue.String()
	//go model value
	var goModelValue = reflect.New(modelGoType)
	var goModel = goModelValue.Interface()
	var tx = x.db.Debug().Model(goModel).Table(tableName)
	//promise, resolve, reject := runtime.NewPromise()
	//var tableName = call.Argument(0).String()
	var jsDBSession = NewJSDBSession(tx, modelGoType)
	return jsDBSession.ToJSValue(runtime)
	//return runtime.ToValue(promise)
}

type JSDBSession struct {
	tx        *gorm.DB
	modelType reflect.Type
}

func NewJSDBSession(tx *gorm.DB, modelType reflect.Type) *JSDBSession {
	//var tx = db.Table()
	return &JSDBSession{
		tx:        tx,
		modelType: modelType,
	}
	//return nil, nil
}

func (x *JSDBSession) ToJSValue(runtime *goja.Runtime) goja.Value {
	var obj = runtime.NewObject()
	obj.Set("where", x.Where)
	obj.Set("find", x.Find)
	return obj
}

func (x *JSDBSession) Where(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var queryObj = call.Arguments[0]
	var query = queryObj.Export()
	var argJSValues = call.Arguments[1:]
	var args = make([]interface{}, 0)
	for _, argJSValue := range argJSValues {
		var arg = argJSValue.Export()
		args = append(args, arg)
	}
	x.tx = x.tx.Where(query, args...)
	return call.This
}

func (x *JSDBSession) Find(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var sliceType = reflect.SliceOf(x.modelType)
	var destPtrValue = reflect.New(sliceType)
	//mirror.SliceValueAddr()
	utils.CallDBFunc(reflect.ValueOf(x.tx.Find), []reflect.Value{destPtrValue})
	var dest = destPtrValue.Elem().Interface()
	var destJSValue = runtime.ToValue(dest)
	fmt.Println(dest, destJSValue)
	return destJSValue
}
