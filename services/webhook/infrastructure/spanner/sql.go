package spanner

import (
	"time"
	"webhook/interface/repository"

	"cloud.google.com/go/spanner"
)

// TODO 外から注入
// var (
// 	project  = os.Getenv("GCP_PROJECT_ID")
// 	instance = os.Getenv("DB_INSTANCE_NAME")
// 	dbName   = os.Getenv("DB_NAME")
// )

type Sql struct {
	project   string
	instance  string
	dbName    string
	client    *spanner.Client
	connected bool
}

func NewSql(project string, instance string, dbName string) repository.Sql {
	return &Sql{project: project, instance: instance, dbName: dbName}
}

func (s *Sql) CommitTimestamp() time.Time {
	return spanner.CommitTimestamp
}

func (s *Sql) Close() {
	if s.client == nil {
		return
	}
	s.client.Close()
}
