// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/gremlin/ent"
)

// The CardFunc type is an adapter to allow the use of ordinary
// function as Card mutator.
type CardFunc func(context.Context, *ent.CardMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CardFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CardMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CardMutation", m)
	}
	return f(ctx, mv)
}

// The CommentFunc type is an adapter to allow the use of ordinary
// function as Comment mutator.
type CommentFunc func(context.Context, *ent.CommentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CommentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommentMutation", m)
	}
	return f(ctx, mv)
}

// The FieldTypeFunc type is an adapter to allow the use of ordinary
// function as FieldType mutator.
type FieldTypeFunc func(context.Context, *ent.FieldTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FieldTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FieldTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FieldTypeMutation", m)
	}
	return f(ctx, mv)
}

// The FileFunc type is an adapter to allow the use of ordinary
// function as File mutator.
type FileFunc func(context.Context, *ent.FileMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileMutation", m)
	}
	return f(ctx, mv)
}

// The FileTypeFunc type is an adapter to allow the use of ordinary
// function as FileType mutator.
type FileTypeFunc func(context.Context, *ent.FileTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FileTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FileTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FileTypeMutation", m)
	}
	return f(ctx, mv)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
	}
	return f(ctx, mv)
}

// The GroupInfoFunc type is an adapter to allow the use of ordinary
// function as GroupInfo mutator.
type GroupInfoFunc func(context.Context, *ent.GroupInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupInfoMutation", m)
	}
	return f(ctx, mv)
}

// The ItemFunc type is an adapter to allow the use of ordinary
// function as Item mutator.
type ItemFunc func(context.Context, *ent.ItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ItemMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemMutation", m)
	}
	return f(ctx, mv)
}

// The NodeFunc type is an adapter to allow the use of ordinary
// function as Node mutator.
type NodeFunc func(context.Context, *ent.NodeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NodeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NodeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NodeMutation", m)
	}
	return f(ctx, mv)
}

// The PetFunc type is an adapter to allow the use of ordinary
// function as Pet mutator.
type PetFunc func(context.Context, *ent.PetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PetMutation", m)
	}
	return f(ctx, mv)
}

// The SpecFunc type is an adapter to allow the use of ordinary
// function as Spec mutator.
type SpecFunc func(context.Context, *ent.SpecMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SpecFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SpecMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SpecMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// On executes the given hook only of the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return nil, fmt.Errorf("%s operation is not allowed", m.Op())
			}
			return next.Mutate(ctx, m)
		})
	}
}
