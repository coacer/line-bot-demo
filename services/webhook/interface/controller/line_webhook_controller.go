package controller

import (
	"context"
	"log"
	"webhook/app/interactor"
	"webhook/infrastructure/grpc/linewebhook/pb"
)

type LineMessageController struct {
	interactor *interactor.NewsInteractor
}

func NewLineMessageController(i *interactor.NewsInteractor) *LineMessageController {
	return &LineMessageController{i}
}

func (c *LineMessageController) Message(ctx context.Context, req *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Println(req)
	channelId := "1111111111"
	uid := req.Event.Source.UserId
	u, err := c.interactor.Reply(ctx, channelId, uid)
	log.Println(u)
	if err != nil {
		return nil, err
	}
	return &pb.MessageReply{}, nil
}
