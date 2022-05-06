# Go Telegram Bot UI

> The project is under development. API may be changed before v1.0.0 version.

> [Telegram Group](https://t.me/gotelegrambotui)

UI controls for telegram bot [go-telegram/bot](https://github.com/go-telegram/bot)

- datepicker
- inline keyboard
- paginator
- slider
- timepicker (todo)

Feel to free to contribute and issues!

## Getting Started

```bash
go get github.com/go-telegram-bot/bot
go get github.com/go-telegram-bot/ui
```

**Important**

UI components register own bot handlers on init. 
If you restart the bot instance, inline buttons in already opened components can't work.

For example, you can handle CallbackQuery in the default handler and send a message to the user.

```go

func defaultHandler(ctx context.Context, b *bot.Bot, update bot.Update)  {
	if update.CallbackQuery != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.CallbackQuery.Message.Chat.ID,
            Text:      "Bot was updated. Try to call calendar again",
        })  
	}
}
```

### Demo bot

You can run demo bot from `examples` folder.

Before start, you should set `EXAMPLE_TELEGRAM_BOT_TOKEN` environment variable to your bot token.

Also, you can try [online version of this bot](https://t.me/gotelegramuidemobot) 

## DatePicker

![datepicker_1.png](datepicker/datepicker.png)

- custom localizations
- set from/to dates
- define include/exclude dates

[Documentation](datepicker/readme.md)

## Inline Keyboard

![inline_keyboard.png](keyboard/inline/inline_keyboard.png)

Small helper for easy building of inline keyboard.

[Documentation](keyboard/inline/readme.md)

## Paginator

![paginator.png](paginator/paginator.png)

- pass any slice of strings
- set perPage value
- set custom lines separator

[Documentation](paginator/readme.md)

## Slider

![slider.png](slider/slider.png)

- pass slides with images and text

[Documentation](slider/readme.md)

## Progress

![progress.png](progress/progress.png)

Progress bar for long tasks

[Documentation](progress/readme.md)
