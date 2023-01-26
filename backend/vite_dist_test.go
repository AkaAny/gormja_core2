package backend

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"gormja_core2/utils"
	"io/fs"
	"os"
	"testing"
)

func GetScriptContent() string {
	var dirFS = os.DirFS("../my-ts-lib")
	const pathPrefix = "dist/assets"
	entries, err := fs.ReadDir(dirFS, pathPrefix)
	if err != nil {
		panic(err)
	}
	var jsFileName = entries[0].Name()
	fmt.Println("js file:", jsFileName)
	rawData, err := fs.ReadFile(dirFS, pathPrefix+"/"+jsFileName)
	if err != nil {
		panic(err)
	}
	return string(rawData)
}

func TestLoad(t *testing.T) {
	var scriptContent = GetScriptContent()
	//fmt.Println(string(rawData))
	var runtime = goja.New()
	var jsConsole = utils.NewJSConsole(logrus.New())
	jsConsole.Attach(runtime)
	_, err := runtime.RunString(scriptContent)
	if err != nil {
		//panic(err)
	}
	var rankPrototype = runtime.Get("Rank").ToObject(runtime)
	var tableNameFuncJSValue = MustAssertJSMemberFunc(rankPrototype, "tableName")
	fmt.Println(tableNameFuncJSValue(rankPrototype))

	var constructorJSValue = rankPrototype.Get("constructor")
	_, ok := goja.AssertConstructor(constructorJSValue)
	if !ok {
		panic("func is not a constructor")
	}
	//var propsJSValue = runtime.ToValue(map[string]interface{}{
	//	"schoolCode": 1,
	//	"staffID":    "20113128",
	//})
	//rankInstance, err := runtime.New(rankPrototype, propsJSValue)
	//if err != nil {
	//	panic(err)
	//}
	var rankInstance = runtime.Get("RankInstance")
	var className = GetInstanceClassName(rankInstance.ToObject(runtime), runtime)
	fmt.Println(className)
	var tableNameOnInstance = GetInstanceStaticMember(rankInstance.ToObject(runtime), "tableName", runtime)
	fmt.Println(tableNameOnInstance)
	//var metadataKey = runtime.Get("metaDataKey")
	//var goNewMetadataKey = runtime.ToValue("dbField") //goja.NewSymbol("dbField") symbol is unique by addr
	var dbFieldAnnotationHandle = NewTSAnnotation("dbField", runtime)
	var attrMap = dbFieldAnnotationHandle.GetProperty(rankInstance, "staffID")
	fmt.Println("attr map from call Reflect from go:", attrMap)
	//var getDBField = MustAssertJSMemberFunc(runtime, "getDBField")
	//attrMap, err = getDBField(runtime.GlobalObject(), rankInstance, runtime.ToValue("staffID"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(attrMap.Export())
	//var jsReflectMirror = NewJSReflectMirror(runtime)
	keyAnnoAttrMap, err := dbFieldAnnotationHandle.GetObject(rankInstance)
	if err != nil {
		panic(err)
	}
	fmt.Println(keyAnnoAttrMap)
}
