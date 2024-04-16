# Progress

![progress.png](progress.png)

If you want to use Cancel button, you should to use option `WithCancel(buttonText string, deleteOnCancel bool, onCancel OnCancelFunc)`.

You can to customize the progress text with option `WithRenderTextFunc(f RenderTextFunc)`

Render func receives the progress value and returns the text.

```
func MyRenderFunc(value float64) string {
    s := fmt.Sprintf("My Progress: %.2f%%", value)
    
    return bot.EscapedMarkdown(s)
}
```

Progress use Markdown for rendering. If you have any Markdown syntax, you should escape your text, if needed.
In this example we have dot in the text. So, we need to escape it.

## Getting Started

```go
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/progress"
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
	p := progress.New()

	p.Show(ctx, b, update.Message.Chat.ID)

	go doSomeLongTaskSimple(ctx, b, p, update.Message.Chat.ID)
}

func doSomeLongTaskSimple(ctx context.Context, b *bot.Bot, p *progress.Progress, chatID int) {
	v := 0.0
	for {
		time.Sleep(time.Second)
		if v == 100 {
			p.Delete(ctx, b)
			p.Done(ctx, b)
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
```

## Options

See in [options.go](options.go) file 
