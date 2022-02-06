// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"oauth-az/ent/oauth2client"
	"oauth-az/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OAuth2ClientUpdate is the builder for updating OAuth2Client entities.
type OAuth2ClientUpdate struct {
	config
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// Where appends a list predicates to the OAuth2ClientUpdate builder.
func (ou *OAuth2ClientUpdate) Where(ps ...predicate.OAuth2Client) *OAuth2ClientUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OAuth2ClientUpdate) SetUpdatedAt(t time.Time) *OAuth2ClientUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ou *OAuth2ClientUpdate) SetNillableUpdatedAt(t *time.Time) *OAuth2ClientUpdate {
	if t != nil {
		ou.SetUpdatedAt(*t)
	}
	return ou
}

// SetInternalID sets the "internal_id" field.
func (ou *OAuth2ClientUpdate) SetInternalID(u uuid.UUID) *OAuth2ClientUpdate {
	ou.mutation.SetInternalID(u)
	return ou
}

// SetSecret sets the "secret" field.
func (ou *OAuth2ClientUpdate) SetSecret(s string) *OAuth2ClientUpdate {
	ou.mutation.SetSecret(s)
	return ou
}

// SetRedirectUris sets the "redirect_uris" field.
func (ou *OAuth2ClientUpdate) SetRedirectUris(s []string) *OAuth2ClientUpdate {
	ou.mutation.SetRedirectUris(s)
	return ou
}

// ClearRedirectUris clears the value of the "redirect_uris" field.
func (ou *OAuth2ClientUpdate) ClearRedirectUris() *OAuth2ClientUpdate {
	ou.mutation.ClearRedirectUris()
	return ou
}

// SetTrustedPeers sets the "trusted_peers" field.
func (ou *OAuth2ClientUpdate) SetTrustedPeers(s []string) *OAuth2ClientUpdate {
	ou.mutation.SetTrustedPeers(s)
	return ou
}

// ClearTrustedPeers clears the value of the "trusted_peers" field.
func (ou *OAuth2ClientUpdate) ClearTrustedPeers() *OAuth2ClientUpdate {
	ou.mutation.ClearTrustedPeers()
	return ou
}

// SetName sets the "name" field.
func (ou *OAuth2ClientUpdate) SetName(s string) *OAuth2ClientUpdate {
	ou.mutation.SetName(s)
	return ou
}

// Mutation returns the OAuth2ClientMutation object of the builder.
func (ou *OAuth2ClientUpdate) Mutation() *OAuth2ClientMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OAuth2ClientUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ou.defaults()
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuth2ClientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OAuth2ClientUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OAuth2ClientUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OAuth2ClientUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OAuth2ClientUpdate) defaults() {
	if _, ok := ou.mutation.UpdateTime(); !ok {
		v := oauth2client.UpdateDefaultUpdateTime()
		ou.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OAuth2ClientUpdate) check() error {
	if v, ok := ou.mutation.Secret(); ok {
		if err := oauth2client.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf("ent: validator failed for field \"secret\": %w", err)}
		}
	}
	if v, ok := ou.mutation.Name(); ok {
		if err := oauth2client.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ou *OAuth2ClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth2client.Table,
			Columns: oauth2client.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth2client.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth2client.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth2client.FieldUpdateTime,
		})
	}
	if value, ok := ou.mutation.InternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oauth2client.FieldInternalID,
		})
	}
	if value, ok := ou.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth2client.FieldSecret,
		})
	}
	if value, ok := ou.mutation.RedirectUris(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth2client.FieldRedirectUris,
		})
	}
	if ou.mutation.RedirectUrisCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: oauth2client.FieldRedirectUris,
		})
	}
	if value, ok := ou.mutation.TrustedPeers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth2client.FieldTrustedPeers,
		})
	}
	if ou.mutation.TrustedPeersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: oauth2client.FieldTrustedPeers,
		})
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth2client.FieldName,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth2client.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OAuth2ClientUpdateOne is the builder for updating a single OAuth2Client entity.
type OAuth2ClientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OAuth2ClientUpdateOne) SetUpdatedAt(t time.Time) *OAuth2ClientUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouo *OAuth2ClientUpdateOne) SetNillableUpdatedAt(t *time.Time) *OAuth2ClientUpdateOne {
	if t != nil {
		ouo.SetUpdatedAt(*t)
	}
	return ouo
}

