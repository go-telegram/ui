package paginator

import (
	"strconv"

	"github.com/go-telegram/bot/models"
)

func (p *Paginator) buildKeyboard() models.InlineKeyboardMarkup {
	var row []models.InlineKeyboardButton

	if p.withoutEmptyButtons && p.pagesCount <= 7 {
		if p.pagesCount > 1 {
			for i := 1; i <= p.pagesCount; i++ {
				callbackCommand := strconv.Itoa(i)
				buttonText := strconv.Itoa(i)
				if i == p.currentPage {
					buttonText = "( " + buttonText + " )"
				}

				row = append(row, models.InlineKeyboardButton{Text: buttonText, CallbackData: p.prefix + callbackCommand})
			}
		}
	} else {
		row = append(row, models.InlineKeyboardButton{Text: "\u00AB 1", CallbackData: p.prefix + cmdStart})

		startPage := p.calcStartPage()

		for i := startPage; i < startPage+5; i++ {
			callbackCommand := strconv.Itoa(i)
			buttonText := strconv.Itoa(i)
			if i > p.pagesCount {
				callbackCommand = cmdNop
				buttonText = " "
			}
			if i == p.currentPage {
				buttonText = "( " + buttonText + " )"
			}

			row = append(row, models.InlineKeyboardButton{Text: buttonText, CallbackData: p.prefix + callbackCommand})
		}

		row = append(row, models.InlineKeyboardButton{Text: strconv.Itoa(p.pagesCount) + " \u00BB", CallbackData: p.prefix + cmdEnd})
	}

	kb := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{},
	}

	if len(row) > 0 {
		kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	}

	if p.closeButton != "" {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []models.InlineKeyboardButton{
			{Text: p.closeButton, CallbackData: p.prefix + cmdClose},
		})
	}

	return kb
}

func (p *Paginator) calcStartPage() int {
	if p.pagesCount < 5 { // 5 is pages buttons count
		return 1
	}
	if p.currentPage < 3 { // 3 is center page button
		return 1
	}
	if p.currentPage >= p.pagesCount-2 {
		return p.pagesCount - 4
	}
	return p.currentPage - 2
}
