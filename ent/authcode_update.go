// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"oauth-az/ent/authcode"
	"oauth-az/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuthCodeUpdate is the builder for updating AuthCode entities.
type AuthCodeUpdate struct {
	config
	hooks    []Hook
	mutation *AuthCodeMutation
}

// Where appends a list predicates to the AuthCodeUpdate builder.
func (acu *AuthCodeUpdate) Where(ps ...predicate.AuthCode) *AuthCodeUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetUpdatedAt sets the "updated_at" field.
func (acu *AuthCodeUpdate) SetUpdatedAt(t time.Time) *AuthCodeUpdate {
	acu.mutation.SetUpdatedAt(t)
	return acu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableUpdatedAt(t *time.Time) *AuthCodeUpdate {
	if t != nil {
		acu.SetUpdatedAt(*t)
	}
	return acu
}

// SetValue sets the "value" field.
func (acu *AuthCodeUpdate) SetValue(u uuid.UUID) *AuthCodeUpdate {
	acu.mutation.SetValue(u)
	return acu
}

// SetClientID sets the "client_id" field.
func (acu *AuthCodeUpdate) SetClientID(s string) *AuthCodeUpdate {
	acu.mutation.SetClientID(s)
	return acu
}

// SetScopes sets the "scopes" field.
func (acu *AuthCodeUpdate) SetScopes(s []string) *AuthCodeUpdate {
	acu.mutation.SetScopes(s)
	return acu
}

// ClearScopes clears the value of the "scopes" field.
func (acu *AuthCodeUpdate) ClearScopes() *AuthCodeUpdate {
	acu.mutation.ClearScopes()
	return acu
}

// SetResponseType sets the "response_type" field.
func (acu *AuthCodeUpdate) SetResponseType(s []string) *AuthCodeUpdate {
	acu.mutation.SetResponseType(s)
	return acu
}

// ClearResponseType clears the value of the "response_type" field.
func (acu *AuthCodeUpdate) ClearResponseType() *AuthCodeUpdate {
	acu.mutation.ClearResponseType()
	return acu
}

// SetNonce sets the "nonce" field.
func (acu *AuthCodeUpdate) SetNonce(s string) *AuthCodeUpdate {
	acu.mutation.SetNonce(s)
	return acu
}

// SetRedirectURI sets the "redirect_uri" field.
func (acu *AuthCodeUpdate) SetRedirectURI(s string) *AuthCodeUpdate {
	acu.mutation.SetRedirectURI(s)
	return acu
}

// SetExpiry sets the "expiry" field.
func (acu *AuthCodeUpdate) SetExpiry(t time.Time) *AuthCodeUpdate {
	acu.mutation.SetExpiry(t)
	return acu
}

// SetCodeChallenge sets the "code_challenge" field.
func (acu *AuthCodeUpdate) SetCodeChallenge(s string) *AuthCodeUpdate {
	acu.mutation.SetCodeChallenge(s)
	return acu
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableCodeChallenge(s *string) *AuthCodeUpdate {
	if s != nil {
		acu.SetCodeChallenge(*s)
	}
	return acu
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (acu *AuthCodeUpdate) SetCodeChallengeMethod(acm authcode.CodeChallengeMethod) *AuthCodeUpdate {
	acu.mutation.SetCodeChallengeMethod(acm)
	return acu
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableCodeChallengeMethod(acm *authcode.CodeChallengeMethod) *AuthCodeUpdate {
	if acm != nil {
		acu.SetCodeChallengeMethod(*acm)
	}
	return acu
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acu *AuthCodeUpdate) Mutation() *AuthCodeMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AuthCodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	acu.defaults()
	if len(acu.hooks) == 0 {
		if err = acu.check(); err != nil {
			return 0, err
		}
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acu.check(); err != nil {
				return 0, err
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			if acu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AuthCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AuthCodeUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AuthCodeUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acu *AuthCodeUpdate) defaults() {
	if _, ok := acu.mutation.UpdateTime(); !ok {
		v := authcode.UpdateDefaultUpdateTime()
		acu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *AuthCodeUpdate) check() error {
	if v, ok := acu.mutation.ClientID(); ok {
		if err := authcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf("ent: validator failed for field \"client_id\": %w", err)}
		}
	}
	if v, ok := acu.mutation.Nonce(); ok {
		if err := authcode.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf("ent: validator failed for field \"nonce\": %w", err)}
		}
	}
	if v, ok := acu.mutation.RedirectURI(); ok {
		if err := authcode.RedirectURIValidator(v); err != nil {
			return &ValidationError{Name: "redirect_uri", err: fmt.Errorf("ent: validator failed for field \"redirect_uri\": %w", err)}
		}
	}
	if v, ok := acu.mutation.CodeChallengeMethod(); ok {
		if err := authcode.CodeChallengeMethodValidator(v); err != nil {
			return &ValidationError{Name: "code_challenge_method", err: fmt.Errorf("ent: validator failed for field \"code_challenge_method\": %w", err)}
		}
	}
	return nil
}

func (acu *AuthCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authcode.Table,
			Columns: authcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authcode.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldUpdatedAt,
		})
	}
	if value, ok := acu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldUpdateTime,
		})
	}
	if value, ok := acu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: authcode.FieldValue,
		})
	}
	if value, ok := acu.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClientID,
		})
	}
	if value, ok := acu.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldScopes,
		})
	}
	if acu.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldScopes,
		})
	}
	if value, ok := acu.mutation.ResponseType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldResponseType,
		})
	}
	if acu.mutation.ResponseTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldResponseType,
		})
	}
	if value, ok := acu.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldNonce,
		})
	}
	if value, ok := acu.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldRedirectURI,
		})
	}
	if value, ok := acu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldExpiry,
		})
	}
	if value, ok := acu.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallenge,
		})
	}
	if value, ok := acu.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authcode.FieldCodeChallengeMethod,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AuthCodeUpdateOne is the builder for updating a single AuthCode entity.
type AuthCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthCodeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (acuo *AuthCodeUpdateOne) SetUpdatedAt(t time.Time) *AuthCodeUpdateOne {
	acuo.mutation.SetUpdatedAt(t)
	return acuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableUpdatedAt(t *time.Time) *AuthCodeUpdateOne {
	if t != nil {
		acuo.SetUpdatedAt(*t)
	}
	return acuo
}

// SetValue sets the "value" field.
func (acuo *AuthCodeUpdateOne) SetValue(u uuid.UUID) *AuthCodeUpdateOne {
	acuo.mutation.SetValue(u)
	return acuo
}

// SetClientID sets the "client_id" field.
func (acuo *AuthCodeUpdateOne) SetClientID(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClientID(s)
	return acuo
}

// SetScopes sets the "scopes" field.
func (acuo *AuthCodeUpdateOne) SetScopes(s []string) *AuthCodeUpdateOne {
	acuo.mutation.SetScopes(s)
	return acuo
}

// ClearScopes clears the value of the "scopes" field.
func (acuo *AuthCodeUpdateOne) ClearScopes() *AuthCodeUpdateOne {
	acuo.mutation.ClearScopes()
	return acuo
}

// SetResponseType sets the "response_type" field.
func (acuo *AuthCodeUpdateOne) SetResponseType(s []string) *AuthCodeUpdateOne {
	acuo.mutation.SetResponseType(s)
	return acuo
}

// ClearResponseType clears the value of the "response_type" field.
func (acuo *AuthCodeUpdateOne) ClearResponseType() *AuthCodeUpdateOne {
	acuo.mutation.ClearResponseType()
	return acuo
}

// SetNonce sets the "nonce" field.
func (acuo *AuthCodeUpdateOne) SetNonce(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetNonce(s)
	return acuo
}

// SetRedirectURI sets the "redirect_uri" field.
func (acuo *AuthCodeUpdateOne) SetRedirectURI(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetRedirectURI(s)
	return acuo
}

// SetExpiry sets the "expiry" field.
func (acuo *AuthCodeUpdateOne) SetExpiry(t time.Time) *AuthCodeUpdateOne {
	acuo.mutation.SetExpiry(t)
	return acuo
}

// SetCodeChallenge sets the "code_challenge" field.
func (acuo *AuthCodeUpdateOne) SetCodeChallenge(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetCodeChallenge(s)
	return acuo
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableCodeChallenge(s *string) *AuthCodeUpdateOne {
	if s != nil {
		acuo.SetCodeChallenge(*s)
	}
	return acuo
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (acuo *AuthCodeUpdateOne) SetCodeChallengeMethod(acm authcode.CodeChallengeMethod) *AuthCodeUpdateOne {
	acuo.mutation.SetCodeChallengeMethod(acm)
	return acuo
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableCodeChallengeMethod(acm *authcode.CodeChallengeMethod) *AuthCodeUpdateOne {
	if acm != nil {
		acuo.SetCodeChallengeMethod(*acm)
	}
	return acuo
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acuo *AuthCodeUpdateOne) Mutation() *AuthCodeMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AuthCodeUpdateOne) Select(field string, fields ...string) *AuthCodeUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AuthCode entity.
func (acuo *AuthCodeUpdateOne) Save(ctx context.Context) (*AuthCode, error) {
	var (
		err  error
		node *AuthCode
	)
	acuo.defaults()
	if len(acuo.hooks) == 0 {
		if err = acuo.check(); err != nil {
			return nil, err
		}
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acuo.check(); err != nil {
				return nil, err
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			if acuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) SaveX(ctx context.Context) *AuthCode {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AuthCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acuo *AuthCodeUpdateOne) defaults() {
	if _, ok := acuo.mutation.UpdateTime(); !ok {
		v := authcode.UpdateDefaultUpdateTime()
		acuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AuthCodeUpdateOne) check() error {
	if v, ok := acuo.mutation.ClientID(); ok {
		if err := authcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf("ent: validator failed for field \"client_id\": %w", err)}
		}
	}
	if v, ok := acuo.mutation.Nonce(); ok {
		if err := authcode.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf("ent: validator failed for field \"nonce\": %w", err)}
		}
	}
	if v, ok := acuo.mutation.RedirectURI(); ok {
		if err := authcode.RedirectURIValidator(v); err != nil {
			return &ValidationError{Name: "redirect_uri", err: fmt.Errorf("ent: validator failed for field \"redirect_uri\": %w", err)}
		}
	}
	if v, ok := acuo.mutation.CodeChallengeMethod(); ok {
		if err := authcode.CodeChallengeMethodValidator(v); err != nil {
			return &ValidationError{Name: "code_challenge_method", err: fmt.Errorf("ent: validator failed for field \"code_challenge_method\": %w", err)}
		}
	}
	return nil
}

func (acuo *AuthCodeUpdateOne) sqlSave(ctx context.Context) (_node *AuthCode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authcode.Table,
			Columns: authcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authcode.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AuthCode.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authcode.FieldID)
		for _, f := range fields {
			if !authcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != authcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldUpdatedAt,
		})
	}
	if value, ok := acuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldUpdateTime,
		})
	}
	if value, ok := acuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: authcode.FieldValue,
		})
	}
	if value, ok := acuo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClientID,
		})
	}
	if value, ok := acuo.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldScopes,
		})
	}
	if acuo.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldScopes,
		})
	}
	if value, ok := acuo.mutation.ResponseType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldResponseType,
		})
	}
	if acuo.mutation.ResponseTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldResponseType,
		})
	}
	if value, ok := acuo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldNonce,
		})
	}
	if value, ok := acuo.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldRedirectURI,
		})
	}
	if value, ok := acuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldExpiry,
		})
	}
	if value, ok := acuo.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallenge,
		})
	}
	if value, ok := acuo.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: authcode.FieldCodeChallengeMethod,
		})
	}
	_node = &AuthCode{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
