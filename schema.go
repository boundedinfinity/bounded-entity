package bounded_entity

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/google/uuid"
)

type SchemaId uuid.UUID

type SchemaVersion rfc3339date.Rfc3339Date

type BoundedEntity interface {
	EntityType() string
	EntityId() SchemaId
	EntityVersion() SchemaVersion
}
