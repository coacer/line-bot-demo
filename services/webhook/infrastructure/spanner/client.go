package spanner

import (
	"context"
	"fmt"
	"time"
	"webhook/exception"
	"webhook/interface/repository"

	"cloud.google.com/go/spanner"
)

type Client struct {
	spanner *spanner.Client
}

func (s *Sql) NewClient(ctx context.Context) (c repository.SqlClient, err error) {
	if s.connected {
		c = &Client{s.client}
		return
	}

	client, err := spanner.NewClient(
		ctx,
		fmt.Sprintf(
			"projects/%s/instances/%s/databases/%s",
			s.project,
			s.instance,
			s.dbName,
		),
	)
	if err != nil {
		err = exception.NewError(exception.DatabaseConnectionError, err)
	}
	c = &Client{client}
	s.client = client
	s.connected = true
	return
}

func (c *Client) Read(ctx context.Context, table string, ids []string, columns []string) repository.Result {
	var t repository.Transaction = &Transaction{roTxn: c.spanner.Single()}
	return t.Read(ctx, table, ids, columns)
}

func (c *Client) ReadAll(ctx context.Context, table string, columns []string) repository.Result {
	var t repository.Transaction = &Transaction{roTxn: c.spanner.Single()}
	return t.ReadAll(ctx, table, columns)
}

// TODO
// func (c *Client) Query(ctx context.Context, query string) repository.Result  {}

func (c *Client) Insert(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	var t repository.Transaction = &Transaction{}
	return t.Insert(ctx, table, columns, values)
}

func (c *Client) Update(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	var t repository.Transaction = &Transaction{}
	return t.Update(ctx, table, columns, values)
}

func (c *Client) Delete(ctx context.Context, table string, id string) repository.WriteQuery {
	var t repository.Transaction = &Transaction{}
	return t.Delete(ctx, table, id)
}

func (c *Client) Commit(ctx context.Context, queries []repository.WriteQuery) (time.Time, error) {
	return c.DoTxn(ctx, func(ctx context.Context, t repository.Transaction) error {
		return t.Commit(ctx, queries)
	})
}

func (c *Client) DoTxn(ctx context.Context, callback func(ctx context.Context, txn repository.Transaction) error) (time.Time, error) {
	return c.spanner.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		var t repository.Transaction = &Transaction{rwTxn: txn}
		return callback(ctx, t)
	})
}

type Result struct {
	iter *spanner.RowIterator
}

type Row struct {
	row *spanner.Row
}

func (r *Row) Bind(args ...interface{}) error {
	return r.row.Columns(args...)
}

func (r *Result) Loop(callback func(row repository.Row) error) error {
	return r.iter.Do(func(r *spanner.Row) error {
		row := &Row{r}
		return callback(row)
	})
}

type WriteQuery struct {
	mutation *spanner.Mutation
}
