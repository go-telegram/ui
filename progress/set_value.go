package progress

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

func (p *Progress) SetValue(ctx context.Context, b *bot.Bot, value float64) {
	if p.canceled {
		return
	}

	p.value = value

	editParams := &methods.EditMessageTextParams{
		ChatID:    p.message.Chat.ID,
		MessageID: p.message.ID,
		Text:      p.renderTextFunc(p.value),
		ParseMode: models.ParseModeMarkdown,
	}
	if p.onCancel != nil {
		editParams.ReplyMarkup = p.buildKeyboard()
	}

	m, err := methods.EditMessageText(ctx, b, editParams)
	if err != nil {
		p.onError(err)
	}

	p.message = m
}
