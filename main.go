package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	. "twigram-go/data"

	database "twigram-go/db/mysql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hashicorp/go-hclog"
)

var (
	msg    tgbotapi.MessageConfig
	update tgbotapi.Update
)

type test struct {
	http.ResponseWriter
	db sql.DB
}

func initDB() {

}

func main() {

	// Creating a logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "twigram",
		Level: hclog.LevelFromString("DEBUG"),
	})

	// Initializing telegram bot.
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		logger.Error("error initializing telegram bot", "error", err)
		os.Exit(1)
	}

	// Setting whether or not to use debug mode in telegram bot.
	bot.Debug, err = strconv.ParseBool(os.Getenv("TELEGRAM_DEBUG"))
	if err != nil {
		bot.Debug = false
	}

	// Initializing the database.
	err = database.InitDB()
	if err != nil {
		logger.Error("error initializing database", "error", err)
		os.Exit(1)
	}

	// err = database.Migrate()
	// if err != nil {
	// 	logger.Error("error migrating database", "error", err)
	// 	os.Exit(1)
	// }

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	SendMessage := func(text string) (tgbotapi.Message, error) {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, HelpMessage)
		msg.ReplyToMessageID = update.Message.MessageID
		return bot.Send(msg)
	}

	for update = range updates {

		local_update := &update

		go func(u *tgbotapi.Update) {

			if !u.Message.IsCommand() {
				fmt.Printf("%v", *u.MyChatMember)
			}

			switch u.Message.Command() {
			case "start":

				// Generate code_challenge specific to the user.
				// Generate the url to have permission.
			case "help":
				SendMessage(HelpMessage)
			case "follow":
				SendMessage(FollowMessage)
				// Extra code
			case "unfollow":
				SendMessage(UnfollowMessage)
			case "list":

			default:
				SendMessage(UnrecognizableCommand)

			}
		}(local_update)
	}

}
