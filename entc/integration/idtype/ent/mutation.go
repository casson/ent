// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/ent"
	"github.com/facebookincubator/ent/entc/integration/idtype/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op               Op
	typ              string
	id               *uint64
	name             *string
	clearedFields    map[string]bool
	spouse           *uint64
	clearedspouse    bool
	followers        map[uint64]struct{}
	removedfollowers map[uint64]struct{}
	following        map[uint64]struct{}
	removedfollowing map[uint64]struct{}
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op) *UserMutation {
	return &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]bool),
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetSpouseID sets the spouse edge to User by id.
func (m *UserMutation) SetSpouseID(id uint64) {
	m.spouse = &id
}

// ClearSpouse clears the spouse edge to User.
func (m *UserMutation) ClearSpouse() {
	m.clearedspouse = true
}

// SpouseCleared returns if the edge spouse was cleared.
func (m *UserMutation) SpouseCleared() bool {
	return m.clearedspouse
}

// SpouseID returns the spouse id in the mutation.
func (m *UserMutation) SpouseID() (id uint64, exists bool) {
	if m.spouse != nil {
		return *m.spouse, true
	}
	return
}

// SpouseIDs returns the spouse ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// SpouseID instead. It exists only for internal usage by the builders.
func (m *UserMutation) SpouseIDs() (ids []uint64) {
	if id := m.spouse; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetSpouse reset all changes of the spouse edge.
func (m *UserMutation) ResetSpouse() {
	m.spouse = nil
	m.clearedspouse = false
}

// AddFollowerIDs adds the followers edge to User by ids.
func (m *UserMutation) AddFollowerIDs(ids ...uint64) {
	if m.followers == nil {
		m.followers = make(map[uint64]struct{})
	}
	for i := range ids {
		m.followers[ids[i]] = struct{}{}
	}
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (m *UserMutation) RemoveFollowerIDs(ids ...uint64) {
	if m.removedfollowers == nil {
		m.removedfollowers = make(map[uint64]struct{})
	}
	for i := range ids {
		m.removedfollowers[ids[i]] = struct{}{}
	}
}

// RemovedFollowers returns the removed ids of followers.
func (m *UserMutation) RemovedFollowersIDs() (ids []uint64) {
	for id := range m.removedfollowers {
		ids = append(ids, id)
	}
	return
}

// FollowersIDs returns the followers ids in the mutation.
func (m *UserMutation) FollowersIDs() (ids []uint64) {
	for id := range m.followers {
		ids = append(ids, id)
	}
	return
}

// ResetFollowers reset all changes of the followers edge.
func (m *UserMutation) ResetFollowers() {
	m.followers = nil
	m.removedfollowers = nil
}

// AddFollowingIDs adds the following edge to User by ids.
func (m *UserMutation) AddFollowingIDs(ids ...uint64) {
	if m.following == nil {
		m.following = make(map[uint64]struct{})
	}
	for i := range ids {
		m.following[ids[i]] = struct{}{}
	}
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (m *UserMutation) RemoveFollowingIDs(ids ...uint64) {
	if m.removedfollowing == nil {
		m.removedfollowing = make(map[uint64]struct{})
	}
	for i := range ids {
		m.removedfollowing[ids[i]] = struct{}{}
	}
}

// RemovedFollowing returns the removed ids of following.
func (m *UserMutation) RemovedFollowingIDs() (ids []uint64) {
	for id := range m.removedfollowing {
		ids = append(ids, id)
	}
	return
}

// FollowingIDs returns the following ids in the mutation.
func (m *UserMutation) FollowingIDs() (ids []uint64) {
	for id := range m.following {
		ids = append(ids, id)
	}
	return
}

// ResetFollowing reset all changes of the following edge.
func (m *UserMutation) ResetFollowing() {
	m.following = nil
	m.removedfollowing = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.spouse != nil {
		edges = append(edges, user.EdgeSpouse)
	}
	if m.followers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.following != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeSpouse:
		if id := m.spouse; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.followers))
		for id := range m.followers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.following))
		for id := range m.following {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedfollowers != nil {
		edges = append(edges, user.EdgeFollowers)
	}
	if m.removedfollowing != nil {
		edges = append(edges, user.EdgeFollowing)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFollowers:
		ids := make([]ent.Value, 0, len(m.removedfollowers))
		for id := range m.removedfollowers {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeFollowing:
		ids := make([]ent.Value, 0, len(m.removedfollowing))
		for id := range m.removedfollowing {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedspouse {
		edges = append(edges, user.EdgeSpouse)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeSpouse:
		return m.clearedspouse
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeSpouse:
		m.ClearSpouse()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeSpouse:
		m.ResetSpouse()
		return nil
	case user.EdgeFollowers:
		m.ResetFollowers()
		return nil
	case user.EdgeFollowing:
		m.ResetFollowing()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
