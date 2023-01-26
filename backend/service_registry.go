package backend

import (
	"fmt"
	"github.com/dop251/goja"
	"gormja_core2/jsutil"
)

type ServiceRegistry struct {
	runtime             *goja.Runtime
	namePrototypeMap    map[string]ClassTypeType
	prototypeServiceMap map[ClassTypeType]*ServiceObject
}

type ClassTypeType goja.Value

func NewServiceRegistry(runtime *goja.Runtime) *ServiceRegistry {
	return &ServiceRegistry{
		runtime:             runtime,
		namePrototypeMap:    make(map[string]ClassTypeType),
		prototypeServiceMap: make(map[ClassTypeType]*ServiceObject),
	}
}

func (x *ServiceRegistry) Get(classType ClassTypeType) *ServiceObject {
	return x.prototypeServiceMap[classType]
}

func (x *ServiceRegistry) LookupByName(name string) (*ServiceObject, error) {
	classType, ok := x.namePrototypeMap[name]
	if !ok {
		return nil, fmt.Errorf("service with name:%s does not exist", name)
	}
	return x.Get(classType), nil
}

func (x *ServiceRegistry) Put(classType ClassTypeType, wrapperObj *ServiceObject) {
	var className = jsutil.GetClassName(classType.ToObject(x.runtime))
	x.namePrototypeMap[className] = classType
	x.prototypeServiceMap[classType] = wrapperObj
}
