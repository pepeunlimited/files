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

// FilesCreate is the builder for creating a Files entity.
type FilesCreate struct {
	config
	filename   *string
	mime_type  *string
	file_size  *int64
	is_draft   *bool
	is_deleted *bool
	user_id    *int64
	created_at *time.Time
	updated_at *time.Time
	spaces     map[int]struct{}
}

// SetFilename sets the filename field.
func (fc *FilesCreate) SetFilename(s string) *FilesCreate {
	fc.filename = &s
	return fc
}

// SetMimeType sets the mime_type field.
func (fc *FilesCreate) SetMimeType(s string) *FilesCreate {
	fc.mime_type = &s
	return fc
}

// SetFileSize sets the file_size field.
func (fc *FilesCreate) SetFileSize(i int64) *FilesCreate {
	fc.file_size = &i
	return fc
}

// SetIsDraft sets the is_draft field.
func (fc *FilesCreate) SetIsDraft(b bool) *FilesCreate {
	fc.is_draft = &b
	return fc
}

// SetNillableIsDraft sets the is_draft field if the given value is not nil.
func (fc *FilesCreate) SetNillableIsDraft(b *bool) *FilesCreate {
	if b != nil {
		fc.SetIsDraft(*b)
	}
	return fc
}

// SetIsDeleted sets the is_deleted field.
func (fc *FilesCreate) SetIsDeleted(b bool) *FilesCreate {
	fc.is_deleted = &b
	return fc
}

// SetNillableIsDeleted sets the is_deleted field if the given value is not nil.
func (fc *FilesCreate) SetNillableIsDeleted(b *bool) *FilesCreate {
	if b != nil {
		fc.SetIsDeleted(*b)
	}
	return fc
}

// SetUserID sets the user_id field.
func (fc *FilesCreate) SetUserID(i int64) *FilesCreate {
	fc.user_id = &i
	return fc
}

// SetCreatedAt sets the created_at field.
func (fc *FilesCreate) SetCreatedAt(t time.Time) *FilesCreate {
	fc.created_at = &t
	return fc
}

// SetUpdatedAt sets the updated_at field.
func (fc *FilesCreate) SetUpdatedAt(t time.Time) *FilesCreate {
	fc.updated_at = &t
	return fc
}

// SetSpacesID sets the spaces edge to Spaces by id.
func (fc *FilesCreate) SetSpacesID(id int) *FilesCreate {
	if fc.spaces == nil {
		fc.spaces = make(map[int]struct{})
	}
	fc.spaces[id] = struct{}{}
	return fc
}

// SetNillableSpacesID sets the spaces edge to Spaces by id if the given value is not nil.
func (fc *FilesCreate) SetNillableSpacesID(id *int) *FilesCreate {
	if id != nil {
		fc = fc.SetSpacesID(*id)
	}
	return fc
}

// SetSpaces sets the spaces edge to Spaces.
func (fc *FilesCreate) SetSpaces(s *Spaces) *FilesCreate {
	return fc.SetSpacesID(s.ID)
}

// Save creates the Files in the database.
func (fc *FilesCreate) Save(ctx context.Context) (*Files, error) {
	if fc.filename == nil {
		return nil, errors.New("ent: missing required field \"filename\"")
	}
	if err := files.FilenameValidator(*fc.filename); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"filename\": %v", err)
	}
	if fc.mime_type == nil {
		return nil, errors.New("ent: missing required field \"mime_type\"")
	}
	if err := files.MimeTypeValidator(*fc.mime_type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"mime_type\": %v", err)
	}
	if fc.file_size == nil {
		return nil, errors.New("ent: missing required field \"file_size\"")
	}
	if fc.is_draft == nil {
		v := files.DefaultIsDraft
		fc.is_draft = &v
	}
	if fc.is_deleted == nil {
		v := files.DefaultIsDeleted
		fc.is_deleted = &v
	}
	if fc.user_id == nil {
		return nil, errors.New("ent: missing required field \"user_id\"")
	}
	if fc.created_at == nil {
		return nil, errors.New("ent: missing required field \"created_at\"")
	}
	if fc.updated_at == nil {
		return nil, errors.New("ent: missing required field \"updated_at\"")
	}
	if len(fc.spaces) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spaces\"")
	}
	return fc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FilesCreate) SaveX(ctx context.Context) *Files {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FilesCreate) sqlSave(ctx context.Context) (*Files, error) {
	var (
		f    = &Files{config: fc.config}
		spec = &sqlgraph.CreateSpec{
			Table: files.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: files.FieldID,
			},
		}
	)
	if value := fc.filename; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldFilename,
		})
		f.Filename = *value
	}
	if value := fc.mime_type; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldMimeType,
		})
		f.MimeType = *value
	}
	if value := fc.file_size; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldFileSize,
		})
		f.FileSize = *value
	}
	if value := fc.is_draft; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDraft,
		})
		f.IsDraft = *value
	}
	if value := fc.is_deleted; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDeleted,
		})
		f.IsDeleted = *value
	}
	if value := fc.user_id; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldUserID,
		})
		f.UserID = *value
	}
	if value := fc.created_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldCreatedAt,
		})
		f.CreatedAt = *value
	}
	if value := fc.updated_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldUpdatedAt,
		})
		f.UpdatedAt = *value
	}
	if nodes := fc.spaces; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   files.SpacesTable,
			Columns: []string{files.SpacesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: spaces.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, fc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}