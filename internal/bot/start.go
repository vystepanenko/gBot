package bot

import (
	telebot "gopkg.in/telebot.v4"
)

func handleStart(ctx telebot.Context) error {
	return ctx.Send("Hi! I'm a test bot for the DevOps course. Type /help for commands.")
}
