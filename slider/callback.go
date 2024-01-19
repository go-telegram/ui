package slider

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Slider) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
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

			_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    update.CallbackQuery.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.MessageID,
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

			_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    update.CallbackQuery.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.MessageID,
			})
			if errDelete != nil {
				s.onError(errDelete)
			}
		}
		s.onCancel(ctx, b, update.CallbackQuery.Message)
		s.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	}

	slide := s.slides[s.current]

	editParams := &bot.EditMessageMediaParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.MessageID,
		Media: &models.InputMediaPhoto{
			Media:     slide.Photo,
			Caption:   slide.Text,
			ParseMode: models.ParseModeMarkdown,
		},
		ReplyMarkup: s.buildKeyboard(),
	}
	if slide.IsUpload {
		editParams.Media = &models.InputMediaPhoto{
			Media:           "attach://image.png",
			Caption:         slide.Text,
			ParseMode:       models.ParseModeMarkdown,
			MediaAttachment: strings.NewReader(slide.Photo),
		}
	}

	_, errEdit := b.EditMessageMedia(ctx, editParams)
	if errEdit != nil {
		s.onError(errEdit)
	}

	s.callbackAnswer(ctx, b, update.CallbackQuery)
}
