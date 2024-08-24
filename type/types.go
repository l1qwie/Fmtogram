package typeS

type ForumTopicClosed struct {
	// Currently holds no information
}

type ForumTopicReopened struct {
	// Currently holds no information
}

type GeneralForumTopicHidden struct {
	// Currently holds no information
}

type GeneralForumTopicUnhidden struct {
	// Currently holds no information
}

type GiveawayCreated struct {
	// Currently holds no information
}

type VideoChatStarted struct {
	// This object represents a service message about a video chat started in the chat. Currently holds no information.
}

type CallbackGame struct {
	// A placeholder, currently holds no information
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

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

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request"`
	WebApName          string `json:"web_app_name"`
	FromAttachmentMenu bool   `json:"from_attachment_menu"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

type ForumTopicEdited struct {
	Name              string `json:"name"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query"`
	AllowUserChats    bool   `json:"allow_user_chats"`
	AllowBotChats     bool   `json:"allow_bot_chats"`
	AllowGroupChats   bool   `json:"allow_group_chats"`
	AllowChannelChats bool   `json:"allow_channel_chats"`
}

type InlineKeyboardButton struct {
	Text                         string                      `json:"text"`
	Url                          string                      `json:"url"`
	CallbackData                 string                      `json:"callback_data"`
	WebApp                       WebAppInfo                  `json:"web_app"`
	LoginUrl                     LoginUrl                    `json:"login_url"`
	SwitchInlineQuery            string                      `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string                      `json:"switch_inline_query_current_chat"`
	SwitchInlineQueryChosenChat  SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
	CallbackGame                 CallbackGame                `json:"callback_game"`
	Pay                          bool                        `json:"pay"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type Story struct {
	Chat Chat `json:"chat"`
	ID   int  `json:"id"`
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

type PaidMedia struct {
	Type     string      `json:"type"`
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Duration int         `json:"duration"`
	Photo    []PhotoSize `json:"photo"`
	Video    Video       `json:"video"`
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

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionID        string    `json:"shipping_option_id"`
	OrderInfo               OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string    `json:"provider_payment_charge_id"`
}

type SharedUsers struct {
	UserID    int         `json:"user_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Username  string      `json:"username"`
	Photo     []PhotoSize `json:"photo"`
}

type UsersShared struct {
	RequestID int           `json:"request_id"`
	Users     []SharedUsers `json:"users"`
}

type ChatShared struct {
	RequestID int         `json:"request_id"`
	ChatID    int         `json:"chat_id"`
	Title     string      `json:"title"`
	Username  string      `json:"username"`
	Photo     []PhotoSize `json:"photo"`
}

type BackgroundFill struct {
	Type          string `json:"type"`
	Color         int    `json:"color"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
	Colors        []int  `json:"colors"`
}

type BackgroundType struct {
	Type             string         `json:"type"`
	Fill             BackgroundFill `json:"fill"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
	Document         Document       `json:"document"`
	IsBlurred        bool           `json:"is_blurred"`
	IsMoving         bool           `json:"is_moving"`
	Intensity        int            `json:"intensity"`
	IsInverted       bool           `json:"is_inverted"`
	ThemeName        string         `json:"theme_name"`
}

type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members"`
	HasPublicWinners              bool     `json:"has_public_winners"`
	PrizeDescription              string   `json:"prize_description"`
	CountryCodes                  []string `json:"country_codes"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count"`
}

type GiveawayCompleted struct {
	WinnerCount         int             `json:"winner_count"`
	UnclaimedPrizeCount int             `json:"unclaimed_prize_count"`
	GiveawayMessage     GiveawayMessage `json:"giveaway_message"`
}

type GiveawayWinners struct {
	Chat                          Chat   `json:"chat"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int    `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int    `json:"additional_chat_count"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count"`
	OnlyNewMembers                bool   `json:"only_new_members"`
	WasRefunded                   bool   `json:"was_refunded"`
	PrizeDescription              string `json:"prize_description"`
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   PassportFile   `json:"front_side"`
	ReverseSide PassportFile   `json:"reverse_side"`
	Selfie      PassportFile   `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
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

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     int    `json:"file_path"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YSHift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
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

type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	VoterCount   int             `json:"voter_count"`
}

type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEnities       []MessageEntity `json:"question_entities"`
	Options               []PollOption    `json:"options"`
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

type GiveawayMessage struct {
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
	VideoChatScheduled            VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   InlineKeyboardMarkup          `json:"reply_markup"`
}

type MaybeInaccessibleMessage struct {
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

type BusinessConnection struct {
	ID         string `json:"id"`
	User       User   `json:"user"`
	UserChatID int    `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}

type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        User           `json:"user"`
	ActorChat   Chat           `json:"actor_chat"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   int      `json:"offset"`
	ChatType string   `json:"chat_type"`
	Location Location `json:"location"`
}

type ChosenInlineResult struct {
	ResultID        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageID string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

type CallbackQuery struct {
	ID              string                   `json:"id"`
	From            User                     `json:"from"`
	Message         MaybeInaccessibleMessage `json:"message"`
	InlineMessageID string                   `json:"inline_message_id"`
	ChatInstance    string                   `json:"chat_instance"`
	Data            string                   `json:"data"`
	GameShortName   string                   `json:"game_short_name"`
}

type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionID string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info"`
}

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat Chat   `json:"voter_chat"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

type ChatMember struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsAnonymous           bool   `json:"is_anonymous"`
	CustomTitle           string `json:"custom_title"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanManageChat         bool   `json:"can_manage_chat"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanManageVideoChats   bool   `json:"can_manage_video_chats"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPostStories        bool   `json:"can_post_stories"`
	CanEditStories        bool   `json:"can_edit_stories"`
	CanDeleteStories      bool   `json:"can_delete_stories"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int    `json:"until_date"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Create                  User   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name"`
	ExpireDate              int    `json:"expire_date"`
	MemberLimit             int    `json:"member_limit"`
	PendingJoinRequestCount int    `jsos:"pending_join_request_count"`
	SubscriptionPeriod      int    `json:"subscription_period"`
	SubscriptionPrice       int    `json:"subscription_price"`
}

type ChatMemberUpdated struct {
	Chat                    Chat           `json:"chat"`
	From                    User           `json:"from"`
	Date                    int            `json:"date"`
	OldChatMember           ChatMember     `json:"old_chat_member"`
	NewChatMember           ChatMember     `json:"new_chat_member"`
	InviteLink              ChatInviteLink `json:"invite_link"`
	ViaJoinRequest          bool           `json:"via_join_request"`
	ViaChatFolderInviteLink bool           `json:"via_chat_folder_invite_link"`
}

type ChatJoinRequest struct {
	Chat       Chat           `json:"chat"`
	User       User           `json:"user"`
	UserChatID int            `json:"user_chat_id"`
	Date       int            `json:"date"`
	Bio        string         `json:"bio"`
	InviteLink ChatInviteLink `json:"invite_link"`
}

type ChatBoostSource struct {
	Source            string `json:"source"`
	User              User   `json:"user"`
	GiveawayMessageID int    `json:"giveaway_message_id"`
	IsUnclaimed       bool   `json:"is_unclaimed"`
}

type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	CxpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
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
	DeletedBusinessMessage BusinessMessagesDeleted     `json:"deleted_business_messages"`
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
