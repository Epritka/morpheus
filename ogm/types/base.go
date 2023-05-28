package types

type Entity interface {
	isEntity() bool
}

type BaseEntity struct {
	ID     *int64   `ogm:"id"`
	Labels []string `ogm:"labels"`
}

func new() *BaseEntity {
	return &BaseEntity{}
}

func (b *BaseEntity) SetId(id int64) {
	b.ID = &id
}

func (b *BaseEntity) SetLabels(labels []string) {
	b.Labels = labels
}

func (b *BaseEntity) isModel() bool {
	return true
}
