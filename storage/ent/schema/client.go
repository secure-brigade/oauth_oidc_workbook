// Inspired from https://github.com/dexidp/dex/blob/f92a6f4457c17b0ec29781a661bcffe7cf8f2f20/storage/ent/schema/client.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// OAuth2Client holds the schema definition for the OAuth2Client entity.
type OAuth2Client struct {
	ent.Schema
}

// Mixin returns the OAuth2Client mixin.
func (OAuth2Client) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		mixin.Time{},
	}
}

// Fields of the OAuth2Client.
func (OAuth2Client) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("internal_id", uuid.UUID{}).
			Default(uuid.New).
			Comment("UUIDの挙動を見たいので入れてるだけ"),
		field.Text("secret").
			MinLen(16).
			SchemaType(textSchema).
			NotEmpty(),
		field.JSON("redirect_uris", []string{}).
			Optional(),
		field.JSON("trusted_peers", []string{}).
			Optional(),
		field.Enum("type").
			Values("confidential", "public", "credentialed").
			Immutable().
			Default("confidential").
			Comment("credentialedはOAuth2.1で定義されたもの"),
		field.Text("name").
			SchemaType(textSchema).
			NotEmpty(),
	}
}

// Edges of the OAuth2Client.
func (OAuth2Client) Edges() []ent.Edge {
	return []ent.Edge{}
}
