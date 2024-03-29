// Code generated by entc, DO NOT EDIT.

package authcode

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the authcode type in the database.
	Label = "auth_code"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldResponseType holds the string denoting the response_type field in the database.
	FieldResponseType = "response_type"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// FieldRedirectURI holds the string denoting the redirect_uri field in the database.
	FieldRedirectURI = "redirect_uri"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldCodeChallenge holds the string denoting the code_challenge field in the database.
	FieldCodeChallenge = "code_challenge"
	// FieldCodeChallengeMethod holds the string denoting the code_challenge_method field in the database.
	FieldCodeChallengeMethod = "code_challenge_method"
	// Table holds the table name of the authcode in the database.
	Table = "auth_codes"
)

// Columns holds all SQL columns for authcode fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreateTime,
	FieldUpdateTime,
	FieldValue,
	FieldClientID,
	FieldScopes,
	FieldResponseType,
	FieldNonce,
	FieldRedirectURI,
	FieldExpiry,
	FieldCodeChallenge,
	FieldCodeChallengeMethod,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue func() uuid.UUID
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// NonceValidator is a validator for the "nonce" field. It is called by the builders before save.
	NonceValidator func(string) error
	// RedirectURIValidator is a validator for the "redirect_uri" field. It is called by the builders before save.
	RedirectURIValidator func(string) error
	// DefaultCodeChallenge holds the default value on creation for the "code_challenge" field.
	DefaultCodeChallenge string
)

// CodeChallengeMethod defines the type for the "code_challenge_method" enum field.
type CodeChallengeMethod string

// CodeChallengeMethodPlain is the default value of the CodeChallengeMethod enum.
const DefaultCodeChallengeMethod = CodeChallengeMethodPlain

// CodeChallengeMethod values.
const (
	CodeChallengeMethodPlain CodeChallengeMethod = "plain"
	CodeChallengeMethodS256  CodeChallengeMethod = "S256"
)

func (ccm CodeChallengeMethod) String() string {
	return string(ccm)
}

// CodeChallengeMethodValidator is a validator for the "code_challenge_method" field enum values. It is called by the builders before save.
func CodeChallengeMethodValidator(ccm CodeChallengeMethod) error {
	switch ccm {
	case CodeChallengeMethodPlain, CodeChallengeMethodS256:
		return nil
	default:
		return fmt.Errorf("authcode: invalid enum value for code_challenge_method field: %q", ccm)
	}
}
