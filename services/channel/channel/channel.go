package channel

import (
	"channel/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Channel struct{}

func (c *Channel) Test(ctx context.Context, in *TestRequest) (*TestReply, error) {
	return &TestReply{Data: "Successfully channel gRPC testing!"}, nil
}

func (c *Channel) Get(ctx context.Context, in *GetRequest) (*GetReply, error) {
	// model.FindLineBotChannel(ctx, )
	return &GetReply{}, nil
}

func (c *Channel) Create(ctx context.Context, in *CreateRequest) (*CreateReply, error) {
	id := models.GeneratePrimaryKey()
	lbc := models.LineBotChannel{
		ID:                 id,
		ChannelID:          in.ChannelId,
		ChannelSecretID:    in.ChannelSecretId,
		ChannelAccessToken: in.ChannelAccessToken,
		CreatedAt:          spanner.CommitTimestamp,
		UpdatedAt:          spanner.CommitTimestamp,
	}
	client, err := spanner.NewClient(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s", os.Getenv("GCP_PROJECT_ID"), os.Getenv("DB_INSTANCE_NAME"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.Internal, "Failed connect spanner.")
	}
	defer client.Close()

	cts, err := client.Apply(ctx, []*spanner.Mutation{lbc.Insert(ctx)})
	if err != nil {
		log.Fatalln(err)
		return nil, status.Errorf(codes.Internal, "Failed apply insert.")
	}
	return &CreateReply{Channel: &ChannelModel{
		ID:                 lbc.ID,
		ChannelId:          lbc.ChannelID,
		ChannelSecretId:    lbc.ChannelSecretID,
		ChannelAccessToken: lbc.ChannelAccessToken,
		CreatedAt:          cts.Format(time.RFC3339),
		UpdatedAt:          cts.Format(time.RFC3339),
	}}, nil
}

func New() ChannelServer {
	return &Channel{}
}
