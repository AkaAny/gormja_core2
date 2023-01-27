package backend

import (
	"context"
	"fmt"
	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"gormja_core2/jsutil"
)

func NewServiceRuntime(dbRegistry *DBRegistry, cacheClient CacheClient, logger *logrus.Logger) *ServiceRuntime {
	var runtime = goja.New()
	var serviceRuntime = &ServiceRuntime{
		runtime:         runtime,
		logger:          logger,
		dbRegistry:      dbRegistry,
		cacheClient:     cacheClient,
		serviceRegistry: NewServiceRegistry(runtime),
	}
	return serviceRuntime
}

type ServiceRuntime struct {
	runtime         *goja.Runtime
	logger          *logrus.Logger
	dbRegistry      *DBRegistry
	cacheClient     CacheClient
	serviceRegistry *ServiceRegistry
}

func (x *ServiceRuntime) Init() {
	jsutil.NewJSConsole(x.logger).Attach(x.runtime)
	var runtimeObj = x.runtime.NewObject()
	if err := runtimeObj.Set("registerService", x.RegisterService); err != nil {
		panic(err)
	}
	if err := runtimeObj.Set("getService", x.GetService); err != nil {
		panic(err)
	}
	if err := runtimeObj.Set("debugBreakpoint", x.DebugBreakPoint); err != nil {
		panic(err)
	}
	if err := x.runtime.Set("Runtime", runtimeObj); err != nil {
		panic(err)
	}
}

func (x *ServiceRuntime) LoadScript(content string) {
	_, err := x.runtime.RunString(content)
	if err != nil {
		panic(err)
	}
}

func (x *ServiceRuntime) Lookup(ctx context.Context, serviceID string,
	condMap map[string]interface{}) (dests []interface{}, err error) {
	serviceObj, err := x.serviceRegistry.LookupByName(serviceID)
	if err != nil {
		return nil, err
	}
	dests, err = serviceObj.Lookup(ctx, condMap, x.cacheClient)
	if err != nil {
		return nil, fmt.Errorf("lookup on service:%s with err:%w", serviceID, err)
	}
	return dests, nil
}

func (x *ServiceRuntime) ManualRefresh(ctx context.Context, serviceName string,
	condMap map[string]interface{}) error {
	serviceObj, err := x.serviceRegistry.LookupByName(serviceName)
	if err != nil {
		return err
	}
	if err := serviceObj.ManualRefresh(ctx, condMap, x.cacheClient); err != nil {
		return fmt.Errorf("manual refresh on service:%s with err:%w", serviceName, err)
	}
	return nil
}

func (x *ServiceRuntime) RegisterService(call goja.FunctionCall) goja.Value {
	promise, resolve, reject := x.runtime.NewPromise()
	var prototype = call.Argument(0).ToObject(x.runtime)
	instanceObj, err := x.runtime.New(prototype)
	if err != nil {
		var goErr = fmt.Errorf("new service:%v with err:%w", prototype, err)
		x.logger.WithError(err).Errorln("err when new service instance")
		reject(goErr)
		return x.runtime.ToValue(promise)
	}
	var serviceID = instanceObj.Get("serviceID").String()
	//var instanceObj = call.Argument(0).ToObject(x.runtime)
	var sourceType = instanceObj.Get("sourceType").String()
	fmt.Println("source type:", sourceType)
	if err := x.registerDBService(instanceObj); err != nil {
		reject(err)
		return x.runtime.ToValue(promise)
	}
	var newUnifyModelCallable = jsutil.MustAssertJSMemberFunc(instanceObj, "newUnifyModel")
	unifyModelJSValue, err := newUnifyModelCallable(instanceObj)
	if err != nil {
		reject(err)
		return x.runtime.ToValue(promise)
	}
	var unifyModelObj = unifyModelJSValue.ToObject(x.runtime)
	var unifyModelType = NewModelReflectType(unifyModelObj, x.runtime)
	var wrapperObj = &ServiceObject{
		runtime:              x.runtime,
		classType:            prototype,
		instanceObj:          instanceObj,
		serviceID:            serviceID,
		unifyEntityModelType: unifyModelType,
	}
	x.serviceRegistry.Put(wrapperObj)
	resolve(instanceObj)
	return x.runtime.ToValue(promise)
}

func (x *ServiceRuntime) registerDBService(instanceObj *goja.Object) error {
	var dbInitParamObj = instanceObj.Get("dbInitParam").ToObject(x.runtime)
	fmt.Println("db init param:", dbInitParamObj.Export())
	var dataSourceID = dbInitParamObj.Get("dataSourceID").String()
	fmt.Println("data source id:", dataSourceID)
	var db = x.dbRegistry.Get(dataSourceID)
	if db == nil {
		var goErr = fmt.Errorf("db with id:%s does not exist", dataSourceID)
		x.logger.WithError(goErr).Errorln("err when get db from registry")
		return goErr
	}
	var dbJSValue = NewJSDB(db).ToJSValue(x.runtime)
	var initFunc = jsutil.MustAssertJSMemberFunc(instanceObj, "init")
	fmt.Println("init func:", initFunc)
	_, err := initFunc(instanceObj, dbJSValue)
	if err != nil {
		var goErr = fmt.Errorf("call init with err:%w", err)
		x.logger.WithError(goErr).Errorln("err when call init on js service instance")
		return goErr
	}
	return nil
}

func (x *ServiceRuntime) GetService(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var classTypeJSValue = call.Argument(0)
	var wrapperObj = x.serviceRegistry.Get(classTypeJSValue)
	return wrapperObj.instanceObj
}

func (x *ServiceRuntime) GetServiceByID(serviceID string) (*ServiceObject, error) {
	return x.serviceRegistry.LookupByName(serviceID)
}

func (x *ServiceRuntime) DebugBreakPoint(call goja.FunctionCall, runtime *goja.Runtime) goja.Value {
	var hint = call.Argument(0).String()
	fmt.Println("breakpoint hit:", hint)
	return goja.Undefined()
}
