package bot

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v4"
)

func handleHelp(ctx telebot.Context) error {
	var data strings.Builder
	for _, c := range commands {
		fmt.Fprintf(&data, "%s - %s\n", c.name, c.description)
	}
	return ctx.Send(data.String())
}
