package repository

import (
	"context"
	"time"
)

type Sql interface {
	NewClient(ctx context.Context) (SqlClient, error)
	Close()

	CommitTimestamp() time.Time
}

type SqlClient interface {
	// Read
	Read(ctx context.Context, table string, ids []string, columns []string) Result
	ReadAll(ctx context.Context, table string, columns []string) Result
	// Query(ctx context.Context, query string, columns []string) (Result error)

	// Write
	Insert(ctx context.Context, table string, columns []string, values []interface{}) WriteQuery
	Update(ctx context.Context, table string, columns []string, values []interface{}) WriteQuery
	Delete(ctx context.Context, table string, id string) WriteQuery
	// write execution
	Commit(ctx context.Context, queries []WriteQuery) (time.Time, error)
	// Transaction
	DoTxn(ctx context.Context, callback func(ctx context.Context, txn Transaction) error) (time.Time, error)
}

type Result interface {
	Loop(callback func(row Row) error) error
}

type Row interface {
	Bind(args ...interface{}) error
}

type WriteQuery interface{}

type Transaction interface {
	// Read
	Read(ctx context.Context, table string, ids []string, columns []string) Result
	ReadAll(ctx context.Context, table string, columns []string) Result
	// Query(ctx context.Context, query string, columns []string) (Result error)

	// Write
	Insert(ctx context.Context, table string, columns []string, values []interface{}) WriteQuery
	Update(ctx context.Context, table string, columns []string, values []interface{}) WriteQuery
	Delete(ctx context.Context, table string, id string) WriteQuery
	// write execution
	Commit(ctx context.Context, queries []WriteQuery) error
}
