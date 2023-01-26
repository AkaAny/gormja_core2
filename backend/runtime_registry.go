package backend

import "fmt"

type RuntimeRegistry struct {
	idRuntimeMap map[string]*ServiceRuntime
}

func (x *RuntimeRegistry) Put(id string, serviceRuntime *ServiceRuntime) {
	x.idRuntimeMap[id] = serviceRuntime
}

func (x *RuntimeRegistry) Get(id string) (*ServiceRuntime, error) {
	serviceRuntime, ok := x.idRuntimeMap[id]
	if !ok {
		return nil, fmt.Errorf("service runtime with id:%s does not exist", id)
	}
	return serviceRuntime, nil
}

func (x *RuntimeRegistry) List() {

}
