package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Blocker holds the schema definition for the Blocker entity.
type Blocker struct {
	ent.Schema
}

// Fields of the Blocker.
func (Blocker) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.String("name"),
		field.String("api_key"), // hash of api_key
		field.Bool("revoked"),
		field.String("ip_address"),
		field.String("type"), // rupture ou stream
		field.Time("expiration").Optional(),
	}
}

// Edges of the Blocker.
func (Blocker) Edges() []ent.Edge {
	return nil
}
