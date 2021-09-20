package repository

import (
	"context"
	"time"
)

type SqlClient interface {
	GetByIds(ctx context.Context, table string, ids []string) Row
	GetAll(ctx context.Context, table string) Row
	Query(ctx context.Context, query string) Row
	Insert(ctx context.Context, table string, columns []string, values []interface{}) time.Time
	Update(ctx context.Context, table string, columns []string, values []interface{}) time.Time
	Delete(ctx context.Context, table string, id string) time.Time
}

type Row interface {
	Bind(args ...interface{}) error
	Loop(func(row interface{}) error) error
}
