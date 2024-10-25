package paginator

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ShowOption func(p *bot.SendMessageParams)

// ShowWithThreadID sets the thread ID for the message
func ShowWithThreadID(threadID int) ShowOption {
	return func(p *bot.SendMessageParams) {
		p.MessageThreadID = threadID
	}
}

// ShowWithReply sets the reply parameters for the message
func ShowWithReply(r *models.ReplyParameters) ShowOption {
	return func(p *bot.SendMessageParams) {
		p.ReplyParameters = r
	}
}
