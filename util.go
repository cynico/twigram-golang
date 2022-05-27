package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func DeepCopy(u *tgbotapi.Update) tgbotapi.Update {
	newUpdate := tgbotapi.Update{
		UpdateID: u.UpdateID,
		Message:  &tgbotapi.Message{},
	}

	return newUpdate
}
