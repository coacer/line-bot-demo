package repository

import (
	"context"
	"time"
	"webhook/domain/model"
	"webhook/domain/repository"
	"webhook/utils"
)

type UserRepository struct {
	*DBRepository
}

func NewUserRepository(sql Sql) repository.UserRepository {
	return &UserRepository{&DBRepository{sql: sql, table: "User"}}
}

// TODO トランザクション対応
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

// TODO トランザクション対応
func (repo *UserRepository) Store(ctx context.Context, user *model.User) error {
	client, err := repo.sql.NewClient(ctx)
	if err != nil {
		return err
	}
	user.SetId(utils.GenerateUuid())
	user.SetCreatedAt(repo.sql.CommitTimestamp())
	user.SetUpdatedAt(repo.sql.CommitTimestamp())

	q := client.Insert(ctx, repo.table, user.GetColumns(), user.GetValues())
	_, err = client.Commit(ctx, []WriteQuery{q})
	return err
}
