package tgbot_core

type Update struct {
	UpdateId      int           `json:"update_id"`
	Message       Message       `json:"message"`
	EditedMessage Message       `json:"edited_message"`
	CallbackQuery CallbackQuery `json:callback_query`
}
