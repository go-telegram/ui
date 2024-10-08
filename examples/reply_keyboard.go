package main

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/reply"
)

var demoReplyKeyboard *reply.ReplyKeyboard

func initReplyKeyboard(b *bot.Bot) {
	demoReplyKeyboard = reply.New(reply.WithPrefix("reply_keyboard"), reply.IsSelective(), reply.IsOneTimeKeyboard()).
		Button("Button", b, bot.MatchTypeExact, onReplyKeyboardSelect).
		Row().
		Button("Cancel", b, bot.MatchTypeExact, onReplyKeyboardSelect)
}

func handlerReplyKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select example command from reply keyboard:",
		ReplyMarkup: demoReplyKeyboard,
	})
}

func onReplyKeyboardSelect(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "You selected: " + string(update.Message.Text),
	})
}
