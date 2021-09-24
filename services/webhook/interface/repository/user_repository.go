package repository

import (
	"context"
	"time"
	"webhook/domain/model"
	"webhook/domain/repository"
)

type UserRepository struct {
	*DBRepository
}

func NewUserRepository(sql Sql) repository.UserRepository {
	return &UserRepository{&DBRepository{sql: sql, table: "User"}}
}

func (repo *UserRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	var u *model.User
	client, err := repo.sql.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	r := client.Read(ctx, repo.table, []string{id}, u.GetColumns())
	r.Loop(func(row Row) error {
		var id string
		var createdAt time.Time
		var updatedAt time.Time
		var lineBotChannelId string
		var lineUID string
		row.Bind(&id, &createdAt, &updatedAt, &lineBotChannelId, &lineUID)
		u, err = model.NewUser(lineBotChannelId, lineUID)
		return err
	})
	return u, nil
}
