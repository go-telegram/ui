package main

import (
	"context"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/datepicker"
)

func handlerDatepickerSimple(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := datepicker.New(b, onDatepickerSimpleSelect)

	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID:      strconv.Itoa(update.Message.Chat.ID),
		Text:        "Select any date",
		ReplyMarkup: kb,
	})
}

func onDatepickerSimpleSelect(ctx context.Context, b *bot.Bot, mes *models.Message, date time.Time) {
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID: strconv.Itoa(mes.Chat.ID),
		Text:   "You select " + date.Format("2006-01-02"),
	})
}
