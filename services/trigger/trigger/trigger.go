package trigger

import (
	"context"
)

type Trigger struct{}

func (t *Trigger) Test(ctx context.Context, in *TestRequest) (*TestReply, error) {
	return &TestReply{Data: "hoge"}, nil
}

func New() TriggerServer {
	return &Trigger{}
}
