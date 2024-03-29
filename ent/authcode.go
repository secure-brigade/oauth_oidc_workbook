// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"oauth-az/ent/authcode"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AuthCode is the model entity for the AuthCode schema.
type AuthCode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	// 作成日時
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	// 更新日時
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Value holds the value of the "value" field.
	Value uuid.UUID `json:"value,omitempty"`
	// ClientID holds the value of the "client_id" field.
	ClientID string `json:"client_id,omitempty"`
	// Scopes holds the value of the "scopes" field.
	Scopes []string `json:"scopes,omitempty"`
	// ResponseType holds the value of the "response_type" field.
	ResponseType []string `json:"response_type,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
	// RedirectURI holds the value of the "redirect_uri" field.
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// CodeChallenge holds the value of the "code_challenge" field.
	// OAuth2.1ではrequired
	CodeChallenge string `json:"code_challenge,omitempty"`
	// CodeChallengeMethod holds the value of the "code_challenge_method" field.
	// OAuth2.1ではrequired
	CodeChallengeMethod authcode.CodeChallengeMethod `json:"code_challenge_method,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthCode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authcode.FieldScopes, authcode.FieldResponseType:
			values[i] = new([]byte)
		case authcode.FieldID:
			values[i] = new(sql.NullInt64)
		case authcode.FieldClientID, authcode.FieldNonce, authcode.FieldRedirectURI, authcode.FieldCodeChallenge, authcode.FieldCodeChallengeMethod:
			values[i] = new(sql.NullString)
		case authcode.FieldCreatedAt, authcode.FieldUpdatedAt, authcode.FieldCreateTime, authcode.FieldUpdateTime, authcode.FieldExpiry:
			values[i] = new(sql.NullTime)
		case authcode.FieldValue:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthCode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthCode fields.
func (ac *AuthCode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authcode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case authcode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Time
			}
		case authcode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Time
			}
		case authcode.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ac.CreateTime = value.Time
			}
		case authcode.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ac.UpdateTime = value.Time
			}
		case authcode.FieldValue:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value != nil {
				ac.Value = *value
			}
		case authcode.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				ac.ClientID = value.String
			}
		case authcode.FieldScopes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ac.Scopes); err != nil {
					return fmt.Errorf("unmarshal field scopes: %w", err)
				}
			}
		case authcode.FieldResponseType:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field response_type", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ac.ResponseType); err != nil {
					return fmt.Errorf("unmarshal field response_type: %w", err)
				}
			}
		case authcode.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				ac.Nonce = value.String
			}
		case authcode.FieldRedirectURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_uri", values[i])
			} else if value.Valid {
				ac.RedirectURI = value.String
			}
		case authcode.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				ac.Expiry = value.Time
			}
		case authcode.FieldCodeChallenge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge", values[i])
			} else if value.Valid {
				ac.CodeChallenge = value.String
			}
		case authcode.FieldCodeChallengeMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge_method", values[i])
			} else if value.Valid {
				ac.CodeChallengeMethod = authcode.CodeChallengeMethod(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AuthCode.
// Note that you need to call AuthCode.Unwrap() before calling this method if this AuthCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AuthCode) Update() *AuthCodeUpdateOne {
	return (&AuthCodeClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AuthCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AuthCode) Unwrap() *AuthCode {
	tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthCode is not a transactional entity")
	}
	ac.config.driver = tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AuthCode) String() string {
	var builder strings.Builder
	builder.WriteString("AuthCode(")
	builder.WriteString(fmt.Sprintf("id=%v", ac.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(ac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(ac.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", create_time=")
	builder.WriteString(ac.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ac.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", ac.Value))
	builder.WriteString(", client_id=")
	builder.WriteString(ac.ClientID)
	builder.WriteString(", scopes=")
	builder.WriteString(fmt.Sprintf("%v", ac.Scopes))
	builder.WriteString(", response_type=")
	builder.WriteString(fmt.Sprintf("%v", ac.ResponseType))
	builder.WriteString(", nonce=")
	builder.WriteString(ac.Nonce)
	builder.WriteString(", redirect_uri=")
	builder.WriteString(ac.RedirectURI)
	builder.WriteString(", expiry=")
	builder.WriteString(ac.Expiry.Format(time.ANSIC))
	builder.WriteString(", code_challenge=")
	builder.WriteString(ac.CodeChallenge)
	builder.WriteString(", code_challenge_method=")
	builder.WriteString(fmt.Sprintf("%v", ac.CodeChallengeMethod))
	builder.WriteByte(')')
	return builder.String()
}

// AuthCodes is a parsable slice of AuthCode.
type AuthCodes []*AuthCode

func (ac AuthCodes) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
