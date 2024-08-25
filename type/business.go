package typeS

type BusinessConnection struct {
	ID         string `json:"id"`
	User       *User  `json:"user"`
	UserChatID int    `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 *Chat  `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                          `json:"time_zone_name"`
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours,omitempty"`
}
