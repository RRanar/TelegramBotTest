package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Cannot decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func sendTextToTelegramChat(chatId int, text string) (string, error) {
	log.Printf("Sending %s to chat id:%d", text, chatId)
	var telegramToken, tokenError = getTelegramToken()
	if tokenError != nil {
		log.Printf("Something wrong with token - %s", tokenError.Error())
		return "", tokenError
	}

	var telegramApi string = "https://api.telegram.org/bot" + telegramToken + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})
	if err != nil {
		log.Printf("error when posting text to the chat:%s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error is parsing telegram answer:%s", errRead.Error())
		return "", errRead
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func getTelegramToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv("TELEGRAM_BOT_TOKEN"), nil
}

func HandlerTelegramWebhook(w http.ResponseWriter, r *http.Request) {
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	if strings.Contains(update.Message.Text, "moroz") {
		var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, "Морозило бродяга!Выгони его из хаты!")
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
