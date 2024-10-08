package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/progress"
)

var demoProgressSimple *progress.Progress

func initProgressSimple(b *bot.Bot) {
	demoProgressSimple = progress.New(b, progress.WithPrefix("progress-simple"))
}

func handlerProgressSimple(ctx context.Context, b *bot.Bot, update *models.Update) {
	demoProgressSimple.Show(ctx, b, update.Message.Chat.ID)

	go doSomeLongTaskSimple(ctx, b, demoProgressSimple, update.Message.Chat.ID)
}

func doSomeLongTaskSimple(ctx context.Context, b *bot.Bot, p *progress.Progress, chatID any) {
	v := 0.0
	for {
		time.Sleep(time.Second)
		if v == 100 {
			p.Delete(ctx, b)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "Completed",
			})
			return
		}

		v += rand.Float64() * 30
		if v > 100 {
			v = 100
		}
		p.SetValue(ctx, b, v)
	}
}
