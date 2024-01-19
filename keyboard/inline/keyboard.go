package inline

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type OnSelect func(ctx context.Context, bot *bot.Bot, mes models.InaccessibleMessage, data []byte)
type OnErrorHandler func(err error)

type handlerData struct {
	Handler OnSelect
	data    []byte
}

type Keyboard struct {
	// configurable
	onError          OnErrorHandler
	deleteAfterClick bool

	// internal
	prefix            string
	handlers          []handlerData
	callbackHandlerID string
	markup            [][]models.InlineKeyboardButton
}

func New(b *bot.Bot, opts ...Option) *Keyboard {
	kb := &Keyboard{
		prefix:           bot.RandomString(16),
		markup:           [][]models.InlineKeyboardButton{{}},
		handlers:         []handlerData{},
		deleteAfterClick: true,
		onError:          defaultOnError,
	}

	for _, opt := range opts {
		opt(kb)
	}

	kb.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, kb.prefix, bot.MatchTypePrefix, kb.callback)

	return kb
}

func (kb *Keyboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(models.InlineKeyboardMarkup{InlineKeyboard: kb.markup})
}

func defaultOnError(err error) {
	log.Printf("[TG-UI-INLINE-KEYBOARD] [ERROR] %s", err)
}
