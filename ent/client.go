// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"oauth-az/ent/migrate"

	"oauth-az/ent/authcode"
	"oauth-az/ent/authrequest"
	"oauth-az/ent/oauth2client"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuthCode is the client for interacting with the AuthCode builders.
	AuthCode *AuthCodeClient
	// AuthRequest is the client for interacting with the AuthRequest builders.
	AuthRequest *AuthRequestClient
	// OAuth2Client is the client for interacting with the OAuth2Client builders.
	OAuth2Client *OAuth2ClientClient
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
	c.AuthCode = NewAuthCodeClient(c.config)
	c.AuthRequest = NewAuthRequestClient(c.config)
	c.OAuth2Client = NewOAuth2ClientClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		AuthCode:     NewAuthCodeClient(cfg),
		AuthRequest:  NewAuthRequestClient(cfg),
		OAuth2Client: NewOAuth2ClientClient(cfg),
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
		config:       cfg,
		AuthCode:     NewAuthCodeClient(cfg),
		AuthRequest:  NewAuthRequestClient(cfg),
		OAuth2Client: NewOAuth2ClientClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuthCode.
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
	c.AuthCode.Use(hooks...)
	c.AuthRequest.Use(hooks...)
	c.OAuth2Client.Use(hooks...)
}

// AuthCodeClient is a client for the AuthCode schema.
type AuthCodeClient struct {
	config
}

