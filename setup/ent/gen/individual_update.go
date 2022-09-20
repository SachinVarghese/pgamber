// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/individual"
	"github.com/SachinVarghese/pgamber/setup/ent/gen/predicate"
)

// IndividualUpdate is the builder for updating Individual entities.
type IndividualUpdate struct {
	config
	hooks    []Hook
	mutation *IndividualMutation
}

// Where appends a list predicates to the IndividualUpdate builder.
func (iu *IndividualUpdate) Where(ps ...predicate.Individual) *IndividualUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetAge sets the "age" field.
func (iu *IndividualUpdate) SetAge(i int) *IndividualUpdate {
	iu.mutation.ResetAge()
	iu.mutation.SetAge(i)
	return iu
}

// AddAge adds i to the "age" field.
func (iu *IndividualUpdate) AddAge(i int) *IndividualUpdate {
	iu.mutation.AddAge(i)
	return iu
}

// Mutation returns the IndividualMutation object of the builder.
func (iu *IndividualUpdate) Mutation() *IndividualMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IndividualUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndividualMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IndividualUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IndividualUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IndividualUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IndividualUpdate) check() error {
	if v, ok := iu.mutation.Age(); ok {
		if err := individual.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`gen: validator failed for field "Individual.age": %w`, err)}
		}
	}
	return nil
}

func (iu *IndividualUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   individual.Table,
			Columns: individual.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: individual.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: individual.FieldAge,
		})
	}
	if value, ok := iu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: individual.FieldAge,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{individual.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// IndividualUpdateOne is the builder for updating a single Individual entity.
type IndividualUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IndividualMutation
}

// SetAge sets the "age" field.
func (iuo *IndividualUpdateOne) SetAge(i int) *IndividualUpdateOne {
	iuo.mutation.ResetAge()
	iuo.mutation.SetAge(i)
	return iuo
}

// AddAge adds i to the "age" field.
func (iuo *IndividualUpdateOne) AddAge(i int) *IndividualUpdateOne {
	iuo.mutation.AddAge(i)
	return iuo
}

// Mutation returns the IndividualMutation object of the builder.
func (iuo *IndividualUpdateOne) Mutation() *IndividualMutation {
	return iuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IndividualUpdateOne) Select(field string, fields ...string) *IndividualUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Individual entity.
func (iuo *IndividualUpdateOne) Save(ctx context.Context) (*Individual, error) {
	var (
		err  error
		node *Individual
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndividualMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (iuo *IndividualUpdateOne) SaveX(ctx context.Context) *Individual {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IndividualUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IndividualUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IndividualUpdateOne) check() error {
	if v, ok := iuo.mutation.Age(); ok {
		if err := individual.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`gen: validator failed for field "Individual.age": %w`, err)}
		}
	}
	return nil
}

func (iuo *IndividualUpdateOne) sqlSave(ctx context.Context) (_node *Individual, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   individual.Table,
			Columns: individual.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: individual.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Individual.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, individual.FieldID)
		for _, f := range fields {
			if !individual.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != individual.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: individual.FieldAge,
		})
	}
	if value, ok := iuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: individual.FieldAge,
		})
	}
	_node = &Individual{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{individual.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
