// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent/migrate"

	"github.com/genglsh/go-trainning-map/app/people/service/internal/data/ent/people"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// People is the client for interacting with the People builders.
	People *PeopleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.People = NewPeopleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		People: NewPeopleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		People: NewPeopleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		People.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.People.Use(hooks...)
}

// PeopleClient is a client for the People schema.
type PeopleClient struct {
	config
}

// NewPeopleClient returns a client for the People from the given config.
func NewPeopleClient(c config) *PeopleClient {
	return &PeopleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `people.Hooks(f(g(h())))`.
func (c *PeopleClient) Use(hooks ...Hook) {
	c.hooks.People = append(c.hooks.People, hooks...)
}

// Create returns a create builder for People.
func (c *PeopleClient) Create() *PeopleCreate {
	mutation := newPeopleMutation(c.config, OpCreate)
	return &PeopleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of People entities.
func (c *PeopleClient) CreateBulk(builders ...*PeopleCreate) *PeopleCreateBulk {
	return &PeopleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for People.
func (c *PeopleClient) Update() *PeopleUpdate {
	mutation := newPeopleMutation(c.config, OpUpdate)
	return &PeopleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PeopleClient) UpdateOne(pe *People) *PeopleUpdateOne {
	mutation := newPeopleMutation(c.config, OpUpdateOne, withPeople(pe))
	return &PeopleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PeopleClient) UpdateOneID(id int) *PeopleUpdateOne {
	mutation := newPeopleMutation(c.config, OpUpdateOne, withPeopleID(id))
	return &PeopleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for People.
func (c *PeopleClient) Delete() *PeopleDelete {
	mutation := newPeopleMutation(c.config, OpDelete)
	return &PeopleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PeopleClient) DeleteOne(pe *People) *PeopleDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PeopleClient) DeleteOneID(id int) *PeopleDeleteOne {
	builder := c.Delete().Where(people.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PeopleDeleteOne{builder}
}

// Query returns a query builder for People.
func (c *PeopleClient) Query() *PeopleQuery {
	return &PeopleQuery{
		config: c.config,
	}
}

// Get returns a People entity by its id.
func (c *PeopleClient) Get(ctx context.Context, id int) (*People, error) {
	return c.Query().Where(people.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PeopleClient) GetX(ctx context.Context, id int) *People {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PeopleClient) Hooks() []Hook {
	return c.hooks.People
}
