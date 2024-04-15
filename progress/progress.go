package progress

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type OnCancelFunc func(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage)
type RenderTextFunc func(value float64) string
type OnErrorHandler func(err error)

type Progress struct {
	prefix         string
	value          float64
	renderTextFunc RenderTextFunc
	onError        OnErrorHandler

	message *models.Message

	cancelText        string
	onCancel          OnCancelFunc
	onCancelHandlerId string
	deleteOnCancel    bool
	canceled          bool
}

func New(opts ...Option) *Progress {
	p := &Progress{
		prefix:         bot.RandomString(16),
		value:          0,
		renderTextFunc: defaultRenderTextFunc,
		onError:        defaultOnError,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// Prefix returns the prefix of the widget
func (p *Progress) Prefix() string {
	return p.prefix
}

func (p *Progress) Show(ctx context.Context, b *bot.Bot, chatID any) error {
	sendParams := &bot.SendMessageParams{
		ChatID:    chatID,
		Text:      p.renderTextFunc(p.value),
		ParseMode: models.ParseModeMarkdown,
	}

	if p.onCancel != nil {
		sendParams.ReplyMarkup = p.buildKeyboard()
	}

	m, err := b.SendMessage(ctx, sendParams)

	if err != nil {
		return err
	}

	if p.onCancel != nil {
		p.onCancelHandlerId = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, p.prefix,
			bot.MatchTypeExact, p.onCancelCall)
	}

	p.message = m

	return nil
}

func (p *Progress) Delete(ctx context.Context, b *bot.Bot) {
	_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    p.message.Chat.ID,
		MessageID: p.message.ID,
	})
	if err != nil {
		p.onError(err)
	}
}

// ctx is being passed for forward compatibility reasons
func (p *Progress) Done(ctx context.Context, b *bot.Bot) {
	if p.onCancelHandlerId != "" {
		b.UnregisterHandler(p.onCancelHandlerId)
		p.onCancelHandlerId = ""
	}
}

func defaultRenderTextFunc(value float64) string {
	return bot.EscapeMarkdown(fmt.Sprintf("%.2f%%", value))
}

func defaultOnError(err error) {
	log.Printf("[TG-UI-PROGRESS] [ERROR] %s", err)
}

func (p *Progress) buildKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: p.cancelText, CallbackData: p.prefix},
			},
		},
	}
}

func (p *Progress) onCancelCall(ctx context.Context, b *bot.Bot, update *models.Update) {
	p.canceled = true
	if p.deleteOnCancel {
		_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
		})
		if err != nil {
			p.onError(err)
		}
	}
	p.onCancel(ctx, b, update.CallbackQuery.Message)
	p.Done(ctx, b)
}
