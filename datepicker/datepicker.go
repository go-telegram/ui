package datepicker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type DatesMode int

const (
	DateModeExclude DatesMode = iota
	DateModeInclude
)

type OnSelectHandler func(ctx context.Context, bot *bot.Bot, mes models.InaccessibleMessage, date time.Time)
type OnCancelHandler func(ctx context.Context, bot *bot.Bot, mes models.InaccessibleMessage)
type OnErrorHandler func(err error)

type DatePicker struct {
	// configurable params
	startFromSunday bool
	language        string
	langs           LangsData
	deleteOnSelect  bool
	deleteOnCancel  bool
	from            time.Time
	to              time.Time
	dates           []time.Time
	datesMode       DatesMode
	onSelect        OnSelectHandler
	onCancel        OnCancelHandler
	onError         OnErrorHandler

	// current date
	month time.Month
	year  int

	// internal
	prefix            string
	callbackHandlerID string
}

func New(b *bot.Bot, onSelect OnSelectHandler, opts ...Option) *DatePicker {
	year, month, _ := time.Now().Date()

	datePicker := &DatePicker{
		language:       "en",
		langs:          loadLangs(),
		deleteOnSelect: true,
		deleteOnCancel: true,
		onSelect:       onSelect,
		onCancel:       defaultOnCancel,
		onError:        defaultOnError,

		month: month,
		year:  year,

		prefix: bot.RandomString(16),
	}

	for _, opt := range opts {
		opt(datePicker)
	}

	datePicker.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, datePicker.prefix, bot.MatchTypePrefix, datePicker.callback)

	return datePicker
}

func (datePicker *DatePicker) MarshalJSON() ([]byte, error) {
	return json.Marshal(&models.InlineKeyboardMarkup{InlineKeyboard: datePicker.buildKeyboard()})
}

func defaultOnError(err error) {
	log.Printf("[TG-UI-DATEPICKER] [ERROR] %s", err)
}

func defaultOnCancel(_ context.Context, _ *bot.Bot, _ models.InaccessibleMessage) {}
