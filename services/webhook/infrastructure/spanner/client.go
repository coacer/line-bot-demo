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

func (c *Client) GetByIds(ctx context.Context, table string, ids []string, columns []string) repository.Result {
	keys := make([]spanner.KeySet, len(ids))
	for i, id := range ids {
		keys[i] = spanner.Key{id}
	}
	iter := c.spanner.Single().Read(ctx, table, spanner.KeySets(keys...), columns)
	return &Result{iter}
}

func (c *Client) GetAll(ctx context.Context, table string, columns []string) repository.Result {
	iter := c.spanner.Single().Read(ctx, table, spanner.AllKeys(), columns)
	return &Result{iter}
}

// TODO
// func (c *Client) Query(ctx context.Context, query string) repository.Result  {}

func (c *Client) Insert(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	mu := spanner.Insert(table, columns, values)
	return &WriteQuery{mu}
}

func (c *Client) Update(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	mu := spanner.Update(table, columns, values)
	return &WriteQuery{mu}
}

func (c *Client) Delete(ctx context.Context, table string, id string) repository.WriteQuery {
	mu := spanner.Delete(table, spanner.Key{id})
	return &WriteQuery{mu}
}

func (c *Client) Commit(ctx context.Context, query repository.WriteQuery) (time.Time, error) {
	return c.spanner.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		return txn.BufferWrite([]*spanner.Mutation{query.(*WriteQuery).mutation})
	})
}

func (c *Client) Transaction(ctx context.Context, callback func(ctx context.Context) ([]repository.WriteQuery, error)) (time.Time, error) {
	return c.spanner.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		queries, err := callback(ctx)
		if err != nil {
			return err
		}
		ms := make([]*spanner.Mutation, len(queries))
		for _, q := range queries {
			ms = append(ms, q.(WriteQuery).mutation)
		}
		return txn.BufferWrite(ms)
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
