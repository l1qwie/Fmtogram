package message

import "github.com/l1qwie/Fmtogram/types"

type Message struct {
	ChatID               interface{}            `json:"chat_id,omitempty"`
	BusinessConnectionID string                 `json:"business_connection_id,omitempty"`
	MessageThreadID      int                    `json:"message_thread_id,omitempty"`
	DisableNotification  bool                   `json:"disable_notification,omitempty"`
	ProtectContent       bool                   `json:"protect_content,omitempty"`
	MessageEffectID      string                 `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
}
