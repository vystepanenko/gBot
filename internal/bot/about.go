package bot

import (
	telebot "gopkg.in/telebot.v4"
)

func handleAbout(ctx telebot.Context) error {
	return ctx.Send(`This is a test task bot.
Created for: DevOps та Kubernetes. Практичний інтенсив+
Author: Valentyn
All rightst reserved ©`)
}
