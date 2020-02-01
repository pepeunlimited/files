// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/pepeunlimited/files/internal/pkg/ent/buckets"
	"github.com/pepeunlimited/files/internal/pkg/ent/files"
	"github.com/pepeunlimited/files/internal/pkg/ent/predicate"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FilesUpdate is the builder for updating Files entities.
type FilesUpdate struct {
	config
	filename       *string
	mime_type      *string
	file_size      *int64
	addfile_size   *int64
	is_draft       *bool
	is_deleted     *bool
	user_id        *int64
	adduser_id     *int64
	created_at     *time.Time
	updated_at     *time.Time
	buckets        map[int]struct{}
	clearedBuckets bool
	predicates     []predicate.Files
}

// Where adds a new predicate for the builder.
func (fu *FilesUpdate) Where(ps ...predicate.Files) *FilesUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetFilename sets the filename field.
func (fu *FilesUpdate) SetFilename(s string) *FilesUpdate {
	fu.filename = &s
	return fu
}

// SetMimeType sets the mime_type field.
func (fu *FilesUpdate) SetMimeType(s string) *FilesUpdate {
	fu.mime_type = &s
	return fu
}

// SetFileSize sets the file_size field.
func (fu *FilesUpdate) SetFileSize(i int64) *FilesUpdate {
	fu.file_size = &i
	fu.addfile_size = nil
	return fu
}

// AddFileSize adds i to file_size.
func (fu *FilesUpdate) AddFileSize(i int64) *FilesUpdate {
	if fu.addfile_size == nil {
		fu.addfile_size = &i
	} else {
		*fu.addfile_size += i
	}
	return fu
}

// SetIsDraft sets the is_draft field.
func (fu *FilesUpdate) SetIsDraft(b bool) *FilesUpdate {
	fu.is_draft = &b
	return fu
}

// SetNillableIsDraft sets the is_draft field if the given value is not nil.
func (fu *FilesUpdate) SetNillableIsDraft(b *bool) *FilesUpdate {
	if b != nil {
		fu.SetIsDraft(*b)
	}
	return fu
}

// SetIsDeleted sets the is_deleted field.
func (fu *FilesUpdate) SetIsDeleted(b bool) *FilesUpdate {
	fu.is_deleted = &b
	return fu
}

// SetNillableIsDeleted sets the is_deleted field if the given value is not nil.
func (fu *FilesUpdate) SetNillableIsDeleted(b *bool) *FilesUpdate {
	if b != nil {
		fu.SetIsDeleted(*b)
	}
	return fu
}

// SetUserID sets the user_id field.
func (fu *FilesUpdate) SetUserID(i int64) *FilesUpdate {
	fu.user_id = &i
	fu.adduser_id = nil
	return fu
}

// AddUserID adds i to user_id.
func (fu *FilesUpdate) AddUserID(i int64) *FilesUpdate {
	if fu.adduser_id == nil {
		fu.adduser_id = &i
	} else {
		*fu.adduser_id += i
	}
	return fu
}

// SetCreatedAt sets the created_at field.
func (fu *FilesUpdate) SetCreatedAt(t time.Time) *FilesUpdate {
	fu.created_at = &t
	return fu
}

// SetUpdatedAt sets the updated_at field.
func (fu *FilesUpdate) SetUpdatedAt(t time.Time) *FilesUpdate {
	fu.updated_at = &t
	return fu
}

// SetBucketsID sets the buckets edge to Buckets by id.
func (fu *FilesUpdate) SetBucketsID(id int) *FilesUpdate {
	if fu.buckets == nil {
		fu.buckets = make(map[int]struct{})
	}
	fu.buckets[id] = struct{}{}
	return fu
}

// SetNillableBucketsID sets the buckets edge to Buckets by id if the given value is not nil.
func (fu *FilesUpdate) SetNillableBucketsID(id *int) *FilesUpdate {
	if id != nil {
		fu = fu.SetBucketsID(*id)
	}
	return fu
}

// SetBuckets sets the buckets edge to Buckets.
func (fu *FilesUpdate) SetBuckets(b *Buckets) *FilesUpdate {
	return fu.SetBucketsID(b.ID)
}

