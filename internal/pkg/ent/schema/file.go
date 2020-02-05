package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type File struct {
	ent.Schema
}

func (File) Config() ent.Config {
	return ent.Config{Table:"files"}
}

func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").MaxLen(100).NotEmpty(),
		field.String("mime_type").MaxLen(255).NotEmpty(),
		field.Int64("file_size"),
		field.Bool("is_draft").Default(false),
		field.Bool("is_deleted").Default(false),
		field.Int64("user_id"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("buckets", Bucket.Type).Ref("files").Unique(), // many-to-one
	}
}