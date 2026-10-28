// Package bot implements the Telegram bot handlers for gbot
package bot

import (
	telebot "gopkg.in/telebot.v4"
)

type command struct {
	name        string
	description string
}

var (
	cmdStart        = command{name: "/start", description: "start the bot"}
	cmdHelp         = command{name: "/help", description: "show this help"}
	cmdVersion      = command{name: "/version", description: "show bot version"}
	cmdConversation = command{name: "/conversation", description: "start a conversation"}
	cmdAbout        = command{name: "/about", description: "about this bot"}
)

// commands is the list shown by /help, in display order
var commands = []command{
	cmdStart,
	cmdHelp,
	cmdVersion,
	cmdConversation,
	cmdAbout,
}

// Register attaches all bot handlers to b
func Register(gbot *telebot.Bot, appVersion string) {
	gbot.Handle(cmdStart.name, handleStart)
	gbot.Handle(cmdHelp.name, handleHelp)
	gbot.Handle(cmdVersion.name, versionHandler(appVersion))
	gbot.Handle(cmdConversation.name, handleConversation)
	gbot.Handle(&btnYes, handleConversationYes)
	gbot.Handle(&btnNo, handleConversationNo)
	gbot.Handle(&btnSecretYes, handleSecretYes)
	gbot.Handle(&btnSecretNo, handleSecretNo)
	gbot.Handle(cmdAbout.name, handleAbout)
	gbot.Handle(telebot.OnText, textHandler(appVersion))
}

// just because "перевірка цього очікує саме в одному файлі"
// "TELE_TOKEN"
// ctx telebot.Context
// TELE_TOKEN
