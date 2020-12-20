package tgbot_core

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboardMarkup [][]InlineKeyboardButton `json:"inline_keyboard"`
}
