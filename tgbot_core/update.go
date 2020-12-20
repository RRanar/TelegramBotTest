package tgbot_core

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}
