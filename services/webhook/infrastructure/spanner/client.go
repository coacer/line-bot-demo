package spanner

import (
	"context"
	"fmt"
	"os"
	"webhook/exception"

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

func NewClient(ctx context.Context) (c *Client, err error) {
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
