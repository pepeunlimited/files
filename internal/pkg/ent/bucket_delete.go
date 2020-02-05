// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/files/internal/pkg/ent/bucket"
	"github.com/pepeunlimited/files/internal/pkg/ent/predicate"
)

// BucketDelete is the builder for deleting a Bucket entity.
type BucketDelete struct {
	config
	predicates []predicate.Bucket
}

// Where adds a new predicate to the delete builder.
func (bd *BucketDelete) Where(ps ...predicate.Bucket) *BucketDelete {
	bd.predicates = append(bd.predicates, ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BucketDelete) Exec(ctx context.Context) (int, error) {
	return bd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BucketDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BucketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bucket.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bucket.FieldID,
			},
		},
	}
	if ps := bd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
}

// BucketDeleteOne is the builder for deleting a single Bucket entity.
type BucketDeleteOne struct {
	bd *BucketDelete
}

// Exec executes the deletion query.
func (bdo *BucketDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bucket.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BucketDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}