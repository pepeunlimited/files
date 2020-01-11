// Code generated by entc, DO NOT EDIT.

package files

import (
	"github.com/pepeunlimited/files/internal/app/app1/ent/schema"
)

const (
	// Label holds the string label denoting the files type in the database.
	Label = "files"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilename holds the string denoting the filename vertex property in the database.
	FieldFilename = "filename"
	// FieldMimeType holds the string denoting the mime_type vertex property in the database.
	FieldMimeType = "mime_type"
	// FieldFileSize holds the string denoting the file_size vertex property in the database.
	FieldFileSize = "file_size"
	// FieldIsDraft holds the string denoting the is_draft vertex property in the database.
	FieldIsDraft = "is_draft"
	// FieldIsDeleted holds the string denoting the is_deleted vertex property in the database.
	FieldIsDeleted = "is_deleted"
	// FieldUserID holds the string denoting the user_id vertex property in the database.
	FieldUserID = "user_id"
	// FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at"

	// Table holds the table name of the files in the database.
	Table = "files"
	// SpacesTable is the table the holds the spaces relation/edge.
	SpacesTable = "files"
	// SpacesInverseTable is the table name for the Spaces entity.
	// It exists in this package in order to avoid circular dependency with the "spaces" package.
	SpacesInverseTable = "spaces"
	// SpacesColumn is the table column denoting the spaces relation/edge.
	SpacesColumn = "spaces_id"
)

// Columns holds all SQL columns are files fields.
var Columns = []string{
	FieldID,
	FieldFilename,
	FieldMimeType,
	FieldFileSize,
	FieldIsDraft,
	FieldIsDeleted,
	FieldUserID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	fields = schema.Files{}.Fields()

	// descFilename is the schema descriptor for filename field.
	descFilename = fields[0].Descriptor()
	// FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	FilenameValidator = func() func(string) error {
		validators := descFilename.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(filename string) error {
			for _, fn := range fns {
				if err := fn(filename); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descMimeType is the schema descriptor for mime_type field.
	descMimeType = fields[1].Descriptor()
	// MimeTypeValidator is a validator for the "mime_type" field. It is called by the builders before save.
	MimeTypeValidator = func() func(string) error {
		validators := descMimeType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mime_type string) error {
			for _, fn := range fns {
				if err := fn(mime_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descIsDraft is the schema descriptor for is_draft field.
	descIsDraft = fields[3].Descriptor()
	// DefaultIsDraft holds the default value on creation for the is_draft field.
	DefaultIsDraft = descIsDraft.Default.(bool)

	// descIsDeleted is the schema descriptor for is_deleted field.
	descIsDeleted = fields[4].Descriptor()
	// DefaultIsDeleted holds the default value on creation for the is_deleted field.
	DefaultIsDeleted = descIsDeleted.Default.(bool)
)