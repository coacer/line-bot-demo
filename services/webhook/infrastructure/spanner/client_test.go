package spanner

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"webhook/interface/repository"
)

var (
	project  = os.Getenv("GCP_PROJECT_ID")
	instance = os.Getenv("DB_INSTANCE_NAME")
	dbName   = os.Getenv("DB_NAME")

	sql       repository.Sql
	sqlClient repository.SqlClient
	ctx       context.Context
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	ctx = context.Background()
	sql = NewSql(project, instance, dbName)
	sqlClient, _ = sql.NewClient(ctx)
}

func TestInsertUser(t *testing.T) {
	patterns := []struct {
		id               string
		lineBotChannelId string
		lineUID          string
	}{
		{"a", "a", "a"},
		{"aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa"},
	}

	for _, p := range patterns {
		var values []interface{}
		values = append(values, p.id, p.lineBotChannelId, p.lineUID, sql.CommitTimestamp(), sql.CommitTimestamp())
		query := sqlClient.Insert(ctx, "User", []string{"Id", "LineBotChannelId", "LineUID", "CreatedAt", "UpdatedAt"}, values)
		_, err := sqlClient.Commit(ctx, []repository.WriteQuery{query})
		if err != nil {
			t.Fatalf("error: %v", err)
		}
	}

	var id string
	var channelId string
	var uid string
	var createdat time.Time
	var updatedat time.Time
	result := sqlClient.ReadAll(ctx, "User", []string{"Id", "LineBotChannelId", "LineUID", "CreatedAt", "UpdatedAt"})
	result.Loop(func(row repository.Row) error {
		return row.Bind(&id, &channelId, &uid, &createdat, &updatedat)
	})
	fmt.Println(id, channelId, uid, createdat)
}
