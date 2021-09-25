package model

import (
	"strconv"
	"webhook/exception"
)

type User struct {
	*DbBaseModel
	lineBotChannelId string `json:"lineBotChannelId"`
	lineUID          string `json:"lineUID"`
}

func NewUser(lineBotChannelId, lineUID string) (*User, error) {
	if !validChannelId(lineBotChannelId) || !validUID(lineUID) {
		return nil, exception.NewError(exception.UserModelInvalidError, nil)
	}

	return &User{
		&DbBaseModel{},
		lineBotChannelId,
		lineUID,
	}, nil
}

func (u *User) GetColumns() []string {
	return []string{"id", "createdAt", "updatedAt", "lineBotChannelId", "lineUID"}
}

func (u *User) GetValues() []interface{} {
	return []interface{}{
		u.id,
		u.createdAt,
		u.updatedAt,
		u.lineBotChannelId,
		u.lineUID,
	}
}

// LINEの仕様準拠
func validChannelId(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil && len(id) == 10
}

// LINEの仕様準拠
func validUID(uid string) bool {
	return uid[0] == 'U' && len(uid) == 33
}
