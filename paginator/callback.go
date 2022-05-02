package paginator

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/methods"
	"github.com/go-telegram/bot/models"
)

func (p *Paginator) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := methods.AnswerCallbackQuery(ctx, b, &methods.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQuery.ID,
	})
	if err != nil {
		p.onError(err)
		return
	}
	if !ok {
		p.onError(fmt.Errorf("callback answer failed"))
	}
}

func (p *Paginator) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cmd := strings.TrimPrefix(update.CallbackQuery.Data, p.prefix)

	switch cmd {
	case cmdNop:
		p.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	case cmdStart:
		if p.currentPage == 1 {
			p.callbackAnswer(ctx, b, update.CallbackQuery)
			return
		}
		p.currentPage = 1
	case cmdEnd:
		if p.currentPage == p.pagesCount {
			p.callbackAnswer(ctx, b, update.CallbackQuery)
			return
		}
		p.currentPage = p.pagesCount
	case cmdClose:
		b.UnregisterHandler(p.callbackHandlerID)

		_, errDelete := methods.DeleteMessage(ctx, b, &methods.DeleteMessageParams{
			ChatID:    strconv.Itoa(update.CallbackQuery.Message.Chat.ID),
			MessageID: update.CallbackQuery.Message.ID,
		})
		if errDelete != nil {
			p.onError(errDelete)
		}
		p.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	default:
		page, _ := strconv.Atoi(cmd)
		p.currentPage = page
	}

	_, errEdit := methods.EditMessageText(ctx, b, &methods.EditMessageTextParams{
		ChatID:          strconv.Itoa(update.CallbackQuery.Message.Chat.ID),
		MessageID:       update.CallbackQuery.Message.ID,
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            p.buildText(),
		ParseMode:       models.ParseModeMarkdown,
		ReplyMarkup:     p.buildKeyboard(),
	})
	if errEdit != nil {
		p.onError(errEdit)
	}

	p.callbackAnswer(ctx, b, update.CallbackQuery)
}
