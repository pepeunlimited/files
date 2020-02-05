// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/pepeunlimited/files/internal/pkg/ent/file"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BucketsColumns holds the columns for the "buckets" table.
	BucketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 63},
		{Name: "endpoint", Type: field.TypeString, Size: 35},
		{Name: "cdn_endpoint", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// BucketsTable holds the schema information for the "buckets" table.
	BucketsTable = &schema.Table{
		Name:        "buckets",
		Columns:     BucketsColumns,
		PrimaryKey:  []*schema.Column{BucketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "filename", Type: field.TypeString, Size: 100},
		{Name: "mime_type", Type: field.TypeString, Size: 255},
		{Name: "file_size", Type: field.TypeInt64},
		{Name: "is_draft", Type: field.TypeBool, Default: file.DefaultIsDraft},
		{Name: "is_deleted", Type: field.TypeBool, Default: file.DefaultIsDeleted},
		{Name: "user_id", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "bucket_files", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "files_buckets_files",
				Columns: []*schema.Column{FilesColumns[9]},

				RefColumns: []*schema.Column{BucketsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BucketsTable,
		FilesTable,
	}
)

func init() {
	FilesTable.ForeignKeys[0].RefTable = BucketsTable
}
