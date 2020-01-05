// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DoBucketsColumns holds the columns for the "do_buckets" table.
	DoBucketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bucket_name", Type: field.TypeString, Unique: true, Size: 63},
		{Name: "endpoint", Type: field.TypeString, Size: 35},
		{Name: "cdn_endpoint", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// DoBucketsTable holds the schema information for the "do_buckets" table.
	DoBucketsTable = &schema.Table{
		Name:        "do_buckets",
		Columns:     DoBucketsColumns,
		PrimaryKey:  []*schema.Column{DoBucketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DoBucketsTable,
	}
)

func init() {
}
