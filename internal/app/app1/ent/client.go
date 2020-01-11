// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/pepeunlimited/files/internal/app/app1/ent/migrate"

	"github.com/pepeunlimited/files/internal/app/app1/ent/files"
	"github.com/pepeunlimited/files/internal/app/app1/ent/spaces"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Files is the client for interacting with the Files builders.
	Files *FilesClient
	// Spaces is the client for interacting with the Spaces builders.
	Spaces *SpacesClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Files:  NewFilesClient(c),
		Spaces: NewSpacesClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config: cfg,
		Files:  NewFilesClient(cfg),
		Spaces: NewSpacesClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Files.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config: cfg,
		Schema: migrate.NewSchema(cfg.driver),
		Files:  NewFilesClient(cfg),
		Spaces: NewSpacesClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// FilesClient is a client for the Files schema.
type FilesClient struct {
	config
}

// NewFilesClient returns a client for the Files from the given config.
func NewFilesClient(c config) *FilesClient {
	return &FilesClient{config: c}
}

// Create returns a create builder for Files.
func (c *FilesClient) Create() *FilesCreate {
	return &FilesCreate{config: c.config}
}

// Update returns an update builder for Files.
func (c *FilesClient) Update() *FilesUpdate {
	return &FilesUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *FilesClient) UpdateOne(f *Files) *FilesUpdateOne {
	return c.UpdateOneID(f.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *FilesClient) UpdateOneID(id int) *FilesUpdateOne {
	return &FilesUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Files.
func (c *FilesClient) Delete() *FilesDelete {
	return &FilesDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FilesClient) DeleteOne(f *Files) *FilesDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FilesClient) DeleteOneID(id int) *FilesDeleteOne {
	return &FilesDeleteOne{c.Delete().Where(files.ID(id))}
}

// Create returns a query builder for Files.
func (c *FilesClient) Query() *FilesQuery {
	return &FilesQuery{config: c.config}
}

// Get returns a Files entity by its id.
func (c *FilesClient) Get(ctx context.Context, id int) (*Files, error) {
	return c.Query().Where(files.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FilesClient) GetX(ctx context.Context, id int) *Files {
	f, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return f
}

// QuerySpaces queries the spaces edge of a Files.
func (c *FilesClient) QuerySpaces(f *Files) *SpacesQuery {
	query := &SpacesQuery{config: c.config}
	id := f.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(files.Table, files.FieldID, id),
		sqlgraph.To(spaces.Table, spaces.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, files.SpacesTable, files.SpacesColumn),
	)
	query.sql = sqlgraph.Neighbors(f.driver.Dialect(), step)

	return query
}

// SpacesClient is a client for the Spaces schema.
type SpacesClient struct {
	config
}

// NewSpacesClient returns a client for the Spaces from the given config.
func NewSpacesClient(c config) *SpacesClient {
	return &SpacesClient{config: c}
}

// Create returns a create builder for Spaces.
func (c *SpacesClient) Create() *SpacesCreate {
	return &SpacesCreate{config: c.config}
}

// Update returns an update builder for Spaces.
func (c *SpacesClient) Update() *SpacesUpdate {
	return &SpacesUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SpacesClient) UpdateOne(s *Spaces) *SpacesUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SpacesClient) UpdateOneID(id int) *SpacesUpdateOne {
	return &SpacesUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Spaces.
func (c *SpacesClient) Delete() *SpacesDelete {
	return &SpacesDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SpacesClient) DeleteOne(s *Spaces) *SpacesDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SpacesClient) DeleteOneID(id int) *SpacesDeleteOne {
	return &SpacesDeleteOne{c.Delete().Where(spaces.ID(id))}
}

// Create returns a query builder for Spaces.
func (c *SpacesClient) Query() *SpacesQuery {
	return &SpacesQuery{config: c.config}
}

// Get returns a Spaces entity by its id.
func (c *SpacesClient) Get(ctx context.Context, id int) (*Spaces, error) {
	return c.Query().Where(spaces.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SpacesClient) GetX(ctx context.Context, id int) *Spaces {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryFiles queries the files edge of a Spaces.
func (c *SpacesClient) QueryFiles(s *Spaces) *FilesQuery {
	query := &FilesQuery{config: c.config}
	id := s.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(spaces.Table, spaces.FieldID, id),
		sqlgraph.To(files.Table, files.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, spaces.FilesTable, spaces.FilesColumn),
	)
	query.sql = sqlgraph.Neighbors(s.driver.Dialect(), step)

	return query
}
