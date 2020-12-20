package tgbot_core

import "strings"

type MessageEntity struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

func (entity *MessageEntity) IsCommand() bool {
	if strings.Contains(entity.Type, "bot_command") {
		return true
	}
	return false
}
