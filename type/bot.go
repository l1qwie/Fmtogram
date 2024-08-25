package typeS

type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatID string `json:"chat_id"`
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatID string `json:"chat_id"`
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatID string `json:"chat_id"`
	UserID int    `json:"user_id"`
}

type BotCommandScope struct {
	BotCommandScopeDefault               *BotCommandScopeDefault
	BotCommandScopeAllPrivateChats       *BotCommandScopeAllPrivateChats
	BotCommandScopeAllGroupChats         *BotCommandScopeAllGroupChats
	BotCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators
	BotCommandScopeChat                  *BotCommandScopeChat
	BotCommandScopeChatAdministrators    *BotCommandScopeChatAdministrators
	BotCommandScopeChatMember            *BotCommandScopeChatMember
}

type BotName struct {
	Name string `json:"name"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}
