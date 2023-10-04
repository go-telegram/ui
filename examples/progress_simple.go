package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/progress"
)

func handlerProgressSimple(ctx context.Context, b *bot.Bot, update *models.Update) {
	p := progress.New()

	p.Show(ctx, b, update.Message.Chat.ID)

	go doSomeLongTaskSimple(ctx, b, p, update.Message.Chat.ID)
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
