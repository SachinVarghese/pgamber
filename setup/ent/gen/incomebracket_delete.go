// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/incomebracket"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/predicate"
)

// IncomeBracketDelete is the builder for deleting a IncomeBracket entity.
type IncomeBracketDelete struct {
	config
	hooks    []Hook
	mutation *IncomeBracketMutation
}

// Where appends a list predicates to the IncomeBracketDelete builder.
func (ibd *IncomeBracketDelete) Where(ps ...predicate.IncomeBracket) *IncomeBracketDelete {
	ibd.mutation.Where(ps...)
	return ibd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ibd *IncomeBracketDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ibd.hooks) == 0 {
		affected, err = ibd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncomeBracketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ibd.mutation = mutation
			affected, err = ibd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ibd.hooks) - 1; i >= 0; i-- {
			if ibd.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ibd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ibd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibd *IncomeBracketDelete) ExecX(ctx context.Context) int {
	n, err := ibd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ibd *IncomeBracketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: incomebracket.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: incomebracket.FieldID,
			},
		},
	}
	if ps := ibd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ibd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// IncomeBracketDeleteOne is the builder for deleting a single IncomeBracket entity.
type IncomeBracketDeleteOne struct {
	ibd *IncomeBracketDelete
}

// Exec executes the deletion query.
func (ibdo *IncomeBracketDeleteOne) Exec(ctx context.Context) error {
	n, err := ibdo.ibd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{incomebracket.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ibdo *IncomeBracketDeleteOne) ExecX(ctx context.Context) {
	ibdo.ibd.ExecX(ctx)
}
