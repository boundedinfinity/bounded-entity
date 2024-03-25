package bounded_entity

type BoundedEntityString struct {
	BoundedEntityBase
	Examples []string `json:"examples,omitempty"`
}

func (BoundedEntityString) EntityType() string {
	return "string"
}

var _ BoundedEntity = &BoundedEntityString{}
