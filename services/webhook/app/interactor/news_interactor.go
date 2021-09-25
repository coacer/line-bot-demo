package interactor

import (
	"context"
	"webhook/domain/model"
	"webhook/domain/repository"
)

type NewsInteractor struct {
	userRepository repository.UserRepository
}

func NewNewsInteractor(repo repository.UserRepository) *NewsInteractor {
	return &NewsInteractor{repo}
}

// TODO news api
func (i *NewsInteractor) Reply(ctx context.Context, channelId string, uid string) (*model.User, error) {
	u, err := model.NewUser(channelId, uid)
	if err != nil {
		return nil, err
	}
	err = i.userRepository.Store(ctx, u)
	return u, err
}
