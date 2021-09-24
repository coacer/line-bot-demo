package repository

import (
	"context"
	"time"
)

type DBRepository interface {
	DoTxn(ctx context.Context, callback func(ctx context.Context) error) (time.Time, error)
}
