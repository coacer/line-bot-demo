package spanner

import (
	"context"
	"fmt"
	"os"
	"time"
	"webhook/exception"
	"webhook/interface/repository"

	"cloud.google.com/go/spanner"
)

var (
	project  = os.Getenv("GCP_PROJECT_ID")
	instance = os.Getenv("DB_INSTANCE_NAME")
	dbName   = os.Getenv("DB_NAME")
)

type Client struct {
	spannerClient *spanner.Client
}

func NewClient(ctx context.Context) (c repository.SqlClient, err error) {
	client, err := spanner.NewClient(
		ctx,
		fmt.Sprintf(
			"projects/%s/instances/%s/databases/%s",
			project,
			instance,
			dbName,
		),
	)
	if err != nil {
		err = exception.NewError(exception.DatabaseConnectionError, err)
	}
	c = &Client{client}
	return
}

func (c *Client) GetByIds(ctx context.Context, table string, ids []string) repository.Row {}
func (c *Client) GetAll(ctx context.Context, table string) repository.Row                 {}
func (c *Client) Query(ctx context.Context, query string) repository.Row                  {}
func (c *Client) Insert(ctx context.Context, table string, columns []string, values []interface{}) time.Time {
}
func (c *Client) Update(ctx context.Context, table string, columns []string, values []interface{}) time.Time {
}
func (c *Client) Delete(ctx context.Context, table string, id string) time.Time {}
