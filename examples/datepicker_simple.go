package main

import (
	"context"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/datepicker"
)

var demoDatePickerSimple *datepicker.DatePicker

func initDatePickerSimple(b *bot.Bot) {
	demoDatePickerSimple = datepicker.New(b, onDatepickerSimpleSelect, datepicker.WithPrefix("datepicker-simple"))
}

func handlerDatepickerSimple(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select any date",
		ReplyMarkup: demoDatePickerSimple,
	})
}

func onDatepickerSimpleSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, date time.Time) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Message.Chat.ID,
		Text:   "You select " + date.Format("2006-01-02"),
	})
}
