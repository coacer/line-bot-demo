package models

//
// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"
//
// 	"cloud.google.com/go/spanner"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )
//
// type Spanner struct {
// 	client *spanner.Client
// }
//
// func NewClient(ctx context.Context) (*Spanner, error) {
// 	s := &Spanner{}
// 	err := s.connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s, nil
// }
//
// func (s *Spanner) Close() {
// 	s.client.Close()
// }
//
// func (s *Spanner) Insert(ctx context.Context, model interface{}) (commitTimeStamp time.Time, err error) {
// 	m, ok := model.(struct {
// 		ID        string
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 	})
// 	if !ok {
// 		err = errors.New("")
// 		return nil, err
// 	}
// 	m.ID = GeneratePrimaryKey()
// 	m.CreatedAt = spanner.CommitTimestamp
// 	m.UpdatedAt = spanner.CommitTimestamp
// 	// s.client.Apply(ctx, )
// 	return time.Now(), nil
// }
//
// func (s *Spanner) connect(ctx context.Context) error {
// 	client, err := spanner.NewClient(ctx,
// 		fmt.Sprintf(
// 			"projects/%s/instances/%s/databases/%s",
// 			os.Getenv("GCP_PROJECT_ID"),
// 			os.Getenv("DB_INSTANCE_NAME"),
// 			os.Getenv("DB_NAME"),
// 		))
// 	s.client = client
// 	if err != nil {
// 		log.Fatalln(err)
// 		return status.Errorf(codes.Internal, "Failed connect spanner.")
// 	}
// 	return nil
// }
