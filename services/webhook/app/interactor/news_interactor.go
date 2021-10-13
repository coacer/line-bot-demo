package interactor

import (
	"context"
	"webhook/domain/model"
	"webhook/domain/repository"
)

type NewsInteractor struct {
	userRepository repository.UserRepository
	lineRepository repository.LineRepository
}

func NewNewsInteractor(ur repository.UserRepository, lr repository.LineRepository) *NewsInteractor {
	return &NewsInteractor{
		userRepository: ur,
		lineRepository: lr,
	}
}

// TODO news api
func (i *NewsInteractor) Reply(ctx context.Context, channelId string, uid string, replyToken string) (*model.User, error) {
	u, err := model.NewUser(channelId, uid)
	if err != nil {
		return nil, err
	}
	err = i.userRepository.Store(ctx, u)

	i.lineRepository.ReplyTextMessages(replyToken, "test", "dayo!!")
	return u, err
}
