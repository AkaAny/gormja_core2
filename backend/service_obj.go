package backend

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dop251/goja"
	"gormja_core2/jsutil"
	"reflect"
)

type ServiceObject struct {
	runtime              *goja.Runtime
	classType            ClassTypeType
	instanceObj          *goja.Object
	serviceID            string
	unifyEntityModelType reflect.Type
}

func (x *ServiceObject) getCacheKey(condMap map[string]interface{}) string {
	//hash query obj json
	rawData, err := json.Marshal(condMap)
	if err != nil {
		panic(err)
	}
	var toHashData = rawData
	var hashData = sha1.New().Sum(toHashData)
	var hashStr = hex.EncodeToString(hashData)
	var className = jsutil.GetInstanceClassName(x.instanceObj, x.runtime)
	var cacheKey = fmt.Sprintf("lookup::%s_%s", className, hashStr)
	return cacheKey
}

func (x *ServiceObject) Lookup(ctx context.Context, condMap map[string]interface{}, cacheClient CacheClient) ([]interface{}, error) {
	var cacheKey = x.getCacheKey(condMap)
	dests, cacheExist, err := cacheClient.Get(ctx, cacheKey)
	if err != nil {
		return nil, fmt.Errorf("get from cache with err:%w", err)
	}
	if cacheExist {
		return dests, nil
	}
	//cache miss, fallback to db
	dests, err = x.forceGetFromServiceAndCacheTo(ctx, condMap, cacheKey, cacheClient)
	if err != nil {
		return dests, err
	}
	return dests, nil
}

func (x *ServiceObject) forceGetFromServiceAndCacheTo(ctx context.Context, condMap map[string]interface{},
	cacheKey string, client CacheClient) ([]interface{}, error) {
	var lookupCallable = jsutil.MustAssertJSMemberFunc(x.instanceObj, "lookup")
	var queryObjJSValue = x.runtime.ToValue(condMap)
	destsJSValue, err := lookupCallable(x.instanceObj, queryObjJSValue)
	if err != nil {
		return nil, fmt.Errorf("lookup from service with err:%w", err)
	}
	exportedGoArray, err := jsutil.ToGoArray[interface{}](destsJSValue)
	if err != nil {
		return nil, fmt.Errorf("to go array with err:%w", err)
	}
	//just prevent someone from giving us a strange value to leak all data
	var dests = exportedGoArray //make([]interface{}, 0)
	//for _, exportedItem := range exportedGoArray {
	//	var destValue = reflect.New(x.unifyEntityModelType).Elem()
	//	var exportItemValue = reflect.ValueOf(exportedItem) //that is a map
	//	for fieldIndex := 0; fieldIndex < x.unifyEntityModelType.NumField(); fieldIndex++ {
	//		var structField = x.unifyEntityModelType.Field(fieldIndex)
	//		var fieldNameValue = reflect.ValueOf(structField.Name)
	//		var exportFieldValue = exportItemValue.MapIndex(fieldNameValue)
	//		destValue.Field(fieldIndex).Set(exportFieldValue)
	//	}
	//	var dest = destValue.Interface()
	//	dests = append(dests, dest)
	//}
	var ttlInMilliSecond = x.instanceObj.Get("ttl").ToInteger()
	if err := client.Set(ctx, cacheKey, dests, ttlInMilliSecond); err != nil {
		return dests, fmt.Errorf("set to cache with err:%w", err)
	}
	return dests, nil
}

func (x *ServiceObject) ManualRefresh(ctx context.Context, condMap map[string]interface{}, cacheClient CacheClient) error {
	var cacheKey = x.getCacheKey(condMap)
	_, err := x.forceGetFromServiceAndCacheTo(ctx, condMap, cacheKey, cacheClient)
	if err != nil {
		return err
	}
	return nil
}
