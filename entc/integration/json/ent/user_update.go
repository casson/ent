// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/json/ent/predicate"
	"github.com/facebook/ent/entc/integration/json/ent/schema"
	"github.com/facebook/ent/entc/integration/json/ent/user"
	"github.com/facebook/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetT sets the t field.
func (uu *UserUpdate) SetT(s *schema.T) *UserUpdate {
	uu.mutation.SetT(s)
	return uu
}

// ClearT clears the value of t.
func (uu *UserUpdate) ClearT() *UserUpdate {
	uu.mutation.ClearT()
	return uu
}

// SetURL sets the url field.
func (uu *UserUpdate) SetURL(u *url.URL) *UserUpdate {
	uu.mutation.SetURL(u)
	return uu
}

// ClearURL clears the value of url.
func (uu *UserUpdate) ClearURL() *UserUpdate {
	uu.mutation.ClearURL()
	return uu
}

// SetRaw sets the raw field.
func (uu *UserUpdate) SetRaw(jm json.RawMessage) *UserUpdate {
	uu.mutation.SetRaw(jm)
	return uu
}

// ClearRaw clears the value of raw.
func (uu *UserUpdate) ClearRaw() *UserUpdate {
	uu.mutation.ClearRaw()
	return uu
}

// SetDirs sets the dirs field.
func (uu *UserUpdate) SetDirs(h []http.Dir) *UserUpdate {
	uu.mutation.SetDirs(h)
	return uu
}

// ClearDirs clears the value of dirs.
func (uu *UserUpdate) ClearDirs() *UserUpdate {
	uu.mutation.ClearDirs()
	return uu
}

// SetInts sets the ints field.
func (uu *UserUpdate) SetInts(i []int) *UserUpdate {
	uu.mutation.SetInts(i)
	return uu
}

// ClearInts clears the value of ints.
func (uu *UserUpdate) ClearInts() *UserUpdate {
	uu.mutation.ClearInts()
	return uu
}

// SetFloats sets the floats field.
func (uu *UserUpdate) SetFloats(f []float64) *UserUpdate {
	uu.mutation.SetFloats(f)
	return uu
}

// ClearFloats clears the value of floats.
func (uu *UserUpdate) ClearFloats() *UserUpdate {
	uu.mutation.ClearFloats()
	return uu
}

// SetStrings sets the strings field.
func (uu *UserUpdate) SetStrings(s []string) *UserUpdate {
	uu.mutation.SetStrings(s)
	return uu
}

// ClearStrings clears the value of strings.
func (uu *UserUpdate) ClearStrings() *UserUpdate {
	uu.mutation.ClearStrings()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.T(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldT,
		})
	}
	if uu.mutation.TCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldT,
		})
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldURL,
		})
	}
	if uu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldURL,
		})
	}
	if value, ok := uu.mutation.Raw(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRaw,
		})
	}
	if uu.mutation.RawCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldRaw,
		})
	}
	if value, ok := uu.mutation.Dirs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldDirs,
		})
	}
	if uu.mutation.DirsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldDirs,
		})
	}
	if value, ok := uu.mutation.Ints(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldInts,
		})
	}
	if uu.mutation.IntsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldInts,
		})
	}
	if value, ok := uu.mutation.Floats(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldFloats,
		})
	}
	if uu.mutation.FloatsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldFloats,
		})
	}
	if value, ok := uu.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldStrings,
		})
	}
	if uu.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldStrings,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetT sets the t field.
func (uuo *UserUpdateOne) SetT(s *schema.T) *UserUpdateOne {
	uuo.mutation.SetT(s)
	return uuo
}

// ClearT clears the value of t.
func (uuo *UserUpdateOne) ClearT() *UserUpdateOne {
	uuo.mutation.ClearT()
	return uuo
}

// SetURL sets the url field.
func (uuo *UserUpdateOne) SetURL(u *url.URL) *UserUpdateOne {
	uuo.mutation.SetURL(u)
	return uuo
}

// ClearURL clears the value of url.
func (uuo *UserUpdateOne) ClearURL() *UserUpdateOne {
	uuo.mutation.ClearURL()
	return uuo
}

// SetRaw sets the raw field.
func (uuo *UserUpdateOne) SetRaw(jm json.RawMessage) *UserUpdateOne {
	uuo.mutation.SetRaw(jm)
	return uuo
}

// ClearRaw clears the value of raw.
func (uuo *UserUpdateOne) ClearRaw() *UserUpdateOne {
	uuo.mutation.ClearRaw()
	return uuo
}

// SetDirs sets the dirs field.
func (uuo *UserUpdateOne) SetDirs(h []http.Dir) *UserUpdateOne {
	uuo.mutation.SetDirs(h)
	return uuo
}

// ClearDirs clears the value of dirs.
func (uuo *UserUpdateOne) ClearDirs() *UserUpdateOne {
	uuo.mutation.ClearDirs()
	return uuo
}

// SetInts sets the ints field.
func (uuo *UserUpdateOne) SetInts(i []int) *UserUpdateOne {
	uuo.mutation.SetInts(i)
	return uuo
}

// ClearInts clears the value of ints.
func (uuo *UserUpdateOne) ClearInts() *UserUpdateOne {
	uuo.mutation.ClearInts()
	return uuo
}

// SetFloats sets the floats field.
func (uuo *UserUpdateOne) SetFloats(f []float64) *UserUpdateOne {
	uuo.mutation.SetFloats(f)
	return uuo
}

// ClearFloats clears the value of floats.
func (uuo *UserUpdateOne) ClearFloats() *UserUpdateOne {
	uuo.mutation.ClearFloats()
	return uuo
}

// SetStrings sets the strings field.
func (uuo *UserUpdateOne) SetStrings(s []string) *UserUpdateOne {
	uuo.mutation.SetStrings(s)
	return uuo
}

// ClearStrings clears the value of strings.
func (uuo *UserUpdateOne) ClearStrings() *UserUpdateOne {
	uuo.mutation.ClearStrings()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.T(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldT,
		})
	}
	if uuo.mutation.TCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldT,
		})
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldURL,
		})
	}
	if uuo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldURL,
		})
	}
	if value, ok := uuo.mutation.Raw(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldRaw,
		})
	}
	if uuo.mutation.RawCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldRaw,
		})
	}
	if value, ok := uuo.mutation.Dirs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldDirs,
		})
	}
	if uuo.mutation.DirsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldDirs,
		})
	}
	if value, ok := uuo.mutation.Ints(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldInts,
		})
	}
	if uuo.mutation.IntsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldInts,
		})
	}
	if value, ok := uuo.mutation.Floats(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldFloats,
		})
	}
	if uuo.mutation.FloatsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldFloats,
		})
	}
	if value, ok := uuo.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldStrings,
		})
	}
	if uuo.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldStrings,
		})
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
