package must

type Check interface {
	RequiredFields(string, []byte) error
}

type MediaCheck struct {
	Media  []string    `json:"media"`
	ChatID interface{} `json:"chat_id"`
}

type PhotoCheck struct {
	Photo  string      `json:"photo"`
	ChatID interface{} `json:"chat_id"`
}

type AudioCheck struct {
	Audio  string      `json:"audio"`
	ChatID interface{} `json:"chat_id"`
}

type DocumentCheck struct {
	Document string      `json:"document"`
	ChatID   interface{} `json:"chat_id"`
}

type VideoCheck struct {
	Video  string      `json:"video"`
	ChatID interface{} `json:"chat_id"`
}

type MessageCheck struct {
	Text   string      `json:"text"`
	ChatID interface{} `json:"chat_id"`
}

type ForwardMessageCheck struct {
	ChatID     interface{} `json:"chat_id"`
	FromChatID interface{} `json:"from_chat_id"`
}

type ForwardMessages struct {
	ChatID     interface{} `json:"chat_id"`
	FromChatID interface{} `json:"from_chat_id"`
	MessageIDs []int       `json:"message_ids"`
}
