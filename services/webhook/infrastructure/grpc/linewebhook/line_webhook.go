package linewebhook

import (
	"context"
)

type LineWebhook struct{}

func New() LineWebhookServer {
	return new(LineWebhook)
}

func (w *LineWebhook) Health(ctx context.Context, in *HealthRequest) (*HealthReply, error) {
	return &HealthReply{Data: "ok"}, nil
}

func (w *LineWebhook) Message(ctx context.Context, in *MessageRequest) (*MessageReply, error) {
	return &MessageReply{}, nil
}
