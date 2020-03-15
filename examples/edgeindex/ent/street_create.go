// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
	"github.com/facebookincubator/ent/schema/field"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	mutation *StreetMutation
	hooks    []Hook
}

// SetName sets the name field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCityID sets the city edge to City by id.
func (sc *StreetCreate) SetCityID(id int) *StreetCreate {
	sc.mutation.SetCityID(id)
	return sc
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (sc *StreetCreate) SetNillableCityID(id *int) *StreetCreate {
	if id != nil {
		sc = sc.SetCityID(*id)
	}
	return sc
}

// SetCity sets the city edge to City.
func (sc *StreetCreate) SetCity(c *City) *StreetCreate {
	return sc.SetCityID(c.ID)
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	if _, ok := sc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	var (
		err  error
		node *Street
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			return node, err
		})
		for i := len(sc.hooks); i > 0; i-- {
			mut = sc.hooks[i-1](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	var (
		s     = &Street{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: street.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: street.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
		s.Name = value
	}
	if nodes := sc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
