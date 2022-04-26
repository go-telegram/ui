package main

import (
	"context"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

func handlerInlineKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := inline.New(b).
		Row().
		Button("Row 1, Btn 1", []byte("1-1"), onInlineKeyboardSelect).
		Button("Row 1, Btn 2", []byte("1-2"), onInlineKeyboardSelect).
		Row().
		Button("Row 2, Btn 1", []byte("2-1"), onInlineKeyboardSelect).
		Button("Row 2, Btn 2", []byte("2-2"), onInlineKeyboardSelect).
		Button("Row 2, Btn 3", []byte("2-3"), onInlineKeyboardSelect).
		Row().
		Button("Row 3, Btn 1", []byte("3-1"), onInlineKeyboardSelect).
		Button("Row 3, Btn 2", []byte("3-2"), onInlineKeyboardSelect).
		Button("Row 3, Btn 3", []byte("3-3"), onInlineKeyboardSelect).
		Button("Row 3, Btn 4", []byte("3-4"), onInlineKeyboardSelect).
		Row().
		Button("Cancel", []byte("cancel"), onInlineKeyboardSelect)

	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID:      strconv.Itoa(update.Message.Chat.ID),
		Text:        "Select the variant",
		ReplyMarkup: kb,
	})
}

func onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes *models.Message, data []byte) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: strconv.Itoa(mes.Chat.ID),
		Text:   "You selected: " + string(data),
	})
}
