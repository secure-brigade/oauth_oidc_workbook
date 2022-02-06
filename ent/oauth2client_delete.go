// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"oauth-az/ent/oauth2client"
	"oauth-az/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OAuth2ClientDelete is the builder for deleting a OAuth2Client entity.
type OAuth2ClientDelete struct {
	config
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// Where appends a list predicates to the OAuth2ClientDelete builder.
func (od *OAuth2ClientDelete) Where(ps ...predicate.OAuth2Client) *OAuth2ClientDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OAuth2ClientDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(od.hooks) == 0 {
		affected, err = od.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OAuth2ClientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			od.mutation = mutation
			affected, err = od.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(od.hooks) - 1; i >= 0; i-- {
			if od.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = od.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, od.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OAuth2ClientDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OAuth2ClientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: oauth2client.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth2client.FieldID,
			},
		},
	}
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, od.driver, _spec)
}

// OAuth2ClientDeleteOne is the builder for deleting a single OAuth2Client entity.
type OAuth2ClientDeleteOne struct {
	od *OAuth2ClientDelete
}

// Exec executes the deletion query.
func (odo *OAuth2ClientDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oauth2client.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OAuth2ClientDeleteOne) ExecX(ctx context.Context) {
	odo.od.ExecX(ctx)
}