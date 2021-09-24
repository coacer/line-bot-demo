package spanner

import (
	"context"
	"webhook/interface/repository"

	"cloud.google.com/go/spanner"
)

type Transaction struct {
	roTxn *spanner.ReadOnlyTransaction
	rwTxn *spanner.ReadWriteTransaction
}

func (t *Transaction) Read(ctx context.Context, table string, ids []string, columns []string) repository.Result {
	keys := make([]spanner.KeySet, len(ids))
	for i, id := range ids {
		keys[i] = spanner.Key{id}
	}
	iter := t.rwTxn.Read(ctx, table, spanner.KeySets(keys...), columns)
	return &Result{iter}
}

func (t *Transaction) ReadAll(ctx context.Context, table string, columns []string) repository.Result {
	iter := t.rwTxn.Read(ctx, table, spanner.AllKeys(), columns)
	return &Result{iter}
}

// TODO
// func (t *Transaction) Query(ctx context.Context, query string) repository.Result  {}

func (t *Transaction) Insert(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	mu := spanner.Insert(table, columns, values)
	return &WriteQuery{mu}
}

func (t *Transaction) Update(ctx context.Context, table string, columns []string, values []interface{}) repository.WriteQuery {
	mu := spanner.Update(table, columns, values)
	return &WriteQuery{mu}
}

func (t *Transaction) Delete(ctx context.Context, table string, id string) repository.WriteQuery {
	mu := spanner.Delete(table, spanner.Key{id})
	return &WriteQuery{mu}
}

func (t *Transaction) Commit(ctx context.Context, queries []repository.WriteQuery) error {
	ms := make([]*spanner.Mutation, len(queries))
	for i, q := range queries {
		ms[i] = q.(*WriteQuery).mutation
	}
	return t.rwTxn.BufferWrite(ms)
}
