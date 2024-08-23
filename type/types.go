package typeS

type User struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business"`
	HasMainWebApp           bool   `json:"has_main_web_app"`
}

type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsForum   bool   `json:"is_forum"`
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url"`
	User          User   `json:"user"`
	Language      string `json:"language"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled"`
	Url              string `json:"url"`
	PreferSmallMedia bool   `json:"prefer_small_media"`
	PreferLargeMedia bool   `json:"prefer_large_media"`
	ShowAboveText    bool   `json:"show_above_text"`
}

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	Vcard       string `json:"vcard"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type Story struct {
	Chat Chat `json:"chat"`
	ID   int  `json:"id"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int     `json:"live_period"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximity_alert_radius"`
}

type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities"`
	Position int             `json:"position"`
	IsManual bool            `json:"is_manual"`
}

type Animation struct {
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
}

type Audio struct {
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Performer    string    `json:"performer"`
	Title        string    `json:"title"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
	Thumbnail    PhotoSize `json:"thumbnail"`
}

type Document struct {
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	Thumbnail    PhotoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
}

type Video struct {
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
}

type VideoNote struct {
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	Length       int       `json:"length"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail"`
	FileSize     int       `json:"file_size"`
}

type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    string   `json:"foursquare_id"`
	FoursquareType  string   `json:"foursquare_type"`
	GooglePlaceID   string   `json:"google_place_id"`
	GooglePlaceType string   `json:"google_place_type"`
}

type MessageOrigin struct {
	Type            string `json:"type"`
	Date            int    `json:"date"`
	SenderUser      User   `json:"sender_user"`
	SenderUserName  string `json:"sender_user_name"`
	SenderChat      Chat   `json:"sender_chat"`
	AuthorSignature string `json:"author_signature"`
	Chat            Chat   `json:"chat"`
	MessageID       int    `json:"message_id"`
}

type Sticker struct {
	FileID           string       `json:"file_id"`
	FileUniqueID     string       `json:"file_unique_id"`
	Type             string       `json:"type"`
	Width            int          `json:"width"`
	Height           int          `json:"height"`
	IsAnimated       bool         `json:"is_animated"`
	IsVideo          bool         `json:"is_video"`
	Thumbnail        PhotoSize    `json:"thumbnail"`
	Emoji            string       `json:"emoji"`
	SetName          string       `json:"set_name"`
	PremiumAnimation File         `json:"premium_animation"`
	MaskPosition     MaskPosition `json:"mask_position"`
	CustomEmojiID    string       `json:"custom_emoji_id"`
	NeedsRepainting  bool         `json:"needs_repainting"`
	FileSize         int          `json:"file_size"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEnities       []MessageEntity `json:"question_entities"`
	Options               []PollOptions   `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       int             `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities"`
	OpenPeriod            int             `json:"open_period"`
	CloseDate             int             `json:"close_date"`
}

type ExternalReplyInfo struct {
	Origin             MessageOrigin      `json:"origin"`
	Chat               Chat               `json:"chat"`
	MessageID          int                `json:"message_id"`
	LinkPreviewOptions LinkPreviewOptions `json:"link_preview_options"`
	Animation          Animation          `json:"animation"`
	Audio              Audio              `json:"audio"`
	Document           Document           `json:"document"`
	PaidMedia          PaidMediaInfo      `json:"paid_media"`
	Photo              []PhotoSize        `json:"photo"`
	Sticker            Sticker            `json:"sticker"`
	Story              Story              `json:"story"`
	Video              []Video            `json:"video"`
	VideoNote          VideoNote          `json:"video_note"`
	Voice              Voice              `json:"voice"`
	HasMediaSpoiler    bool               `json:"has_media_spoiler"`
	Contact            Contact            `json:"contact"`
	Dice               Dice               `json:"dice"`
	Game               Game               `json:"game"`
	Giveaway           Giveaway           `json:"giveaway"`
	GiveawayWinners    GiveawayWinners    `json:"giveaway_winners"`
	Invoice            Invoice            `json:"invoice"`
	Location           Location           `json:"location"`
	Poll               Poll               `json:"poll"`
	Venue              Venue              `json:"venue"`
}

