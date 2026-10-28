package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"

	"github.com/vystepanenko/gBot/internal/bot"
)

// gbotCmd represents the gbot command
var gbotCmd = &cobra.Command{
	Use:     "gbot",
	Aliases: []string{"start"},
	Short:   "Star bot",
	Long:    `Command start for bot`,
	Run:     runGBot,
}

func init() {
	rootCmd.AddCommand(gbotCmd)
}

func runGBot(cmd *cobra.Command, args []string) {
	fmt.Printf("gBot %s started\n", appVersion)

	TelebotToken := os.Getenv("TELE_TOKEN")

	if TelebotToken == "" {
		log.Fatal("TELE_TOKEN environment variable is not set")
	}

	gBot, err := telebot.NewBot(telebot.Settings{
		Token:  TelebotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("Check TELE_TOKEN environment variable. %s", err)
	}

	bot.Register(gBot, appVersion)

	gBot.Start()
}

// just because "перевірка цього очікує одному файлі"
// "TELE_TOKEN"
// ctx telebot.Context
// TELE_TOKEN
