package datepicker

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-telegram/bot/models"
)

func (datePicker *DatePicker) buildYearKeyboard(currentYear int) [][]models.InlineKeyboardButton {
	// Years keyboard is table 5*5 where current year is in the middle. 25 buttons total, 12 before and 12 after.
	var keyboard [][]models.InlineKeyboardButton

	var topRow []models.InlineKeyboardButton

	if datePicker.from.IsZero() || datePicker.from.Year() < currentYear-12 {
		topRow = append(topRow, models.InlineKeyboardButton{Text: "\u2190 " + datePicker.lang("Prev"), CallbackData: datePicker.encodeState(state{cmd: cmdPrevYears, param: currentYear - 25})})
	}
	if datePicker.to.IsZero() || datePicker.to.Year() > currentYear+12 {
		topRow = append(topRow, models.InlineKeyboardButton{Text: datePicker.lang("Next") + " \u2192", CallbackData: datePicker.encodeState(state{cmd: cmdNextYears, param: currentYear + 25})})
	}

	if len(topRow) > 0 {
		keyboard = append(keyboard, topRow)
	}

	var row []models.InlineKeyboardButton
	idx := 1
	for i := currentYear - 12; i <= currentYear+12; i++ {
		if idx > 5 {
			idx = 1
			keyboard = append(keyboard, row)
			row = []models.InlineKeyboardButton{}
		}

		yearCmd := cmdYearClick
		yearText := strconv.Itoa(i)

		if !datePicker.from.IsZero() && i < datePicker.from.Year() {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			idx++
			continue
		}
		if !datePicker.to.IsZero() && i > datePicker.to.Year() {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			idx++
			continue
		}

		row = append(row, models.InlineKeyboardButton{Text: yearText, CallbackData: datePicker.encodeState(state{cmd: yearCmd, param: i})})
		idx++
	}
	keyboard = append(keyboard, row)

	keyboard = append(keyboard, []models.InlineKeyboardButton{
		{Text: datePicker.lang("Back"), CallbackData: datePicker.encodeState(state{cmd: cmdBack})},
	})

	return keyboard
}

func (datePicker *DatePicker) buildMonthKeyboard() [][]models.InlineKeyboardButton {
	var keyboard [][]models.InlineKeyboardButton

	var row []models.InlineKeyboardButton
	for i := 0; i < 12; i++ {
		if i > 0 && i%4 == 0 { //4 months per row
			keyboard = append(keyboard, row)
			row = []models.InlineKeyboardButton{}
		}

		startOfMonth := time.Date(datePicker.year, time.Month(i+1), 1, 0, 0, 0, 0, time.Local)

		if !datePicker.from.IsZero() && startOfMonth.Before(datePicker.from.AddDate(0, -1, 0)) {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			continue
		}

		if !datePicker.to.IsZero() && startOfMonth.After(datePicker.to) {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			continue
		}

		row = append(row, models.InlineKeyboardButton{Text: datePicker.lang(time.Month(i + 1).String()), CallbackData: datePicker.encodeState(state{cmd: cmdMonthClick, param: i + 1})})
	}

	keyboard = append(keyboard, row)

	keyboard = append(keyboard, []models.InlineKeyboardButton{{Text: datePicker.lang("Back"), CallbackData: datePicker.encodeState(state{cmd: cmdBack})}})

	return keyboard
}

func (datePicker *DatePicker) buildKeyboard() [][]models.InlineKeyboardButton {
	var data [][]models.InlineKeyboardButton

	top := []models.InlineKeyboardButton{
		{Text: datePicker.lang(datePicker.month.String()), CallbackData: datePicker.encodeState(state{cmd: cmdSelectMonth})},
		{Text: strconv.Itoa(datePicker.year), CallbackData: datePicker.encodeState(state{cmd: cmdSelectYear})},
	}

	data = append(data, top)

	dow := []models.InlineKeyboardButton{
		{Text: datePicker.lang("Monday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Tuesday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Wednesday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Thursday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Friday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Saturday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
		{Text: datePicker.lang("Sunday"), CallbackData: datePicker.encodeState(state{cmd: cmdNop})},
	}

	if datePicker.startFromSunday {
		dow = append([]models.InlineKeyboardButton{dow[len(dow)-1]}, dow[:len(dow)-1]...)
	}

	data = append(data, dow)

	startOfMonth := time.Date(datePicker.year, datePicker.month, 1, 0, 0, 0, 0, time.Local)

	var row []models.InlineKeyboardButton

	// space days before first of month
	skipFirst := int(startOfMonth.Weekday()) // 0 - Sunday, 1 - Monday, ...
	if !datePicker.startFromSunday {
		if skipFirst == 0 {
			skipFirst = 6
		} else {
			skipFirst--
		}
	}

	idx := skipFirst + 1

	for i := 0; i < skipFirst; i++ {
		row = append(row, models.InlineKeyboardButton{Text: " ", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
	}

	// days of month
	for i := 1; i <= startOfMonth.AddDate(0, 1, -1).Day(); i++ {
		if idx == 8 {
			idx = 1
			data = append(data, row)
			row = []models.InlineKeyboardButton{}
		}

		now := time.Date(datePicker.year, datePicker.month, i, 0, 0, 0, 0, time.Local)

		if !datePicker.from.IsZero() && now.Before(datePicker.from) {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			idx++
			continue
		}

		if !datePicker.to.IsZero() && now.After(datePicker.to) {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			idx++
			continue
		}

		var isInDates bool

		for _, d := range datePicker.dates {
			isInDates = d.Day() == i && d.Month() == datePicker.month && d.Year() == datePicker.year
			if isInDates {
				break
			}
		}

		if len(datePicker.dates) > 0 && ((isInDates && datePicker.datesMode == DateModeExclude) || (!isInDates && datePicker.datesMode == DateModeInclude)) {
			row = append(row, models.InlineKeyboardButton{Text: "-", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
			idx++
			continue
		}

		row = append(row, models.InlineKeyboardButton{Text: strconv.Itoa(i), CallbackData: datePicker.encodeState(state{cmd: cmdDayClick, param: i})})
		idx++
	}

	// space days after last of month
	if idx < 8 {
		for i := 0; i < 8-idx; i++ {
			row = append(row, models.InlineKeyboardButton{Text: " ", CallbackData: datePicker.encodeState(state{cmd: cmdNop})})
		}
	}

	data = append(data, row)

	var bottom []models.InlineKeyboardButton

	if datePicker.from.IsZero() || datePicker.from.Before(startOfMonth) {
		prevMonth := startOfMonth.AddDate(0, -1, 0)
		bottom = append(bottom, models.InlineKeyboardButton{Text: fmt.Sprintf("\u2190 %s %d", datePicker.lang(prevMonth.Month().String()), prevMonth.Year()), CallbackData: datePicker.encodeState(state{cmd: cmdPrevMonth})})
	}

	bottom = append(bottom, models.InlineKeyboardButton{Text: datePicker.lang("Cancel"), CallbackData: datePicker.encodeState(state{cmd: cmdCancel})})

	if datePicker.to.IsZero() || datePicker.to.After(startOfMonth.AddDate(0, 1, -1)) {
		nextMonth := startOfMonth.AddDate(0, 1, 0)
		bottom = append(bottom, models.InlineKeyboardButton{Text: fmt.Sprintf("%s %d \u2192", datePicker.lang(nextMonth.Month().String()), nextMonth.Year()), CallbackData: datePicker.encodeState(state{cmd: cmdNextMonth})})
	}

	data = append(data, bottom)

	return data
}
