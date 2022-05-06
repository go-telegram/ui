package main

import (
	"context"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/datepicker"
)

func makeTime(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func handlerDatepickerCustom(ctx context.Context, b *bot.Bot, update *models.Update) {
	excludeDays := []time.Time{
		makeTime(2020, 1, 10),
		makeTime(2020, 1, 13),
		makeTime(2019, 12, 27),
		makeTime(2019, 12, 28),
		makeTime(2019, 12, 29),
	}

	dpOpts := []datepicker.Option{
		datepicker.StartFromSunday(),
		datepicker.CurrentDate(makeTime(2020, 1, 15)),
		datepicker.From(makeTime(2019, 12, 15)),
		datepicker.To(makeTime(2020, 1, 25)),
		datepicker.OnCancel(onDatepickerCustomCancel),
		datepicker.Language("ru"),
		datepicker.Dates(datepicker.DateModeExclude, excludeDays),
	}

	kb := datepicker.New(b, onDatepickerCustomSelect, dpOpts...)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select the date",
		ReplyMarkup: kb,
	})
}

func onDatepickerCustomCancel(ctx context.Context, b *bot.Bot, mes *models.Message) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Chat.ID,
		Text:   "Canceled",
	})
}

func onDatepickerCustomSelect(ctx context.Context, b *bot.Bot, mes *models.Message, date time.Time) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Chat.ID,
		Text:   "You select " + date.Format("2006-01-02"),
	})
}
