package backend

type ServiceDescription struct {
	ServiceID             string
	ClassName             string
	UnifyModelDescription *ModelDescription
}

type ModelDescription struct {
	ClassName string
	Fields    []*FieldDescription
}

type FieldDescription struct {
	Name string
	Type string
}