// ClearBuckets clears the buckets edge to Buckets.
func (fu *FilesUpdate) ClearBuckets() *FilesUpdate {
	fu.clearedBuckets = true
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FilesUpdate) Save(ctx context.Context) (int, error) {
	if fu.filename != nil {
		if err := files.FilenameValidator(*fu.filename); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"filename\": %v", err)
		}
	}
	if fu.mime_type != nil {
		if err := files.MimeTypeValidator(*fu.mime_type); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"mime_type\": %v", err)
		}
	}
	if len(fu.buckets) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"buckets\"")
	}
	return fu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FilesUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FilesUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FilesUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FilesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   files.Table,
			Columns: files.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: files.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := fu.filename; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldFilename,
		})
	}
	if value := fu.mime_type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldMimeType,
		})
	}
	if value := fu.file_size; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldFileSize,
		})
	}
	if value := fu.addfile_size; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldFileSize,
		})
	}
	if value := fu.is_draft; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDraft,
		})
	}
	if value := fu.is_deleted; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDeleted,
		})
	}
	if value := fu.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldUserID,
		})
	}
	if value := fu.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldUserID,
		})
	}
	if value := fu.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldCreatedAt,
		})
	}
	if value := fu.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldUpdatedAt,
		})
	}
	if fu.clearedBuckets {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   files.BucketsTable,
			Columns: []string{files.BucketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: buckets.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.buckets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   files.BucketsTable,
			Columns: []string{files.BucketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: buckets.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FilesUpdateOne is the builder for updating a single Files entity.
type FilesUpdateOne struct {
	config
	id             int
	filename       *string
	mime_type      *string
	file_size      *int64
	addfile_size   *int64
	is_draft       *bool
	is_deleted     *bool
	user_id        *int64
	adduser_id     *int64
	created_at     *time.Time
	updated_at     *time.Time
	buckets        map[int]struct{}
	clearedBuckets bool
}

// SetFilename sets the filename field.
func (fuo *FilesUpdateOne) SetFilename(s string) *FilesUpdateOne {
	fuo.filename = &s
	return fuo
}

// SetMimeType sets the mime_type field.
func (fuo *FilesUpdateOne) SetMimeType(s string) *FilesUpdateOne {
	fuo.mime_type = &s
	return fuo
}

// SetFileSize sets the file_size field.
func (fuo *FilesUpdateOne) SetFileSize(i int64) *FilesUpdateOne {
	fuo.file_size = &i
	fuo.addfile_size = nil
	return fuo
}

// AddFileSize adds i to file_size.
func (fuo *FilesUpdateOne) AddFileSize(i int64) *FilesUpdateOne {
	if fuo.addfile_size == nil {
		fuo.addfile_size = &i
	} else {
		*fuo.addfile_size += i
	}
	return fuo
}

// SetIsDraft sets the is_draft field.
func (fuo *FilesUpdateOne) SetIsDraft(b bool) *FilesUpdateOne {
	fuo.is_draft = &b
	return fuo
}

// SetNillableIsDraft sets the is_draft field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableIsDraft(b *bool) *FilesUpdateOne {
	if b != nil {
		fuo.SetIsDraft(*b)
	}
	return fuo
}

// SetIsDeleted sets the is_deleted field.
func (fuo *FilesUpdateOne) SetIsDeleted(b bool) *FilesUpdateOne {
	fuo.is_deleted = &b
	return fuo
}

// SetNillableIsDeleted sets the is_deleted field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableIsDeleted(b *bool) *FilesUpdateOne {
	if b != nil {
		fuo.SetIsDeleted(*b)
	}
	return fuo
}

// SetUserID sets the user_id field.
func (fuo *FilesUpdateOne) SetUserID(i int64) *FilesUpdateOne {
	fuo.user_id = &i
	fuo.adduser_id = nil
	return fuo
}

// AddUserID adds i to user_id.
func (fuo *FilesUpdateOne) AddUserID(i int64) *FilesUpdateOne {
	if fuo.adduser_id == nil {
		fuo.adduser_id = &i
	} else {
		*fuo.adduser_id += i
	}
	return fuo
}

// SetCreatedAt sets the created_at field.
func (fuo *FilesUpdateOne) SetCreatedAt(t time.Time) *FilesUpdateOne {
	fuo.created_at = &t
	return fuo
}

// SetUpdatedAt sets the updated_at field.
func (fuo *FilesUpdateOne) SetUpdatedAt(t time.Time) *FilesUpdateOne {
	fuo.updated_at = &t
	return fuo
}

// SetBucketsID sets the buckets edge to Buckets by id.
func (fuo *FilesUpdateOne) SetBucketsID(id int) *FilesUpdateOne {
	if fuo.buckets == nil {
		fuo.buckets = make(map[int]struct{})
	}
	fuo.buckets[id] = struct{}{}
	return fuo
}

// SetNillableBucketsID sets the buckets edge to Buckets by id if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableBucketsID(id *int) *FilesUpdateOne {
	if id != nil {
		fuo = fuo.SetBucketsID(*id)
	}
	return fuo
}

// SetBuckets sets the buckets edge to Buckets.
func (fuo *FilesUpdateOne) SetBuckets(b *Buckets) *FilesUpdateOne {
	return fuo.SetBucketsID(b.ID)
}

// ClearBuckets clears the buckets edge to Buckets.
func (fuo *FilesUpdateOne) ClearBuckets() *FilesUpdateOne {
	fuo.clearedBuckets = true
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FilesUpdateOne) Save(ctx context.Context) (*Files, error) {
	if fuo.filename != nil {
		if err := files.FilenameValidator(*fuo.filename); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"filename\": %v", err)
		}
	}
	if fuo.mime_type != nil {
		if err := files.MimeTypeValidator(*fuo.mime_type); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"mime_type\": %v", err)
		}
	}
	if len(fuo.buckets) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"buckets\"")
	}
	return fuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FilesUpdateOne) SaveX(ctx context.Context) *Files {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FilesUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FilesUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FilesUpdateOne) sqlSave(ctx context.Context) (f *Files, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   files.Table,
			Columns: files.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  fuo.id,
				Type:   field.TypeInt,
				Column: files.FieldID,
			},
		},
	}
	if value := fuo.filename; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldFilename,
		})
	}
	if value := fuo.mime_type; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: files.FieldMimeType,
		})
	}
	if value := fuo.file_size; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldFileSize,
		})
	}
	if value := fuo.addfile_size; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldFileSize,
		})
	}
	if value := fuo.is_draft; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDraft,
		})
	}
	if value := fuo.is_deleted; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: files.FieldIsDeleted,
		})
	}
	if value := fuo.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldUserID,
		})
	}
	if value := fuo.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: files.FieldUserID,
		})
	}
	if value := fuo.created_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldCreatedAt,
		})
	}
	if value := fuo.updated_at; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: files.FieldUpdatedAt,
		})
	}
	if fuo.clearedBuckets {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   files.BucketsTable,
			Columns: []string{files.BucketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: buckets.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.buckets; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   files.BucketsTable,
			Columns: []string{files.BucketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: buckets.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	f = &Files{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}