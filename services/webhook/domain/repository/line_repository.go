package repository

type LineRepository interface {
	ReplyTextMessages(replyToken string, contents ...string) error
}
