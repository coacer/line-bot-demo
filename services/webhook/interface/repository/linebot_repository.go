package repository

import (
	"webhook/domain/repository"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineRepository struct {
	bot *linebot.Client
}

func NewLineRepository(channelSecret string, channelToken string) repository.LineRepository {
	lbc, _ := linebot.New(channelSecret, channelToken)
	return &LineRepository{lbc}
}

func (r *LineRepository) ReplyTextMessages(replyToken string, contents ...string) error {
	messages := make([]*linebot.TextMessage, len(contents))
	for i, c := range contents {
		messages[i] = linebot.NewTextMessage(c)
	}
	_, err := r.bot.ReplyMessage(replyToken, messages[0]).Do()
	return err
}
