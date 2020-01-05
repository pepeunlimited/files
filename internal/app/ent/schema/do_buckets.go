package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type DoBuckets struct {
	ent.Schema
}


func (DoBuckets) Config() ent.Config {
	return ent.Config{
		Table: "do_buckets",
	}
}

// Fields of the DoBuckets.
func (DoBuckets) Fields() []ent.Field {
	return []ent.Field{
		field.String("bucket_name").MinLen(3).MaxLen(63).NotEmpty().Unique(),
		field.String("endpoint").MaxLen(35).NotEmpty(),
		field.String("cdn_endpoint").Nillable().Optional(),
		field.Time("created_at"),
	}
}

// Edges of the DoBuckets.
func (DoBuckets) Edges() []ent.Edge {
	return nil
}