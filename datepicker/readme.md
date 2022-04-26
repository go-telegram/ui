# Datepicker

![datepicker_1.png](datepicker.png)

## Getting Started

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/datepicker"
)

func main() {
	telegramBotToken := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b := bot.New(telegramBotToken, opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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
```

## Options

See in [options.go](options.go) file 
