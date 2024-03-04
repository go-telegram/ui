package datepicker

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	cmdPrevMonth = iota
	cmdNextMonth
	cmdPrevYears
	cmdNextYears
	cmdCancel
	cmdBack
	cmdMonthClick
	cmdYearClick
	cmdNop

	cmdDayClick
	cmdSelectMonth
	cmdSelectYear
)

func (datePicker *DatePicker) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQuery.ID,
	})
	if err != nil {
		datePicker.onError(err)
		return
	}
	if !ok {
		datePicker.onError(fmt.Errorf("callback answer failed"))
	}
}

func (datePicker *DatePicker) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	st := datePicker.decodeState(update.CallbackQuery.Data)

	switch st.cmd {
	case cmdYearClick:
		datePicker.year = st.param
		datePicker.showMain(ctx, b, update.CallbackQuery.Message)
	case cmdMonthClick:
		datePicker.month = time.Month(st.param)
		datePicker.showMain(ctx, b, update.CallbackQuery.Message)
	case cmdDayClick:
		if datePicker.deleteOnSelect {
			_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.Message.ID,
			})
			if errDelete != nil {
				datePicker.onError(fmt.Errorf("failed to delete message onSelect: %w", errDelete))
			}
			b.UnregisterHandler(datePicker.callbackHandlerID)
		}
		datePicker.onSelect(ctx, b, update.CallbackQuery.Message, time.Date(datePicker.year, datePicker.month, st.param, 0, 0, 0, 0, time.Local))
	case cmdCancel:
		if datePicker.deleteOnCancel {
			_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
				ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
				MessageID: update.CallbackQuery.Message.Message.ID,
			})
			if errDelete != nil {
				datePicker.onError(fmt.Errorf("failed to delete message onCancel: %w", errDelete))
			}
			b.UnregisterHandler(datePicker.callbackHandlerID)
		}
		datePicker.onCancel(ctx, b, update.CallbackQuery.Message)
	case cmdPrevYears:
		datePicker.showSelectYear(ctx, b, update.CallbackQuery.Message, st.param)
	case cmdNextYears:
		datePicker.showSelectYear(ctx, b, update.CallbackQuery.Message, st.param)
	case cmdPrevMonth:
		datePicker.month--
		if datePicker.month == 0 {
			datePicker.month = 12
			datePicker.year--
		}
		datePicker.showMain(ctx, b, update.CallbackQuery.Message)
	case cmdNextMonth:
		datePicker.month++
		if datePicker.month > 12 {
			datePicker.month = 1
			datePicker.year++
		}
		datePicker.showMain(ctx, b, update.CallbackQuery.Message)
	case cmdBack:
		datePicker.showMain(ctx, b, update.CallbackQuery.Message)
	case cmdSelectMonth:
		datePicker.showSelectMonth(ctx, b, update.CallbackQuery.Message)
	case cmdSelectYear:
		datePicker.showSelectYear(ctx, b, update.CallbackQuery.Message, datePicker.year)
	case cmdNop:
		// do nothing
	default:
		datePicker.onError(fmt.Errorf("unknown command: %d", st.cmd))
	}

	datePicker.callbackAnswer(ctx, b, update.CallbackQuery)
}

func (datePicker *DatePicker) showSelectMonth(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage) {
	_, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      mes.Message.Chat.ID,
		MessageID:   mes.Message.ID,
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: datePicker.buildMonthKeyboard()},
	})
	if err != nil {
		datePicker.onError(fmt.Errorf("error edit message in showSelectMonth, %w", err))
	}
}

func (datePicker *DatePicker) showSelectYear(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, currentYear int) {
	_, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      mes.Message.Chat.ID,
		MessageID:   mes.Message.ID,
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: datePicker.buildYearKeyboard(currentYear)},
	})
	if err != nil {
		datePicker.onError(fmt.Errorf("error edit message in showSelectYear, %w", err))
	}
}

func (datePicker *DatePicker) showMain(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage) {
	_, err := b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
		ChatID:      mes.Message.Chat.ID,
		MessageID:   mes.Message.ID,
		ReplyMarkup: models.InlineKeyboardMarkup{InlineKeyboard: datePicker.buildKeyboard()},
	})
	if err != nil {
		datePicker.onError(fmt.Errorf("error edit message in showMain, %w", err))
	}
}
