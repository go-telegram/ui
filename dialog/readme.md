# Dialog

![dialog.gif](dialog.gif)

Dialog component allows you to create a simple text dialogs. For example, you can use it to display a some menu to the user.

Use `dialog.Inline()` option for change message text, instead of send new messages.

While you create node keyboard, you can to define NodeID or URL.


## Getting Started

```go
package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
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

var (
	dialogNodes = []dialog.Node{
		{ID: "start", Text: "Start Node", Keyboard: [][]dialog.Button{{{Text: "Go to node 2", NodeID: "2"}, {Text: "Go to node 3", NodeID: "3"}}, {{Text: "Go Telegram UI", URL: "https://github.com/go-telegram/ui"}}}},
		{ID: "2", Text: "node 2 without keyboard"},
		{ID: "3", Text: "node 3", Keyboard: [][]dialog.Button{{{Text: "Go to start", NodeID: "start"}, {Text: "Go to node 4", NodeID: "4"}}}},
		{ID: "4", Text: "node 4", Keyboard: [][]dialog.Button{{{Text: "Back to 3", NodeID: "3"}}}},
	}
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	d := dialog.New(dialogNodes)

	d.Show(ctx, b, strconv.Itoa(update.Message.Chat.ID), "start")
}
```



## Options

See in [options.go](options.go) file 
