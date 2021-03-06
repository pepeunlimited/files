package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Bucket struct {
	ent.Schema
}


func (Bucket) Config() ent.Config {
	return ent.Config{
		Table: "buckets",
	}
}

// Fields of the Spaces.
func (Bucket) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(3).MaxLen(63).NotEmpty().Unique(),
		field.String("endpoint").MaxLen(35).NotEmpty(),
		field.String("cdn_endpoint").Nillable().Optional(),
		field.Time("created_at"),
	}
}

// Edges of the Spaces.
func (Bucket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type), // one-to-many
	}
}