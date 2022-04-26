package slider

import (
	"context"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
	"log"

	"github.com/go-telegram/bot"
)

type OnSelectFunc func(ctx context.Context, b *bot.Bot, message *models.Message, item int)
type OnCancelFunc func(ctx context.Context, b *bot.Bot, message *models.Message)
type OnErrorFunc func(err error)

type Slide struct {
	Photo string
	Text  string
}

var (
	cmdPrev   = "prev"
	cmdNext   = "next"
	cmdNop    = "nop"
	cmdSelect = "select"
	cmdCancel = "cancel"
)

type Slider struct {
	prefix string
	slides []Slide

	selectButtonText string
	onSelect         OnSelectFunc
	cancelButtonText string
	onCancel         OnCancelFunc
	onError          OnErrorFunc

	deleteOnSelect bool
	deleteOnCancel bool

	current           int
	callbackHandlerID string
}

func New(slides []Slide, opts ...Option) *Slider {
	s := &Slider{
		prefix:  bot.RandomString(16),
		slides:  slides,
		onError: defaultOnError,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func defaultOnError(err error) {
	log.Printf("[TG-UI-SLIDER] [ERROR] %s", err)
}

func (s *Slider) Show(ctx context.Context, b *bot.Bot, chatID string) (*models.Message, error) {
	s.callbackHandlerID, _ = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, s.prefix, bot.MatchTypePrefix, s.callback)

	slide := s.slides[s.current]

	return methods.SendPhoto(ctx, b, &methods.SendPhotoParams{
		ChatID:      chatID,
		Photo:       &models.InputFileString{Data: slide.Photo},
		Caption:     slide.Text,
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: s.buildKeyboard(),
	})
}