// NewAuthCodeClient returns a client for the AuthCode from the given config.
func NewAuthCodeClient(c config) *AuthCodeClient {
	return &AuthCodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authcode.Hooks(f(g(h())))`.
func (c *AuthCodeClient) Use(hooks ...Hook) {
	c.hooks.AuthCode = append(c.hooks.AuthCode, hooks...)
}

// Create returns a create builder for AuthCode.
func (c *AuthCodeClient) Create() *AuthCodeCreate {
	mutation := newAuthCodeMutation(c.config, OpCreate)
	return &AuthCodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthCode entities.
func (c *AuthCodeClient) CreateBulk(builders ...*AuthCodeCreate) *AuthCodeCreateBulk {
	return &AuthCodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthCode.
func (c *AuthCodeClient) Update() *AuthCodeUpdate {
	mutation := newAuthCodeMutation(c.config, OpUpdate)
	return &AuthCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthCodeClient) UpdateOne(ac *AuthCode) *AuthCodeUpdateOne {
	mutation := newAuthCodeMutation(c.config, OpUpdateOne, withAuthCode(ac))
	return &AuthCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthCodeClient) UpdateOneID(id int) *AuthCodeUpdateOne {
	mutation := newAuthCodeMutation(c.config, OpUpdateOne, withAuthCodeID(id))
	return &AuthCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthCode.
func (c *AuthCodeClient) Delete() *AuthCodeDelete {
	mutation := newAuthCodeMutation(c.config, OpDelete)
	return &AuthCodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthCodeClient) DeleteOne(ac *AuthCode) *AuthCodeDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthCodeClient) DeleteOneID(id int) *AuthCodeDeleteOne {
	builder := c.Delete().Where(authcode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthCodeDeleteOne{builder}
}

// Query returns a query builder for AuthCode.
func (c *AuthCodeClient) Query() *AuthCodeQuery {
	return &AuthCodeQuery{
		config: c.config,
	}
}

// Get returns a AuthCode entity by its id.
func (c *AuthCodeClient) Get(ctx context.Context, id int) (*AuthCode, error) {
	return c.Query().Where(authcode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthCodeClient) GetX(ctx context.Context, id int) *AuthCode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthCodeClient) Hooks() []Hook {
	return c.hooks.AuthCode
}

// AuthRequestClient is a client for the AuthRequest schema.
type AuthRequestClient struct {
	config
}

// NewAuthRequestClient returns a client for the AuthRequest from the given config.
func NewAuthRequestClient(c config) *AuthRequestClient {
	return &AuthRequestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authrequest.Hooks(f(g(h())))`.
func (c *AuthRequestClient) Use(hooks ...Hook) {
	c.hooks.AuthRequest = append(c.hooks.AuthRequest, hooks...)
}

// Create returns a create builder for AuthRequest.
func (c *AuthRequestClient) Create() *AuthRequestCreate {
	mutation := newAuthRequestMutation(c.config, OpCreate)
	return &AuthRequestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthRequest entities.
func (c *AuthRequestClient) CreateBulk(builders ...*AuthRequestCreate) *AuthRequestCreateBulk {
	return &AuthRequestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthRequest.
func (c *AuthRequestClient) Update() *AuthRequestUpdate {
	mutation := newAuthRequestMutation(c.config, OpUpdate)
	return &AuthRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthRequestClient) UpdateOne(ar *AuthRequest) *AuthRequestUpdateOne {
	mutation := newAuthRequestMutation(c.config, OpUpdateOne, withAuthRequest(ar))
	return &AuthRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthRequestClient) UpdateOneID(id int) *AuthRequestUpdateOne {
	mutation := newAuthRequestMutation(c.config, OpUpdateOne, withAuthRequestID(id))
	return &AuthRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthRequest.
func (c *AuthRequestClient) Delete() *AuthRequestDelete {
	mutation := newAuthRequestMutation(c.config, OpDelete)
	return &AuthRequestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthRequestClient) DeleteOne(ar *AuthRequest) *AuthRequestDeleteOne {
	return c.DeleteOneID(ar.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthRequestClient) DeleteOneID(id int) *AuthRequestDeleteOne {
	builder := c.Delete().Where(authrequest.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthRequestDeleteOne{builder}
}

// Query returns a query builder for AuthRequest.
func (c *AuthRequestClient) Query() *AuthRequestQuery {
	return &AuthRequestQuery{
		config: c.config,
	}
}

// Get returns a AuthRequest entity by its id.
func (c *AuthRequestClient) Get(ctx context.Context, id int) (*AuthRequest, error) {
	return c.Query().Where(authrequest.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthRequestClient) GetX(ctx context.Context, id int) *AuthRequest {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthRequestClient) Hooks() []Hook {
	return c.hooks.AuthRequest
}

// OAuth2ClientClient is a client for the OAuth2Client schema.
type OAuth2ClientClient struct {
	config
}

// NewOAuth2ClientClient returns a client for the OAuth2Client from the given config.
func NewOAuth2ClientClient(c config) *OAuth2ClientClient {
	return &OAuth2ClientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `oauth2client.Hooks(f(g(h())))`.
func (c *OAuth2ClientClient) Use(hooks ...Hook) {
	c.hooks.OAuth2Client = append(c.hooks.OAuth2Client, hooks...)
}

// Create returns a create builder for OAuth2Client.
func (c *OAuth2ClientClient) Create() *OAuth2ClientCreate {
	mutation := newOAuth2ClientMutation(c.config, OpCreate)
	return &OAuth2ClientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OAuth2Client entities.
func (c *OAuth2ClientClient) CreateBulk(builders ...*OAuth2ClientCreate) *OAuth2ClientCreateBulk {
	return &OAuth2ClientCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OAuth2Client.
func (c *OAuth2ClientClient) Update() *OAuth2ClientUpdate {
	mutation := newOAuth2ClientMutation(c.config, OpUpdate)
	return &OAuth2ClientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OAuth2ClientClient) UpdateOne(o *OAuth2Client) *OAuth2ClientUpdateOne {
	mutation := newOAuth2ClientMutation(c.config, OpUpdateOne, withOAuth2Client(o))
	return &OAuth2ClientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OAuth2ClientClient) UpdateOneID(id int) *OAuth2ClientUpdateOne {
	mutation := newOAuth2ClientMutation(c.config, OpUpdateOne, withOAuth2ClientID(id))
	return &OAuth2ClientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OAuth2Client.
func (c *OAuth2ClientClient) Delete() *OAuth2ClientDelete {
	mutation := newOAuth2ClientMutation(c.config, OpDelete)
	return &OAuth2ClientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OAuth2ClientClient) DeleteOne(o *OAuth2Client) *OAuth2ClientDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OAuth2ClientClient) DeleteOneID(id int) *OAuth2ClientDeleteOne {
	builder := c.Delete().Where(oauth2client.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OAuth2ClientDeleteOne{builder}
}

// Query returns a query builder for OAuth2Client.
func (c *OAuth2ClientClient) Query() *OAuth2ClientQuery {
	return &OAuth2ClientQuery{
		config: c.config,
	}
}

// Get returns a OAuth2Client entity by its id.
func (c *OAuth2ClientClient) Get(ctx context.Context, id int) (*OAuth2Client, error) {
	return c.Query().Where(oauth2client.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OAuth2ClientClient) GetX(ctx context.Context, id int) *OAuth2Client {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OAuth2ClientClient) Hooks() []Hook {
	return c.hooks.OAuth2Client
}