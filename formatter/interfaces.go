package formatter

import (
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/types"
)

type handlerMedia interface {
	multipartFields(*multipart.Writer, *[]interface{}, int, bool) error
	jsonFileds() ([]byte, error)
}

type handlerKB interface {
	MultipartFields(*multipart.Writer) error
	JsonFields() ([]byte, error)
}

type IPhoto interface {
	// Receives Telegram ID of a photo, URL-link with the a photo or name of a photo in storage
	WritePhoto(string)

	// Receives a string that will be a caption of the photo. If you're going to send
	// one picture and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(string)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(string)

	WriteCaptionEntities([]*types.MessageEntity)
	WriteShowCaptionAboveMedia()

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler()

	// Must be used. There are three options to send data: formatter.Storage, formatter.URL, formatter.Telegram.
	// It shows where the bot must wait the photo from
	WriteGottenFrom(int)

	// You can call this function after calling Send(). It returns you an array with some data about the photo you just sent. 4 is the biggest amount of data. Sometimes there might be nil
	GetResponse() [4]*responseData
}

type IVideo interface {
	// Receives Telegram ID of a video, Url-link with the a video or name of a video in storage
	WriteVideo(string)

	// Receives a string that will be a caption of the video. If you're going to send
	// one video and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string)

	WriteCaptionEntities(captionEntities []*types.MessageEntity)
	WriteDuration(duration int)

	// Must be used. There are three options to send data: formatter.Storage, formatter.URL, formatter.Telegram
	// It shows where the bot must wait the video from
	WriteGottenFrom(gottenfrom int)

	// Doesn't recieve anything. If the function was called, the photo will be blured
	WriteHasSpoiler()
	WriteHeight(height int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string)

	WriteShowCaptionAboveMedia()
	WriteSupportsStreaming()
	WriteThumbnail(thumbnail string)
	WriteThumbnailGottenFrom(gottenfrom int)
	WriteWidth(width int)

	// You can call this function after calling Send(). It returns you a structure with some data about the video you just sent
	GetResponse() *types.Video
}

type IAudio interface {
	// Receives Telegram ID of an audio, Url-link with the an audio or name of an audio in storage
	WriteAudio(string)

	// Receives a string that will be a caption of the audio. If you're going to send
	// one audio and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string)

	WriteCaptionEntities(captionEntities []*types.MessageEntity)
	WriteDuration(duration int)

	// Must be used. There are three options to send data: formatter.Storage, formatter.URL, formatter.Telegram
	// It shows where the bot must wait the audio from
	WriteGottenFrom(gottenfrom int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string)

	WritePerformer(performer string)
	WriteThumbnail(thumbnail string)
	WriteThumbnailGottenFrom(gottenfrom int)
	WriteTitle(title string)

	// You can call this function after calling Send(). It returns you a structure with some data about the audio you just sent
	GetResponse() *types.Audio
}

type IDocument interface {
	// Receives a string that will be a caption of the document. If you're going to send
	// one document and you want to send a text-message you have to use it insted of WriteString(string)
	// in MSGInformation interface
	WriteCaption(caption string)

	WriteCaptionEntities(captionEntities []*types.MessageEntity)
	WriteDisableContentTypeDetection()

	// Receives Telegram ID of a document, Url-link with the a document or name of a document in storage
	WriteDocument(document string)

	// Must be used. There are three options to send data: formatter.Storage, formatter.URL, formatter.Telegram. It shows where the bot must wait the document from
	WriteGottenFrom(gottenfrom int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string)

	WriteThumbnail(thumbnail string)
	WriteThumbnailGottenFrom(gottenfrom int)

	// You can call this function after calling Send(). It returns you a structure with some data about the document you just sent
	GetResponse() *types.Document
}

type MSGInformation interface {
	// Doesn't revieve anything. If the function was called, the message you send to a client will be gotten without a notification
	WriteDisableNotification()

	WriteEntities(entities []*types.MessageEntity)
	WriteLinkPreviewOptions(lpo *types.LinkPreviewOptions)
	WriteMessageEffectID(messageID string)
	WriteMessageThreadID(messageID int)

	// Receives spicific kind of string. Adds some opportunity to transform the text-message. There are 3 options: types.HTML, types.Markdown and...
	WriteParseMode(parsemode string)

	// Doesn't recieve anything. If the function was called, the text-message will be blured
	WriteProtectContent()

	// ???
	WriteReplyMarkup(markup interface{}) error
	WriteReplyParametrs(params *types.ReplyParameters)

	// The text content of the message
	WriteString(text string)
}

type IChat interface {
	WriteBusinessConnectionID(connectionID string)

	// These the functions might be used only if you need to send a message (!NOT RESPONSE!) to a client.
	// There are some information about the chat the request was gotten from, so you do not need it at all.
	//
	// Receives a number (int). Must be the chatID you want to send a message to
	WriteChatID(chatID int)

	// Receives the name of the chat. Must be the the name of you want to send a message to
	WriteChatName(chatname string)

	// Receives the URL of the chat. Must be the the URL of you want to send a message to
	WriteChatURL(chatURL string)
}

type IReply interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(line int, pos int) (IReplyButton, error)

	// Receives a slice of integers and sets the plan of future reply-keyboard. Be careful with data, because one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the reply-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set(plan []int)
}

type IReplyButton interface {
	WriteRequestChat(reqch *types.KeyboardButtonRequestChat)
	WriteRequestContact()
	WriteRequestLocation()
	WriteRequestPoll(poll *types.KeyboardButtonPollType)
	WriteRequestUsers(requs *types.KeyboardButtonRequestUsers)

	// The text content that will be on the button
	WriteString(text string)

	WriteWebApp(webapp *types.WebAppInfo)
}

type IInline interface {
	// Receives 2 coordinates. The first one is a line you expect to see the button and the second one is a position on the line you just mentioned.
	// The coordinates mustn't be a random numbers. You have to call function Set() first. Returns an interface of the new created button
	// and might return and error, if the coordinates are wrong
	NewButton(line int, pos int) (IInlineButton, error)

	// Receives a slice of integers and sets the plan of future inline-keyboard. Be careful with data, because one array cell is one array cell is considered
	// as row, so, for example: a slice []int{2, 1, 4} means that in the inline-keyboard will be 3 rows and in the first one there'll be 2 buttons, in the next one: just one
	// and in the last one: 4 buttons
	Set(plan []int)
}

type IInlineButton interface {
	// Receives some text string that will be received by the bot after a client pressed on the button
	WriteCallbackData(text string)

	WriteCallbackGame(game *types.CallbackGame)
	WriteLoginUrl(logurl *types.LoginUrl)
	WritePay()

	// Receives some text string that will be seen by client on the button
	WriteString(text string)

	WriteSwitchInlineQuery(sw string)
	WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat)
	WriteSwitchInlineQueryCurrentChat(swcch string)

	// Receives a URL string that will be received by the bot after a client pressed on the button
	WriteURL(url string)

	WriteWebApp(wbapp *types.WebAppInfo)
}
