package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

//go:embed default_text.txt
var defaultMessage string

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	telegramBotToken := os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN")

	opts := []bot.Option{
		bot.WithDebug(),
		bot.WithDefaultHandler(defaultHandler),
		bot.WithMessageTextHandler("/datepicker_simple", bot.MatchTypeExact, handlerDatepickerSimple),
		bot.WithMessageTextHandler("/datepicker_custom", bot.MatchTypeExact, handlerDatepickerCustom),
		bot.WithMessageTextHandler("/inline_keyboard", bot.MatchTypeExact, handlerInlineKeyboard),
		bot.WithMessageTextHandler("/reply_keyboard", bot.MatchTypeExact, handlerReplyKeyboard),
		bot.WithMessageTextHandler("/paginator", bot.MatchTypeExact, handlerPaginator),
		bot.WithMessageTextHandler("/slider", bot.MatchTypeExact, handlerSlider),
		bot.WithMessageTextHandler("/progress_simple", bot.MatchTypeExact, handlerProgressSimple),
		bot.WithMessageTextHandler("/progress_custom", bot.MatchTypeExact, handlerProgressCustom),
		bot.WithMessageTextHandler("/dialog", bot.MatchTypeExact, handlerDialog),
		bot.WithMessageTextHandler("/dialog_inline", bot.MatchTypeExact, handlerDialogInline),
	}

	b, err := bot.New(telegramBotToken, opts...)
	if err != nil {
		panic(err)
	}

	initDatePickerSimple(b)
	initDatePickerCustom(b)
	initInlineKeyboard(b)
	initReplyKeyboard(b)
	initPaginator(b)
	initProgressSimple(b)

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      defaultMessage,
		ParseMode: models.ParseModeMarkdown,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
	})
}
