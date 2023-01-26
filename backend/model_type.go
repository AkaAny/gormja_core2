package backend

import (
	"fmt"
	"github.com/dop251/goja"
	"gormja_core2/jsutil"
	"gormja_core2/mirror"
	"reflect"
	"strings"
)

var dbTypeReflectTypeMap = map[string]reflect.Type{
	"string":  reflect.TypeOf(""),
	"int64":   reflect.TypeOf(int64(0)),
	"bool":    reflect.TypeOf(true),
	"float64": reflect.TypeOf(float64(0.0)),
}

func NewModelReflectType(modelObj *goja.Object, runtime *goja.Runtime) reflect.Type {
	var className = jsutil.GetInstanceClassName(modelObj, runtime)
	fmt.Println(className)
	var dbFieldAnnoHandle = NewTSAnnotation("dbField", runtime)
	keyAttrMap, err := dbFieldAnnoHandle.GetObject(modelObj)
	if err != nil {
		panic(fmt.Errorf("get obj db field attr with err:%w", err))
	}
	var reflectFields = make([]reflect.StructField, 0)
	for key, attrMap := range keyAttrMap {
		var (
			dbName       = attrMap["column"].(string)
			dbTypeStr    = attrMap["dbType"].(string)
			isPrimaryKey = attrMap["isPrimaryKey"].(bool)
		)
		var tagItems = make([]string, 0)
		//column
		tagItems = append(tagItems, "column:"+dbName)
		//primary key
		if isPrimaryKey {
			tagItems = append(tagItems, "primaryKey")
		}
		//dbType
		var reflectType = dbTypeReflectTypeMap[dbTypeStr]
		var tagContentStr = strings.Join(tagItems, ";")
		var reflectFieldItem = reflect.StructField{
			Name:      key,
			PkgPath:   "",
			Type:      reflectType,
			Tag:       reflect.StructTag(fmt.Sprintf(`gorm:"%s"`, tagContentStr)),
			Anonymous: false,
		}
		reflectFields = append(reflectFields, reflectFieldItem)
	}
	var modelGoType = reflect.StructOf(reflectFields)
	mirror.SetTypeName(modelGoType, className)
	return modelGoType
}

func GetModelCacheKey(modelObj *goja.Object, runtime *goja.Runtime) string {
	var className = jsutil.GetInstanceClassName(modelObj, runtime)
	return className
}
