package bot

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v4"
)

func textHandler(appVersion string) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		switch strings.ToLower(strings.TrimSpace(ctx.Text())) {
		case "hello", "hi":
			return ctx.Send(fmt.Sprintf("Hello, %s!", ctx.Sender().FirstName))
		case "version":
			return ctx.Send(fmt.Sprintf("Application version: %s", appVersion))
		default:
			return ctx.Send("I don't understand. Type /help for commands.")
		}
	}
}
