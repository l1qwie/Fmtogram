package types

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	BotID string
	Db    *sql.DB
)

const (
	Markdown     string = "Markdown"
	HTML         string = "HTML"
	HttpsRequest string = "https://api.telegram.org/"
)

type BadResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type JustForUpdate struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
}

type StorageOfJson struct {
	ID int `json:"update_id"`
}

type Media struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

type Responser interface {
	RequestOffset(string, *int) error
	Updates(string, *int, *Telegram) error
}

type User struct {
	UserID   int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Language string `json:"language_code"`
	Phone    string `json:"phone_number"`
}

type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Photo struct {
	FileId string `json:"file_id"`
}

type Video struct {
	FileId string `json:"file_id"`
}

type InfMessage struct {
	TypeFrom User    `json:"from"`
	Chat     Chat    `json:"chat"`
	Text     string  `json:"text"`
	Photo    []Photo `json:"photo"`
	Video    []Video `json:"video"`
}

type Message struct {
	MessageId int     `json:"message_id"`
	Chat      Chat    `json:"chat"`
	Photo     []Photo `json:"photo"`
	Video     []Video `json:"video"`
}

type Callback struct {
	TypeFrom User    `json:"from"`
	Data     string  `json:"data"`
	Message  Message `json:"message"`
}

type Update struct {
	UpdateID  int        `json:"update_id"`
	MessageId int        `json:"message_id"`
	Message   InfMessage `json:"message"`
	Query     Callback   `json:"callback_query"`
}

type Returned struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type TelegramError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Telegram struct {
	Ok     bool           `json:"ok"`
	Result []Update       `json:"result"`
	Error  *TelegramError `json:"error,omitempty"`
	AppErr *chan error
}

type Send struct {
	ChatID      interface{} `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup string      `json:"reply_markup"`
	Photo       string      `json:"photo"`
	Video       string      `json:"video"`
	ParseMode   string      `json:"parse_mode"`
	MessageId   int         `json:"message_id"`
	InputMedia  []Media     `json:"media"`
}

type DelMessage struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}
