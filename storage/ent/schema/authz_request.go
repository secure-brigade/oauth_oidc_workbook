package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// AuthRequest holds the schema definition for the AuthRequest entity.
type AuthRequest struct {
	ent.Schema
}

// Mixin returns the mixin schema for AuthRequest.
func (AuthRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		mixin.Time{},
	}
}

// Fields of the AuthRequest.
func (AuthRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("request_id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUIDの挙動を見たいので入れてるだけ"),
		field.Text("client_id").
			SchemaType(textSchema),
		field.JSON("scopes", []string{}).
			Optional(),
		field.JSON("response_type", []string{}).
			Optional(),
		field.Text("redirect_uri").
			SchemaType(textSchema),
		field.String("nonce").
			SchemaType(textSchema),
		field.Text("code_challenge").
			SchemaType(textSchema).
			Default("").
			Comment("OAuth2.1ではrequired"),
		field.Enum("code_challenge_method").
			Values("plain", "S256").
			Default("plain").
			Comment("OAuth2.1ではrequired"),
	}
}

// Edges of the AuthRequest.
func (AuthRequest) Edges() []ent.Edge {
	return []ent.Edge{}
}
