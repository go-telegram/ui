package paginator

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type OnErrorHandler func(err error)

const (
	cmdNop   = "nop"
	cmdStart = "start"
	cmdEnd   = "end"
	cmdClose = "close"
)

type Paginator struct {
	data                []string
	perPage             int
	currentPage         int
	pagesCount          int
	separator           string
	prefix              string
	closeButton         string
	onError             OnErrorHandler
	withoutEmptyButtons bool

	callbackHandlerID string
}

func New(b *bot.Bot, data []string, opts ...Option) *Paginator {
	p := &Paginator{
		prefix:      bot.RandomString(16),
		data:        data,
		currentPage: 1,
		separator:   "\n\n",
		perPage:     10,
		onError:     defaultOnError,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.pagesCount = len(data) / p.perPage
	if len(data)%p.perPage != 0 {
		p.pagesCount++
	}

	p.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, p.prefix, bot.MatchTypePrefix, p.callback)

	return p
}

// Prefix returns the prefix of the widget
func (p *Paginator) Prefix() string {
	return p.prefix
}

func defaultOnError(err error) {
	log.Printf("[TG-UI-PAGINATOR] [ERROR] %s", err)
}

func (p *Paginator) Show(ctx context.Context, b *bot.Bot, chatID any, opts ...ShowOption) (*models.Message, error) {
	params := &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        p.buildText(),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: p.buildKeyboard(),
	}

	for _, o := range opts {
		o(params)
	}

	return b.SendMessage(ctx, params)
}
