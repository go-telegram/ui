package reply

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (kb *ReplyKeyboard) Row() *ReplyKeyboard {
	if len(kb.markup[len(kb.markup)-1]) > 0 {
		kb.markup = append(kb.markup, []models.KeyboardButton{})
	}
	return kb
}

func (kb *ReplyKeyboard) Button(text string, b *bot.Bot, matchType bot.MatchType, handler bot.HandlerFunc) *ReplyKeyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.KeyboardButton{
		Text: text,
	})

	if handler != nil {
		b.RegisterHandler(bot.HandlerTypeMessageText, text, matchType, handler)
	}

	return kb
}
