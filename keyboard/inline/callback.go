package inline

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (kb *Keyboard) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQuery.ID,
	})
	if err != nil {
		kb.onError(err)
		return
	}
	if !ok {
		kb.onError(fmt.Errorf("callback answer failed"))
	}
}

func (kb *Keyboard) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	if kb.deleteAfterClick {
		b.UnregisterHandler(kb.callbackHandlerID)

		_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.MessageID,
		})
		if errDelete != nil {
			kb.onError(fmt.Errorf("error delete message in callback, %w", errDelete))
		}
	}

	btnNum, errBtnNum := strconv.Atoi(strings.TrimPrefix(update.CallbackQuery.Data, kb.prefix))
	if errBtnNum != nil {
		kb.onError(fmt.Errorf("wrong callback data btnNum, %s", update.CallbackQuery.Data))
		kb.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	}

	if len(kb.handlers) <= btnNum {
		kb.onError(fmt.Errorf("wrong callback data, %s", update.CallbackQuery.Data))
		kb.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	}

	if kb.handlers[btnNum].Handler != nil {
		kb.handlers[btnNum].Handler(ctx, b, update.CallbackQuery.Message, kb.handlers[btnNum].data)
	}

	kb.callbackAnswer(ctx, b, update.CallbackQuery)
}
