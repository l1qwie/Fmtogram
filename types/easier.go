package types

type MessageFmt struct {
	ID   int
	Text string
	Date int
	// ReplyMarkup
	// InlineKeyboard
}

type UserFmt struct {
}

type ChatFmt struct {
}

type ResponseFmt struct {
	Message MessageFmt
	User    UserFmt
	Chat    ChatFmt
}