// SetInternalID sets the "internal_id" field.
func (ouo *OAuth2ClientUpdateOne) SetInternalID(u uuid.UUID) *OAuth2ClientUpdateOne {
	ouo.mutation.SetInternalID(u)
	return ouo
}

// SetSecret sets the "secret" field.
func (ouo *OAuth2ClientUpdateOne) SetSecret(s string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetSecret(s)
	return ouo
}

// SetRedirectUris sets the "redirect_uris" field.
func (ouo *OAuth2ClientUpdateOne) SetRedirectUris(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetRedirectUris(s)
	return ouo
}

// ClearRedirectUris clears the value of the "redirect_uris" field.
func (ouo *OAuth2ClientUpdateOne) ClearRedirectUris() *OAuth2ClientUpdateOne {
	ouo.mutation.ClearRedirectUris()
	return ouo
}

// SetTrustedPeers sets the "trusted_peers" field.
func (ouo *OAuth2ClientUpdateOne) SetTrustedPeers(s []string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetTrustedPeers(s)
	return ouo
}

// ClearTrustedPeers clears the value of the "trusted_peers" field.
func (ouo *OAuth2ClientUpdateOne) ClearTrustedPeers() *OAuth2ClientUpdateOne {
	ouo.mutation.ClearTrustedPeers()
	return ouo
}

// SetName sets the "name" field.
func (ouo *OAuth2ClientUpdateOne) SetName(s string) *OAuth2ClientUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// Mutation returns the OAuth2ClientMutation object of the builder.
func (ouo *OAuth2ClientUpdateOne) Mutation() *OAuth2ClientMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OAuth2ClientUpdateOne) Select(field string, fields ...string) *OAuth2ClientUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated OAuth2Client entity.
func (ouo *OAuth2ClientUpdateOne) Save(ctx context.Context) (*OAuth2Client, error) {
	var (
		err  error
		node *OAuth2Client
	)
	ouo.defaults()
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuth2ClientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OAuth2ClientUpdateOne) SaveX(ctx context.Context) *OAuth2Client {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OAuth2ClientUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OAuth2ClientUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OAuth2ClientUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdateTime(); !ok {
		v := oauth2client.UpdateDefaultUpdateTime()
		ouo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OAuth2ClientUpdateOne) check() error {
	if v, ok := ouo.mutation.Secret(); ok {
		if err := oauth2client.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf("ent: validator failed for field \"secret\": %w", err)}
		}
	}
	if v, ok := ouo.mutation.Name(); ok {
		if err := oauth2client.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (ouo *OAuth2ClientUpdateOne) sqlSave(ctx context.Context) (_node *OAuth2Client, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth2client.Table,
			Columns: oauth2client.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth2client.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OAuth2Client.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oauth2client.FieldID)
		for _, f := range fields {
			if !oauth2client.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oauth2client.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth2client.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: oauth2client.FieldUpdateTime,
		})
	}
	if value, ok := ouo.mutation.InternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: oauth2client.FieldInternalID,
		})
	}
	if value, ok := ouo.mutation.Secret(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth2client.FieldSecret,
		})
	}
	if value, ok := ouo.mutation.RedirectUris(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth2client.FieldRedirectUris,
		})
	}
	if ouo.mutation.RedirectUrisCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: oauth2client.FieldRedirectUris,
		})
	}
	if value, ok := ouo.mutation.TrustedPeers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: oauth2client.FieldTrustedPeers,
		})
	}
	if ouo.mutation.TrustedPeersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: oauth2client.FieldTrustedPeers,
		})
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: oauth2client.FieldName,
		})
	}
	_node = &OAuth2Client{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oauth2client.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
