package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin holds the schema definition for the TimeMixin entity.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			SchemaType(timeSchema).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Immutable().
			Default(time.Now).
			Comment("作成日時"),
		field.Time("updated_at").
			SchemaType(timeSchema).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			Default(time.Now).
			Comment("更新日時"),
	}
}

// Edges of the TimeMixin.
func (TimeMixin) Edges() []ent.Edge {
	return nil
}
