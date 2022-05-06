# Datepicker

![datepicker_1.png](datepicker.png)

## Getting Started

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/datepicker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	telegramBotToken := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b := bot.New(telegramBotToken, opts...)

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := datepicker.New(b, onDatepickerSimpleSelect)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select any date",
		ReplyMarkup: kb,
	})
}

func onDatepickerSimpleSelect(ctx context.Context, b *bot.Bot, mes *models.Message, date time.Time) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Chat.ID,
		Text:   "You select " + date.Format("2006-01-02"),
	})
}
```

## Languages

You can define datepicker language by using `Language(lang string)` option. Dy default it is `en`.

Supported languages are defined in [langs.json](langs.json).

You can define your own language by following this example:

```go
langsData := map[string]map[string]string{
    "mylang": map[string]string{
        "Monday":    "Mo",
        "Tuesday":   "Tu",
        "Wednesday": "We",
        "Thursday":  "Th",
        "Friday":    "Fr",
        "Saturday":  "Sa",
        "Sunday":    "Su",
        "January":   "January",
        "February":  "February",
        "March":     "March",
        "April":     "April",
        "May":       "May",
        "June":      "June",
        "July":      "July",
        "August":    "August",
        "September": "September",
        "October":   "October",
        "November":  "November",
        "December":  "December",
        "Cancel":    "Cancel",
        "Prev":      "Prev",
        "Next":      "Next",
        "Back":      "Back"
    },
}

telegramBotToken := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")

opts := []bot.Option{
    bot.Languages(langsData),
    bot.Language("mylang"),
}

b := bot.New(telegramBotToken, opts...)
```

In this example you can see all supported language keys

## Options

See in [options.go](options.go) file 