type ReplyMessage struct {
	MessageID                     int                           `json:"message_id"`
	MessageThreadID               int                           `json:"message_thread_id"`
	From                          User                          `json:"from"`
	SenderChat                    Chat                          `json:"sender_chat"`
	SenderBoostCount              int                           `json:"sender_boost_count"`
	SenderBusinessBot             User                          `json:"sender_business_bot"`
	Date                          int                           `json:"date"`
	BusinessConnectionID          string                        `json:"business_connection_id"`
	Chat                          Chat                          `json:"chat"`
	ForwardOrigin                 MessageOrigin                 `json:"forward_origin"`
	IsTopicMessage                bool                          `json:"is_topic_message"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`
	ExternalReply                 ExternalReplyInfo             `json:"external_reply"`
	Quote                         TextQuote                     `json:"quote"`
	ReplyToStory                  Story                         `json:"reply_to_story"`
	ViaBot                        User                          `json:"via_bot"`
	EditDate                      int                           `json:"edit_date"`
	HasProtectedContent           bool                          `json:"has_protected_content"`
	IsFrommOffline                bool                          `json:"is_from_offline"`
	MediaGroupID                  string                        `json:"media_group_id"`
	AuthorSignature               string                        `json:"author_signature"`
	Text                          string                        `json:"text"`
	Entities                      []MessageEntity               `json:"entities"`
	LinkPreviewOptions            LinkPreviewOptions            `json:"link_preview_options"`
	EffectID                      string                        `json:"effect_id"`
	Animation                     Animation                     `json:"animation"`
	Audio                         Audio                         `json:"audio"`
	Document                      Document                      `json:"document"`
	PaidMedia                     PaidMediaInfo                 `json:"paid_media"`
	Photo                         []PhotoSize                   `json:"photo"`
	Sticker                       Sticker                       `json:"sticker"`
	Story                         Story                         `json:"story"`
	Video                         Video                         `json:"video"`
	VideoNote                     VideoNote                     `json:"video_note"`
	Voice                         Voice                         `json:"voice"`
	Caption                       string                        `json:"caption"`
	CaptionEntities               []MessageEntity               `json:"caption_entities"`
	ShowCaptionAboveMedia         bool                          `json:"show_caption_above_media"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler"`
	Contact                       Contact                       `json:"contact"`
	Dice                          Dice                          `json:"dice"`
	Game                          Game                          `json:"game"`
	Poll                          Poll                          `json:"poll"`
	Venue                         Venue                         `json:"venue"`
	Location                      Location                      `json:"location"`
	NewChatMembers                []User                        `json:"new_chat_members"`
	LeftChatMember                User                          `json:"left_chat_member"`
	NewChatTitle                  string                        `json:"new_chat_title"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo"`
	GroupChatCreated              bool                          `json:"group_chat_created"`
	SuperGroupChatCreated         bool                          `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int                           `json:"migrate_from_chat_id"`
	PinnedMessage                 MaybeInaccessibleMessage      `json:"pinned_message"`
	Invoice                       Invoice                       `json:"invoice"`
	SuccessfulPayment             SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               RefundedPayment               `json:"refunded_payment"`
	UsersShared                   UsersShared                   `json:"users_shared"`
	ChatShared                    ChatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                        `json:"connected_website"`
	WriteAccessAllowed            WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             ChatBackground                `json:"chat_background_set"`
	ForumTopicCreated             ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      Giveaway                      `json:"giveaway"`
	GiveawayWinners               GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             GiveawayCompleted             `json:"giveaway_completed"`
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup"`
}

