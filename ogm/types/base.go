package types

type Entity interface {
	SetId(id int64)
	isEntity() bool
}

type BaseEntity struct {
	ID *int64 `ogm:"id"`
}

func new() *BaseEntity {
	return &BaseEntity{}
}

func (b *BaseEntity) SetId(id int64) {
	b.ID = &id
}

func (b *BaseEntity) isEntity() bool {
	return true
}
