package main

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
)

var (
	dialogNodes = []dialog.Node{
		{ID: "start", Text: "Start Node", Keyboard: [][]dialog.Button{{{Text: "Go to node 2", NodeID: "2"}, {Text: "Go to node 3", NodeID: "3"}}, {{Text: "Go Telegram UI", URL: "https://github.com/go-telegram/ui"}}}},
		{ID: "2", Text: "node 2 without keyboard"},
		{ID: "3", Text: "node 3", Keyboard: [][]dialog.Button{{{Text: "Go to start", NodeID: "start"}, {Text: "Go to node 4", NodeID: "4"}}}},
		{ID: "4", Text: "node 4", Keyboard: [][]dialog.Button{{{Text: "Back to 3", NodeID: "3"}}}},
	}
)

func handlerDialog(ctx context.Context, b *bot.Bot, update *models.Update) {
	p := dialog.New(dialogNodes)

	p.Show(ctx, b, update.Message.Chat.ID, "start")
}

func handlerDialogInline(ctx context.Context, b *bot.Bot, update *models.Update) {
	p := dialog.New(dialogNodes, dialog.Inline())

	p.Show(ctx, b, update.Message.Chat.ID, "start")
}
