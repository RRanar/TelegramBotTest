package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rranar/telegram-api-webapp/tgbot_core"
)

func HandlerTelegramWebhook(w http.ResponseWriter, r *http.Request) {
	var update, err = tgbot_core.ParseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}
	var telegramResponseBody string
	var errTelegram error

	if update.Message.Entities != nil {
		var keyboard tgbot_core.InlineKeyboardMarkup
		keyboard.InlineKeyboardMarkup = [][]tgbot_core.InlineKeyboardButton{
			[]tgbot_core.InlineKeyboardButton{{Text: "ClickToPunchMoroz", CallbackData: "tg://19117"}, {Text: "FuckTheWorld", CallbackData: "https://t.me/dvachannel"}},
		}
		b, err := json.Marshal(keyboard)
		if err != nil {
			log.Printf("Some error on encode occured:%s", err.Error())
		}
		telegramResponseBody, errTelegram = tgbot_core.SendTextToTelegramChat(update.Message.Chat.Id, "Шалость удалась", b)
	}

	if strings.Contains(update.Message.Text, "moroz") || strings.Contains(update.EditedMessage.Text, "haluka") {
		telegramResponseBody, errTelegram = tgbot_core.SendTextToTelegramChat(update.Message.Chat.Id, "Морозило бродяга!Выгони его из хаты!", nil)
		if errTelegram != nil {
			log.Printf("got error from Telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
		}
	}

	log.Printf("message from chat(%d) - %s", update.Message.Chat.Id, update.Message.Text)
}

func main() {
	http.HandleFunc("/api-telegram", HandlerTelegramWebhook)

	fmt.Printf("Starting web server on :8089\n")
	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}
}
