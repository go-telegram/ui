package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/progress"
)

func handlerProgressCustom(ctx context.Context, b *bot.Bot, update *models.Update) {
	cancelCtx, cancel := context.WithCancel(context.Background())

	progressCancelFunc := func(ctx context.Context, b *bot.Bot, mes *models.Message) {
		cancel()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: mes.Chat.ID,
			Text:   "Progress cancelled",
		})
	}

	renderFunc := func(value float64) string {
		return bot.EscapeMarkdown(fmt.Sprintf("Task progress: %.2f%%", value))
	}

	opts := []progress.Option{
		progress.WithCancel("Cancel", true, progressCancelFunc),
		progress.WithRenderTextFunc(renderFunc),
	}

	p := progress.New(opts...)

	p.Show(ctx, b, update.Message.Chat.ID)

	go doSomeLongTaskCustom(cancelCtx, ctx, b, p, update.Message.Chat.ID)
}

func doSomeLongTaskCustom(cancelCtx, ctx context.Context, b *bot.Bot, p *progress.Progress, chatID any) {
	v := 0.0
	for {
		select {
		case <-cancelCtx.Done():
			return
		default:
		}
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
