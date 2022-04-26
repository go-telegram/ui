package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

//go:embed default_text.txt
var defaultMessage string

func main() {
	telegramBotToken := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")

	opts := []bot.Option{
		//bot.WithDebug(),
		bot.WithDefaultHandler(defaultHandler),
		bot.WithMessageTextHandler("/datepicker_simple", bot.MatchTypeExact, handlerDatepickerSimple),
		bot.WithMessageTextHandler("/datepicker_custom", bot.MatchTypeExact, handlerDatepickerCustom),
		bot.WithMessageTextHandler("/inline_keyboard", bot.MatchTypeExact, handlerInlineKeyboard),
		bot.WithMessageTextHandler("/paginator", bot.MatchTypeExact, handlerPaginator),
		bot.WithMessageTextHandler("/slider", bot.MatchTypeExact, handlerSlider),
	}

	b := bot.New(telegramBotToken, opts...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	methods.SendMessage(ctx, b, &methods.SendMessageParams{
		ChatID:    strconv.Itoa(update.Message.Chat.ID),
		Text:      defaultMessage,
		ParseMode: models.ParseModeMarkdown,
	})
}