type Message struct {
	MessageID                     int                           `json:"message_id"`
	MessageThreadID               int                           `json:"message_thread_id"`
	From                          User                          `json:"from"`
	SenderChat                    Chat                          `json:"sender_chat"`
	SenderBoostCount              int                           `json:"sender_boost_count"`
	SenderBusinessBot             User                          `json:"sender_business_bot"`
	Date                          int                           `json:"date"`
	BusinessConnectionID          string                        `json:"business_connection_id"`
	Chat                          Chat                          `json:"chat"`
	ForwardOrigin                 MessageOrigin                 `json:"forward_origin"`
	IsTopicMessage                bool                          `json:"is_topic_message"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`
	ReplyToMessage                ReplyMessage                  `json:"reply_to_message"` // Message
	ExternalReply                 ExternalReplyInfo             `json:"external_reply"`
	Quote                         TextQuote                     `json:"quote"`
	ReplyToStory                  Story                         `json:"reply_to_story"`
	ViaBot                        User                          `json:"via_bot"`
	EditDate                      int                           `json:"edit_date"`
	HasProtectedContent           bool                          `json:"has_protected_content"`
	IsFrommOffline                bool                          `json:"is_from_offline"`
	MediaGroupID                  string                        `json:"media_group_id"`
	AuthorSignature               string                        `json:"author_signature"`
	Text                          string                        `json:"text"`
	Entities                      []MessageEntity               `json:"entities"`
	LinkPreviewOptions            LinkPreviewOptions            `json:"link_preview_options"`
	EffectID                      string                        `json:"effect_id"`
	Animation                     Animation                     `json:"animation"`
	Audio                         Audio                         `json:"audio"`
	Document                      Document                      `json:"document"`
	PaidMedia                     PaidMediaInfo                 `json:"paid_media"`
	Photo                         []PhotoSize                   `json:"photo"`
	Sticker                       Sticker                       `json:"sticker"`
	Story                         Story                         `json:"story"`
	Video                         []Video                       `json:"video"`
	VideoNote                     VideoNote                     `json:"video_note"`
	Voice                         Voice                         `json:"voice"`
	Caption                       string                        `json:"caption"`
	CaptionEntities               []MessageEntity               `json:"caption_entities"`
	ShowCaptionAboveMedia         bool                          `json:"show_caption_above_media"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler"`
	Contact                       Contact                       `json:"contact"`
	Dice                          Dice                          `json:"dice"`
	Game                          Game                          `json:"game"`
	Poll                          Poll                          `json:"poll"`
	Venue                         Venue                         `json:"venue"`
	Location                      Location                      `json:"location"`
	NewChatMembers                []User                        `json:"new_chat_members"`
	LeftChatMember                User                          `json:"left_chat_member"`
	NewChatTitle                  string                        `json:"new_chat_title"`
	NewChatPhoto                  []PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo"`
	GroupChatCreated              bool                          `json:"group_chat_created"`
	SuperGroupChatCreated         bool                          `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatID               int                           `json:"migrate_from_chat_id"`
	PinnedMessage                 MaybeInaccessibleMessage      `json:"pinned_message"`
	Invoice                       Invoice                       `json:"invoice"`
	SuccessfulPayment             SuccessfulPayment             `json:"successful_payment"`
	RefundedPayment               RefundedPayment               `json:"refunded_payment"`
	UsersShared                   UsersShared                   `json:"users_shared"`
	ChatShared                    ChatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                        `json:"connected_website"`
	WriteAccessAllowed            WriteAccessAllowed            `json:"write_access_allowed"`
	PassportData                  PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	BoostAdded                    ChatBoostAdded                `json:"boost_added"`
	ChatBackgroundSet             ChatBackground                `json:"chat_background_set"`
	ForumTopicCreated             ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              ForumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              ForumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            ForumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	GiveawayCreated               GiveawayCreated               `json:"giveaway_created"`
	Giveaway                      Giveaway                      `json:"giveaway"`
	GiveawayWinners               GiveawayWinners               `json:"giveaway_winners"`
	GiveawayCompleted             GiveawayCompleted             `json:"giveaway_completed"`
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup"`
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
