// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pepeunlimited/files/internal/app/app1/ent/files"
	"github.com/pepeunlimited/files/internal/app/app1/ent/spaces"
)

// SpacesCreate is the builder for creating a Spaces entity.
type SpacesCreate struct {
	config
	name         *string
	endpoint     *string
	cdn_endpoint *string
	created_at   *time.Time
	files        map[int]struct{}
}

// SetName sets the name field.
func (sc *SpacesCreate) SetName(s string) *SpacesCreate {
	sc.name = &s
	return sc
}

// SetEndpoint sets the endpoint field.
func (sc *SpacesCreate) SetEndpoint(s string) *SpacesCreate {
	sc.endpoint = &s
	return sc
}

// SetCdnEndpoint sets the cdn_endpoint field.
func (sc *SpacesCreate) SetCdnEndpoint(s string) *SpacesCreate {
	sc.cdn_endpoint = &s
	return sc
}

// SetNillableCdnEndpoint sets the cdn_endpoint field if the given value is not nil.
func (sc *SpacesCreate) SetNillableCdnEndpoint(s *string) *SpacesCreate {
	if s != nil {
		sc.SetCdnEndpoint(*s)
	}
	return sc
}

// SetCreatedAt sets the created_at field.
func (sc *SpacesCreate) SetCreatedAt(t time.Time) *SpacesCreate {
	sc.created_at = &t
	return sc
}

// AddFileIDs adds the files edge to Files by ids.
func (sc *SpacesCreate) AddFileIDs(ids ...int) *SpacesCreate {
	if sc.files == nil {
		sc.files = make(map[int]struct{})
	}
	for i := range ids {
		sc.files[ids[i]] = struct{}{}
	}
	return sc
}

// AddFiles adds the files edges to Files.
func (sc *SpacesCreate) AddFiles(f ...*Files) *SpacesCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return sc.AddFileIDs(ids...)
}

// Save creates the Spaces in the database.
func (sc *SpacesCreate) Save(ctx context.Context) (*Spaces, error) {
	if sc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := spaces.NameValidator(*sc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if sc.endpoint == nil {
		return nil, errors.New("ent: missing required field \"endpoint\"")
	}
	if err := spaces.EndpointValidator(*sc.endpoint); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"endpoint\": %v", err)
	}
	if sc.created_at == nil {
		return nil, errors.New("ent: missing required field \"created_at\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpacesCreate) SaveX(ctx context.Context) *Spaces {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SpacesCreate) sqlSave(ctx context.Context) (*Spaces, error) {
	var (
		s    = &Spaces{config: sc.config}
		spec = &sqlgraph.CreateSpec{
			Table: spaces.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spaces.FieldID,
			},
		}
	)
	if value := sc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: spaces.FieldName,
		})
		s.Name = *value
	}
	if value := sc.endpoint; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: spaces.FieldEndpoint,
		})
		s.Endpoint = *value
	}
	if value := sc.cdn_endpoint; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: spaces.FieldCdnEndpoint,
		})
		s.CdnEndpoint = value
	}
	if value := sc.created_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: spaces.FieldCreatedAt,
		})
		s.CreatedAt = *value
	}
	if nodes := sc.files; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spaces.FilesTable,
			Columns: []string{spaces.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: files.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
