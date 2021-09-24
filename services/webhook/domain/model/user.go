package model

import (
	"strconv"
	"webhook/exception"
	"webhook/utils"
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
	return utils.GetKeys(*u)
}

// LINEの仕様準拠
func validChannelId(id string) bool {
	_, err := strconv.Atoi(id)
	return err != nil && len(id) != 10
}

// LINEの仕様準拠
func validUID(uid string) bool {
	return uid[0] == 'U' && len(uid) == 33
}
