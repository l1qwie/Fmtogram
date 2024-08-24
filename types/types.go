package types

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

type Update struct {
	UpdateID               int                         `json:"update_id"`
	Message                Message                     `json:"message"`
	EditedMessage          Message                     `json:"edited_message"`
	ChanPost               Message                     `json:"channel_post"`
	EditedChanPost         Message                     `json:"edited_channel_post"`
	BusinessConn           BusinessConnection          `json:"business_connection"`
	BusinessMessage        Message                     `json:"business_message"`
	EditedBusinessMessage  Message                     `json:"edited_business_message"`
	DeletedBusinessMessage BusinessMessageDeleted      `json:"deleted_business_messages"`
	MessageReaction        MessageReactionUpdated      `json:"message_reaction"`
	MessageReactionCount   MessageReactionCountUpdated `json:"message_reaction_count"`
	InlineQuery            InlineQuery                 `json:"inline_query"`
	ChosenInlineResult     ChosenInlineResult          `json:"chosen_inline_result"`
	CallbackQuery          CallbackQuery               `json:"callback_query"`
	ShippingQuery          ShippingQuery               `json:"shipping_query"`
	PreCheckoutQuery       PreCheckoutQuery            `json:"pre_checkout_query"`
	Poll                   Poll                        `json:"poll"`
	PollAnswer             PollAnswer                  `json:"poll_answer"`
	MyChatMember           ChatMemberUpdated           `json:"my_chat_member"`
	ChatMember             ChatMemberUpdated           `json:"chat_member"`
	ChatJoinRequest        ChatJoinRequest             `json:"chat_join_request"`
	ChatBoost              ChatBoostUpdated            `json:"chat_boost"`
	RemovedChatBoost       ChatBoostRemoved            `json:"removed_chat_boost"`
}
