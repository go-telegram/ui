package slider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

func (s *Slider) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := methods.SendAnswerCallbackQuery(ctx, b, &methods.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQuery.ID,
	})
	if err != nil {
		s.onError(err)
		return
	}
	if !ok {
		s.onError(fmt.Errorf("callback answer failed"))
	}
}

func (s *Slider) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cmd := strings.TrimPrefix(update.CallbackQuery.Data, s.prefix)

	switch cmd {
	case cmdPrev:
		s.current--
		if s.current < 0 {
			s.current = len(s.slides) - 1
		}
	case cmdNext:
		s.current++
		if s.current >= len(s.slides) {
			s.current = 0
		}
	case cmdSelect:
		if s.deleteOnSelect {
			b.UnregisterHandler(s.callbackHandlerID)

			_, errDelete := methods.DeleteMessage(ctx, b, &methods.DeleteMessageParams{
				ChatID:    strconv.Itoa(update.CallbackQuery.Message.Chat.ID),
				MessageID: update.CallbackQuery.Message.ID,
			})
			if errDelete != nil {
				s.onError(errDelete)
			}
		}
		s.onSelect(ctx, b, update.CallbackQuery.Message, s.current)
		s.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	case cmdCancel:
		if s.deleteOnCancel {
			b.UnregisterHandler(s.callbackHandlerID)

			_, errDelete := methods.DeleteMessage(ctx, b, &methods.DeleteMessageParams{
				ChatID:    strconv.Itoa(update.CallbackQuery.Message.Chat.ID),
				MessageID: update.CallbackQuery.Message.ID,
			})
			if errDelete != nil {
				s.onError(errDelete)
			}
		}
		s.onCancel(ctx, b, update.CallbackQuery.Message)
		s.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	}

	_, errEdit := methods.EditMessageMedia(ctx, b, &methods.EditMessageMediaParams{
		ChatID:    strconv.Itoa(update.CallbackQuery.Message.Chat.ID),
		MessageID: update.CallbackQuery.Message.ID,
		Media: &methods.InputMediaPhoto{
			Type:      "photo",
			Media:     s.slides[s.current].Photo,
			Caption:   s.slides[s.current].Text,
			ParseMode: models.ParseModeMarkdown,
		},
		ReplyMarkup: s.buildKeyboard(),
	})
	if errEdit != nil {
		s.onError(errEdit)
	}

	s.callbackAnswer(ctx, b, update.CallbackQuery)
}
