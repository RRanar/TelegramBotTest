package main

import (
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

	if strings.Contains(update.Message.Text, "moroz") {
		var telegramResponseBody, errTelegram = tgbot_core.SendTextToTelegramChat(update.Message.Chat.Id, "Морозило бродяга!Выгони его из хаты!")
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
