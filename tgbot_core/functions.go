package tgbot_core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ParseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Cannot decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func SendTextToTelegramChat(chatId int, text string) (string, error) {
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
	err := godotenv.Load("../.env")
	if err != nil {
		return "", err
	}
	return os.Getenv("TELEGRAM_BOT_TOKEN"), nil
}
