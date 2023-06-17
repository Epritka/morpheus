package types

type Entity interface {
	Type() string
	Labels() []string
	SetElementId(elementId string)
	GetElementId() string
	SetId(int64)
	GetId() int64
	Properies() map[string]any
}

type BaseEntity struct {
	Id        int64
	ElementId string
}

func new() *BaseEntity {
	return &BaseEntity{}
}

func (b *BaseEntity) SetElementId(elementId string) {
	b.ElementId = elementId
}

func (b *BaseEntity) GetElementId() string {
	return b.ElementId
}

func (b *BaseEntity) SetId(id int64) {
	b.Id = id
}

func (b *BaseEntity) GetId() int64 {
	return b.Id
}
