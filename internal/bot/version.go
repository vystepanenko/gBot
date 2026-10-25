package bot

import (
	"fmt"

	telebot "gopkg.in/telebot.v4"
)

func versionHandler(appVersion string) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		return ctx.Send(fmt.Sprintf("Application version: %s", appVersion))
	}
}
