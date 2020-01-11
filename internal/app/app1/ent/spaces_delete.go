// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/files/internal/app/app1/ent/predicate"
	"github.com/pepeunlimited/files/internal/app/app1/ent/spaces"
)

// SpacesDelete is the builder for deleting a Spaces entity.
type SpacesDelete struct {
	config
	predicates []predicate.Spaces
}

// Where adds a new predicate to the delete builder.
func (sd *SpacesDelete) Where(ps ...predicate.Spaces) *SpacesDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SpacesDelete) Exec(ctx context.Context) (int, error) {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SpacesDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SpacesDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: spaces.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spaces.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, spec)
}

// SpacesDeleteOne is the builder for deleting a single Spaces entity.
type SpacesDeleteOne struct {
	sd *SpacesDelete
}

// Exec executes the deletion query.
func (sdo *SpacesDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{spaces.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SpacesDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}