package bounded_entity

type BoundedEntityBase struct {
	SchemaId        SchemaId        `json:"schema-id,omitempty"`
	SchemaVersion   SchemaVersion   `json:"schema-version,omitempty"`
	Name            string          `json:"name,omitempty"`
	Description     string          `json:"description,omitempty"`
	ReadOnly        bool            `json:"read-only,omitempty"`
	WriteOnly       bool            `json:"write-only,omitempty"`
	DeprecatedSince bool            `json:"deprecated-since,omitempty"`
	RemovedIn       bool            `json:"removed-in,omitempty"`
	ReplacedBy      []BoundedEntity `json:"replaced-by,omitempty"`
}

func (b BoundedEntityBase) EntityVersion() SchemaVersion {
	return b.SchemaVersion
}

func (b BoundedEntityBase) EntityId() SchemaId {
	return b.SchemaId
}
