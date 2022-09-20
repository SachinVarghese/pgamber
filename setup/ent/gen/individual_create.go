// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/individual"
)

// IndividualCreate is the builder for creating a Individual entity.
type IndividualCreate struct {
	config
	mutation *IndividualMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (ic *IndividualCreate) SetAge(i int) *IndividualCreate {
	ic.mutation.SetAge(i)
	return ic
}

// Mutation returns the IndividualMutation object of the builder.
func (ic *IndividualCreate) Mutation() *IndividualMutation {
	return ic.mutation
}

// Save creates the Individual in the database.
func (ic *IndividualCreate) Save(ctx context.Context) (*Individual, error) {
	var (
		err  error
		node *Individual
	)
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndividualMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Individual)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IndividualMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IndividualCreate) SaveX(ctx context.Context) *Individual {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IndividualCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IndividualCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IndividualCreate) check() error {
	if _, ok := ic.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`gen: missing required field "Individual.age"`)}
	}
	if v, ok := ic.mutation.Age(); ok {
		if err := individual.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`gen: validator failed for field "Individual.age": %w`, err)}
		}
	}
	return nil
}

func (ic *IndividualCreate) sqlSave(ctx context.Context) (*Individual, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *IndividualCreate) createSpec() (*Individual, *sqlgraph.CreateSpec) {
	var (
		_node = &Individual{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: individual.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: individual.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: individual.FieldAge,
		})
		_node.Age = value
	}
	return _node, _spec
}

// IndividualCreateBulk is the builder for creating many Individual entities in bulk.
type IndividualCreateBulk struct {
	config
	builders []*IndividualCreate
}

// Save creates the Individual entities in the database.
func (icb *IndividualCreateBulk) Save(ctx context.Context) ([]*Individual, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Individual, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IndividualMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IndividualCreateBulk) SaveX(ctx context.Context) []*Individual {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IndividualCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IndividualCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
