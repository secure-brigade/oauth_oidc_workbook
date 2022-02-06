package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// AuthCode holds the schema definition for the AuthCode entity.
type AuthCode struct {
	ent.Schema
}

// Mixin returns the mixin schema for AuthCode.
func (AuthCode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		mixin.Time{},
	}
}

// Fields of the AuthCode.
func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("value", uuid.UUID{}).
			Default(uuid.New),
		field.Text("client_id").
			SchemaType(textSchema).
			NotEmpty(),
		field.JSON("scopes", []string{}).
			Optional(),
		field.JSON("response_type", []string{}).
			Optional(),
		field.Text("nonce").
			SchemaType(textSchema).
			NotEmpty(),
		field.Text("redirect_uri").
			SchemaType(textSchema).
			NotEmpty(),
		field.Time("expiry").
			SchemaType(timeSchema),
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

// Edges of the AuthCode.
func (AuthCode) Edges() []ent.Edge {
	return []ent.Edge{}
}
