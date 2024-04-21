package reply

import (
	"encoding/json"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ReplyKeyboard struct {
	// configurable
	inputFieldPlaceholder string
	selective             bool
	oneTimeKeyboard       bool
	resizeKeyboard        bool
	persistent            bool

	// internal
	prefix string
	markup [][]models.KeyboardButton
}

func New(b *bot.Bot, opts ...Option) *ReplyKeyboard {
	kb := &ReplyKeyboard{
		inputFieldPlaceholder: "",
		selective:             false,
		oneTimeKeyboard:       false,
		resizeKeyboard:        false,
		persistent:            false,
		prefix:                bot.RandomString(16),
		markup:                [][]models.KeyboardButton{{}},
	}

	for _, opt := range opts {
		opt(kb)
	}

	return kb
}

// Prefix returns the prefix of the widget
func (kb *ReplyKeyboard) Prefix() string {
	return kb.prefix
}

func (kb *ReplyKeyboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		models.ReplyKeyboardMarkup{
			Keyboard:              kb.markup,
			Selective:             kb.selective,
			OneTimeKeyboard:       kb.oneTimeKeyboard,
			ResizeKeyboard:        kb.resizeKeyboard,
			InputFieldPlaceholder: kb.inputFieldPlaceholder,
			IsPersistent:          kb.persistent,
		},
	)
}
