package typestt

var (
	BotID string
	DEBUG bool
	INFO  bool
	// StartFunc func(*Telegram, *Returned) *formatter.Formatter
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

type StorageOfJson struct {
	ID int `json:"update_id"`
}

type JustForUpdate struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
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

type ReturnedMessage struct {
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

type Returned struct {
	Ok     bool            `json:"ok"`
	Result ReturnedMessage `json:"result"`
}

type TelegramError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Telegram struct {
	Ok     bool           `json:"ok"`
	Result []interface{}  `json:"result"`
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
	ChatId    interface{} `json:"chat_id"`
	MessageId interface{} `json:"message_id"`
}

type Message struct {
	MessageID       int  `json:"message_id"`
	MessageThreadID int  `json:"message_thread_id"`
	From            User `json:"from"`
	SenderChat      Chat `json:""`

	Chat  Chat    `json:"chat"`
	Text  string  `json:"text"`
	Photo []Photo `json:"photo"`
	Video []Video `json:"video"`
}
