package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

// TelebotToken token
var TelebotToken = os.Getenv("TELEBOT_TOKEN")

// gbotCmd represents the gbot command
var gbotCmd = &cobra.Command{
	Use:     "gbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("qBot %s started", appVersion)

		gBot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TelebotToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Check TELEBOT_TOKEN environment variable. %s", err)
		}

		gBot.Handle(telebot.OnText, func(ctx telebot.Context) error {
			payload := ctx.Message().Payload

			switch payload {
			case "version":
				err = ctx.Send(fmt.Sprintf("Application version: %s", appVersion))
			}
			return err
		})

		gBot.Start()
	},
}

func init() {
	rootCmd.AddCommand(gbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
