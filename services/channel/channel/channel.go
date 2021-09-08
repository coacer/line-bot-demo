package channel

import (
	"context"
	// model "channel/models/golang"
)

type Channel struct{}

func (c *Channel) Test(ctx context.Context, in *TestRequest) (*TestReply, error) {
	return &TestReply{Data: "hoge"}, nil
}

func (c *Channel) Get(ctx context.Context, in *GetRequest) (*GetReply, error) {
	// model.FindLineBotChannel(ctx, )
	return &GetReply{}, nil
}

func (c *Channel) Create(ctx context.Context, in *CreateRequest) (*CreateReply, error) {
	// model.FindLineBotChannel(ctx, )
	return &CreateReply{}, nil
}

func New() ChannelServer {
	return &Channel{}
}
