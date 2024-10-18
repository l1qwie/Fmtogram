package formatter

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	interfacePhoto    string = "IPhoto"
	interfaceVideo    string = "IVideo"
	interfaceAudio    string = "IAudio"
	interfaceDocument string = "IDocument"
	interfaceInf      string = "MSG Information"
	interfaceChat     string = "IChat"
	inKB              string = "Inline Keyboard"
	replyKB           string = "Reply Keyboard"
	button            string = "Button"
	inbtn             string = "Inline Button"
	rpbtn             string = "Reply Button"
	checkString       int    = -1
	checkArray        int    = -2
	checkBool         int    = -3
	checkInt          int    = -4
)

func (ph *photo) WritePhoto(photo string) {
	if !isItEmply(ph, checkString, ph.Photo) {
		logs.DataIsntEmply(interfacePhoto, "Photo", ph.Photo)
	}
	ph.Photo, ph.Media = photo, photo
	logs.DataWrittenSuccessfully(interfacePhoto, "Photo")
}

func (ph *photo) WriteCaption(caption string) {
	if !isItEmply(ph, checkString, ph.Caption) {
		logs.DataIsntEmply(interfacePhoto, "Caption", ph.Caption)
	}
	ph.Caption = caption
	logs.DataWrittenSuccessfully(interfacePhoto, "Caption")
}

func (ph *photo) WriteParseMode(parsemode string) {
	if !isItEmply(ph, checkString, ph.ParseMode) {
		logs.DataIsntEmply(interfacePhoto, "Parse Mode", ph.ParseMode)
	}
	ph.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfacePhoto, "Parse Mode")
}

func (ph *photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ph, checkArray, ph.CaptionEntities) {
		logs.DataIsntEmply(interfacePhoto, "Caption Entities", ph.CaptionEntities)
	}
	ph.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfacePhoto, "Caption Entities")
}

func (ph *photo) WriteShowCaptionAboveMedia() {
	if !isItEmply(ph, checkBool, ph.ShowCaptionAboveMedia) {
		logs.DataIsntEmply(interfacePhoto, "Show Caption Above Media", ph.ShowCaptionAboveMedia)
	}
	ph.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", interfacePhoto, ph.ShowCaptionAboveMedia)
}

func (ph *photo) WriteHasSpoiler() {
	if !isItEmply(ph, checkBool, ph.HasSpoiler) {
		logs.DataIsntEmply(interfacePhoto, "Has Spoiler", ph.HasSpoiler)
	}
	ph.HasSpoiler = true
	logs.SettedParam("Has Spoiler", interfacePhoto, ph.HasSpoiler)
}

func (ph *photo) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(ph, checkInt, ph.GottenFrom) {
		logs.DataIsntEmply(interfacePhoto, "Gotten From", ph.GottenFrom)
	}
	ph.GottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfacePhoto, "Gotten From")
}

func (ph *photo) GetResponse() [4]*responseData {
	return ph.ResponseData
}

func (vd *video) WriteVideo(video string) {
	if !isItEmply(vd, checkString, vd.Video) {
		logs.DataIsntEmply(interfaceVideo, "Video", vd.Video)
	}
	vd.Video, vd.Media = video, video
	logs.DataWrittenSuccessfully(interfaceVideo, "Video")
}

func (vd *video) WriteThumbnail(thumbnail string) {
	if !isItEmply(vd, checkString, vd.Thumbnail) {
		logs.DataIsntEmply("Video", "Thumbnail", vd.Thumbnail)
	}
	vd.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail")
}

func (vd *video) WriteCaption(caption string) {
	if !isItEmply(vd, checkString, vd.Caption) {
		logs.DataIsntEmply("Video", "Caption", vd.Thumbnail)
	}
	vd.Caption = caption
	logs.DataWrittenSuccessfully(interfaceVideo, "Caption")
}

func (vd *video) WriteParseMode(parsemode string) {
	if !isItEmply(vd, checkString, vd.ParseMode) {
		logs.DataIsntEmply("Video", "Parse Mode", vd.Thumbnail)
	}
	vd.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceVideo, "Parse Mode")
}

func (vd *video) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(vd, checkArray, vd.CaptionEntities) {
		logs.DataIsntEmply("Video", "Caption Entities", vd.Thumbnail)
	}
	vd.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceVideo, "Caption Entities")
}

func (vd *video) WriteShowCaptionAboveMedia() {
	if !isItEmply(vd, checkBool, vd.ShowCaptionAboveMedia) {
		logs.DataIsntEmply("Video", "Show Caption Above Media", vd.ShowCaptionAboveMedia)
	}
	vd.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", "Video", vd.ShowCaptionAboveMedia)
}

func (vd *video) WriteWidth(width int) {
	if !isItEmply(vd, checkInt, vd.Width) {
		logs.DataIsntEmply("Video", "Width", vd.Width)
	}
	vd.Width = width
	logs.DataWrittenSuccessfully(interfaceVideo, "Width")
}

func (vd *video) WriteHeight(height int) {
	if !isItEmply(vd, checkInt, vd.Height) {
		logs.DataIsntEmply("Video", "Height", vd.Height)
	}
	vd.Height = height
	logs.DataWrittenSuccessfully(interfaceVideo, "Height")
}

func (vd *video) WriteDuration(duration int) {
	if !isItEmply(vd, checkInt, vd.Duration) {
		logs.DataIsntEmply("Video", "Duration", vd.Duration)
	}
	vd.Duration = duration
	logs.DataWrittenSuccessfully(interfaceVideo, "Duration")
}

func (vd *video) WriteSupportsStreaming() {
	if !isItEmply(vd, checkBool, vd.SupportsStreaming) {
		logs.DataIsntEmply("Video", "Supports Streaming", vd.SupportsStreaming)
	}
	vd.SupportsStreaming = true
	logs.SettedParam("Supports Streaming", "Video", vd.SupportsStreaming)
}

func (vd *video) WriteHasSpoiler() {
	if !isItEmply(vd, checkBool, vd.HasSpoiler) {
		logs.DataIsntEmply("Video", "Has Spoiler", vd.HasSpoiler)
	}
	vd.HasSpoiler = true
	logs.SettedParam("Has Spoiler", "Video", vd.SupportsStreaming)
}

func (vd *video) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(vd, checkInt, vd.VideoGottenFrom) {
		logs.DataIsntEmply(interfaceVideo, "Gotten From", vd.VideoGottenFrom)
	}
	vd.VideoGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceVideo, "Gotten From")
}

func (vd *video) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(vd, checkInt, vd.ThumbnailGottenFrom) {
		logs.DataIsntEmply(interfaceVideo, "Thumbnail Gotten From", vd.ThumbnailGottenFrom)
	}
	vd.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceVideo, "Thumbnail Gotten From")
}

func (vd *video) GetResponse() *types.Video {
	return vd.ResponseData
}

func (ad *audio) WriteAudio(audio string) {
	if !isItEmply(ad, checkString, ad.Audio) {
		logs.DataIsntEmply(interfaceAudio, "Audio", ad.Audio)
	}
	ad.Audio, ad.Media = audio, audio
	logs.DataWrittenSuccessfully(interfaceAudio, "Audio")
}

func (ad *audio) WriteThumbnail(thumbnail string) {
	if !isItEmply(ad, checkString, ad.Thumbnail) {
		logs.DataIsntEmply("Audio", "Thumbnail", ad.Thumbnail)
	}
	ad.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceAudio, "Thumbnail")
}

func (ad *audio) WriteCaption(caption string) {
	if !isItEmply(ad, checkString, ad.Caption) {
		logs.DataIsntEmply("Audio", "Caption", ad.Caption)
	}
	ad.Caption = caption
	logs.DataWrittenSuccessfully(interfaceAudio, "Caption")
}

func (ad *audio) WriteParseMode(parsemode string) {
	if !isItEmply(ad, checkString, ad.ParseMode) {
		logs.DataIsntEmply("Audio", "Parse Mode", ad.ParseMode)
	}
	ad.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceAudio, "Parse Mode")
}

func (ad *audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ad, checkArray, ad.CaptionEntities) {
		logs.DataIsntEmply("Audio", "Caption Entities", ad.CaptionEntities)
	}
	ad.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceAudio, "Caption Entities")
}

func (ad *audio) WriteDuration(duration int) {
	if !isItEmply(ad, checkInt, ad.Duration) {
		logs.DataIsntEmply("Audio", "Duration", ad.Duration)
	}
	ad.Duration = duration
	logs.DataWrittenSuccessfully(interfaceAudio, "Duration")
}

func (ad *audio) WritePerformer(performer string) {
	if !isItEmply(ad, checkString, ad.Performer) {
		logs.DataIsntEmply("Audio", "Performer", ad.Performer)
	}
	ad.Performer = performer
	logs.DataWrittenSuccessfully(interfaceAudio, "Performer")
}

func (ad *audio) WriteTitle(title string) {
	if !isItEmply(ad, checkString, ad.Title) {
		logs.DataIsntEmply("Audio", "Title", ad.Title)
	}
	ad.Title = title
	logs.DataWrittenSuccessfully(interfaceAudio, "Title")
}

func (ad *audio) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(ad, checkInt, ad.AudioGottenFrom) {
		logs.DataIsntEmply(interfaceAudio, "Gotten From", ad.AudioGottenFrom)
	}
	ad.AudioGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceAudio, "Gotten From")
}

func (ad *audio) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(ad, checkInt, ad.ThumbnailGottenFrom) {
		logs.DataIsntEmply(interfaceAudio, "Thumbnail Gotten From", ad.ThumbnailGottenFrom)
	}
	ad.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceAudio, "Thumbnail Gotten From")
}

func (ad *audio) GetResponse() *types.Audio {
	return ad.ResponseData
}

func (dc *document) WriteDocument(document string) {
	if !isItEmply(dc, checkString, dc.Document) {
		logs.DataIsntEmply("Document", "Document", dc.Document)
	}
	dc.Document = document
	logs.DataWrittenSuccessfully(interfaceDocument, "Document")
}

func (dc *document) WriteThumbnail(thumbnail string) {
	if !isItEmply(dc, checkString, dc.Thumbnail) {
		logs.DataIsntEmply("Document", "Thumbnail", dc.Thumbnail)
	}
	dc.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail")
}

func (dc *document) WriteCaption(caption string) {
	if !isItEmply(dc, checkString, dc.Caption) {
		logs.DataIsntEmply("Document", "Caption", dc.Caption)
	}
	dc.Caption = caption
	logs.DataWrittenSuccessfully(interfaceDocument, "Caption")
}

func (dc *document) WriteParseMode(parsemode string) {
	if !isItEmply(dc, checkString, dc.ParseMode) {
		logs.DataIsntEmply("Document", "Parse Mode", dc.ParseMode)
	}
	dc.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceDocument, "Parse Mode")
}

func (dc *document) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(dc, checkArray, dc.CaptionEntities) {
		logs.DataIsntEmply("Document", "Caption Entities", dc.CaptionEntities)
	}
	dc.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(interfaceDocument, "Caption Entities")
}

func (dc *document) WriteDisableContentTypeDetection() {
	if !isItEmply(dc, checkArray, dc.DisableContentTypeDetection) {
		logs.DataIsntEmply("Document", "Disable Content Type Detection", dc.DisableContentTypeDetection)
	}
	dc.DisableContentTypeDetection = true
	logs.SettedParam("Disable Content Type Detection", "Document", dc.DisableContentTypeDetection)
}

func (dc *document) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.DocumentGottenFrom) {
		logs.DataIsntEmply(interfaceDocument, "Gotten From", dc.DocumentGottenFrom)
	}
	dc.DocumentGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceDocument, "Gotten From")
}

func (dc *document) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.ThumbnailGottenFrom) {
		logs.DataIsntEmply(interfaceDocument, "Thumbnail Gotten From", dc.ThumbnailGottenFrom)
	}
	dc.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(interfaceDocument, "Thumbnail Gotten From")
}

func (dc *document) GetResponse() *types.Document {
	return dc.ResponseData
}

func (inf *information) WriteString(text string) {
	if inf.Text != "" {
		logs.DataIsntEmply(interfaceInf, "Text", inf.Text)
	}
	inf.Text = text
	logs.DataWrittenSuccessfully(interfaceInf, "Text")
}

func (inf *information) WriteParseMode(parsemode string) {
	if inf.ParseMode != "" {
		logs.DataIsntEmply(interfaceInf, "Parse Mode", inf.ParseMode)
	}
	inf.ParseMode = parsemode
	logs.DataWrittenSuccessfully(interfaceInf, "Parse Mode")
}

func (inf *information) WriteMessageThreadID(messageID int) {
	if inf.MessageThreadID != 0 {
		logs.DataIsntEmply(interfaceInf, "Message Thread ID", inf.MessageThreadID)
	}
	inf.MessageThreadID = messageID
	logs.DataWrittenSuccessfully(interfaceInf, "Message Thread ID")
}

func (inf *information) WriteDisableNotification() {
	if inf.DisableNotification {
		logs.DataIsntEmply(interfaceInf, "Disable Notification", inf.DisableNotification)
	}
	inf.DisableNotification = true
	logs.SettedParam("Disable Notification", interfaceInf, inf.DisableNotification)
}

func (inf *information) WriteProtectContent() {
	if inf.ProtectContent {
		logs.DataIsntEmply(interfaceInf, "Protect Content", inf.ProtectContent)
	}
	inf.ProtectContent = true
	logs.SettedParam("Protect Content", interfaceInf, inf.ProtectContent)
}

func (inf *information) WriteMessageEffectID(messageID string) {
	if inf.MessageEffectID != "" {
		logs.DataIsntEmply(interfaceInf, "Message Effect ID", inf.MessageEffectID)
	}
	inf.MessageEffectID = messageID
	logs.DataWrittenSuccessfully(interfaceInf, "Message Effect ID")
}

func (inf *information) WriteReplyParametrs(params *types.ReplyParameters) {
	if inf.ReplyParameters != nil {
		logs.DataIsntEmply(interfaceInf, "Reply Parametrs", inf.ReplyParameters)
	}
	inf.ReplyParameters = params
	logs.DataWrittenSuccessfully(interfaceInf, "Reply Parametrs")
}

func (inf *information) WriteEntities(entities []*types.MessageEntity) {
	if len(inf.Entities) != 0 {
		logs.DataIsntEmply(interfaceInf, "Entities", inf.Entities)
	}
	inf.Entities = entities
	logs.DataWrittenSuccessfully(interfaceInf, "Entities")
}

func (inf *information) WriteLinkPreviewOptions(lpo *types.LinkPreviewOptions) {
	if inf.LinkPreviewOptions != nil {
		logs.DataIsntEmply(interfaceInf, "Link Preview Options", inf.LinkPreviewOptions)
	}
	inf.LinkPreviewOptions = lpo
	logs.DataWrittenSuccessfully(interfaceInf, "Link Preview Options")
}

func checkReplyTypes(markup interface{}) bool {
	var ok bool
	switch markup.(type) {
	case *types.InlineKeyboardMarkup:
		ok = true
	case *types.ReplyKeyboardMarkup:
		ok = true
	case *types.ReplyKeyboardRemove:
		ok = true
	case *types.ForceReply:
		ok = true
	}

	return ok
}

func (inf *information) WriteReplyMarkup(markup interface{}) error {
	var err error
	ok := checkReplyTypes(markup)
	if ok {
		if inf.ReplyMarkup != nil {
			logs.DataIsntEmply(interfaceInf, "Reply Markup", inf.ReplyMarkup)
		}
		inf.ReplyMarkup = markup
		logs.DataWrittenSuccessfully(interfaceInf, "Reply Markup")
	} else {
		err = fmt.Errorf("WriteReplyMarkup is waiting an interface{}, but the exactly type of the interface{} must be *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove or *types.ForceReply")
	}
	return err
}

func (ch *chat) WriteChatID(chatID int) {
	if ch.ID != nil {
		logs.DataIsntEmply(interfaceChat, "Chat ID", ch.ID)
	}
	ch.ID = chatID
	logs.DataWrittenSuccessfully(interfaceChat, "Chat ID")
}

func (ch *chat) WriteChatName(chatname string) {
	if ch.ID != nil {
		logs.DataIsntEmply(interfaceChat, "Chat Name", ch.ID)
	}
	ch.ID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully(interfaceChat, "Chat Name")
}

func (ch *chat) WriteChatURL(chatURL string) {
	if ch.ID != nil {
		logs.DataIsntEmply(interfaceChat, "Chat URL", ch.ID)
	}
	ch.ID = chatURL
	logs.DataWrittenSuccessfully(interfaceChat, "Chat URL")
}

func (ch *chat) WriteBusinessConnectionID(connectionID string) {
	if ch.BusinessConnectionID != "" {
		logs.DataIsntEmply(interfaceChat, "Business Connection ID", ch.BusinessConnectionID)
	}
	ch.BusinessConnectionID = connectionID
	logs.DataWrittenSuccessfully(interfaceChat, "Business Connection ID")
}

func (in *inline) Set(plan []int) {
	in.Keyboard = new(inlineKeyboard)
	in.Keyboard.InlineKeyboard = make([][]*inlineKeyboardButton, len(plan))
	for i := range in.Keyboard.InlineKeyboard {
		in.Keyboard.InlineKeyboard[i] = make([]*inlineKeyboardButton, plan[i])
	}
}

func (in *inline) NewButton(line, pos int) (IInlineButton, error) {
	var err error

	if (line >= 0) && (pos >= 0) && len(in.Keyboard.InlineKeyboard) > line && len(in.Keyboard.InlineKeyboard[line]) > pos {

		if in.Keyboard.InlineKeyboard[line][pos] != nil {
			logs.DataIsntEmply(inKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), in.Keyboard.InlineKeyboard[line][pos])
		}
		in.Keyboard.InlineKeyboard[line][pos] = new(inlineKeyboardButton)

		return in.Keyboard.InlineKeyboard[line][pos], nil

	} else {

		if len(in.Keyboard.InlineKeyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(in.Keyboard.InlineKeyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (inb *inlineKeyboardButton) WriteString(text string) {
	if inb.Text != "" {
		logs.DataIsntEmply(inbtn, "Text", inb.Text)
	}
	inb.Text = text
	logs.DataWrittenSuccessfully(inbtn, "Text")
}

func (inb *inlineKeyboardButton) WriteURL(url string) {
	if inb.Url != "" {
		logs.DataIsntEmply(inbtn, "URL", inb.Url)
	}
	inb.Url = url
	logs.DataWrittenSuccessfully(inbtn, "URL")
}

func (inb *inlineKeyboardButton) WriteCallbackData(text string) {
	if inb.CallbackData != "" {
		logs.DataIsntEmply(inbtn, "Callback Data", inb.CallbackData)
	}
	inb.CallbackData = text
	logs.DataWrittenSuccessfully(inbtn, "Callback Data")
}

func (inb *inlineKeyboardButton) WriteWebApp(wbapp *types.WebAppInfo) {
	if inb.WebApp != nil {
		logs.DataIsntEmply(inbtn, "Web App", *inb.WebApp)
	}
	inb.WebApp = wbapp
	logs.DataWrittenSuccessfully(inbtn, "Web App")
}

func (inb *inlineKeyboardButton) WriteLoginUrl(logurl *types.LoginUrl) {
	if inb.LoginUrl != nil {
		logs.DataIsntEmply(inbtn, "Login URL", *inb.LoginUrl)
	}
	inb.LoginUrl = logurl
	logs.DataWrittenSuccessfully(inbtn, "Login URL")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQuery(sw string) {
	if inb.SwitchInlineQuery != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query", inb.SwitchInlineQuery)
	}
	inb.SwitchInlineQuery = sw
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) {
	if inb.SwitchInlineQueryCurrentChat != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Current Chat", inb.SwitchInlineQueryCurrentChat)
	}
	inb.SwitchInlineQueryCurrentChat = swcch
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Current Chat")
}

func (inb *inlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) {
	if inb.SwitchInlineQueryChosenChat != nil {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Chosen Chat", inb.SwitchInlineQueryChosenChat)
	}
	inb.SwitchInlineQueryChosenChat = sw
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Chosen Chat")
}

func (inb *inlineKeyboardButton) WriteCallbackGame(game *types.CallbackGame) {
	if inb.CallbackGame != nil {
		logs.DataIsntEmply(inbtn, "Callback Game", inb.CallbackGame)
	}
	inb.CallbackGame = game
	logs.DataWrittenSuccessfully(inbtn, "Callback Game")
}

func (inb *inlineKeyboardButton) WritePay() {
	if inb.Pay {
		logs.DataIsntEmply(inbtn, "Pay", inb.Pay)
	}
	inb.Pay = true
	logs.SettedParam("Pay", inbtn, true)
}

func (rp *reply) Set(plan []int) {
	rp.Keyboard = new(replyKeyboard)
	rp.Keyboard.Keyboard = make([][]*replyKeyboardButton, len(plan))
	for i := range rp.Keyboard.Keyboard {
		rp.Keyboard.Keyboard[i] = make([]*replyKeyboardButton, plan[i])
	}
}

func (rp *reply) NewButton(line, pos int) (IReplyButton, error) {
	var err error

	if (line >= 0) && (pos >= 0) && (len(rp.Keyboard.Keyboard) > line) && (len(rp.Keyboard.Keyboard[line]) > pos) {

		if rp.Keyboard.Keyboard[line][pos] != nil {
			logs.DataIsntEmply(replyKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), rp.Keyboard.Keyboard[line][pos])
		}
		rp.Keyboard.Keyboard[line][pos] = new(replyKeyboardButton)

		return rp.Keyboard.Keyboard[line][pos], nil

	} else {
		if len(rp.Keyboard.Keyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(rp.Keyboard.Keyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (rpb *replyKeyboardButton) WriteString(text string) {
	if rpb.Text != "" {
		logs.DataIsntEmply(rpbtn, "Text", rpb.Text)
	}
	rpb.Text = text
	logs.DataWrittenSuccessfully(rpbtn, "Text")
}

func (rpb *replyKeyboardButton) WriteRequestUsers(requs *types.KeyboardButtonRequestUsers) {
	if rpb.RequestUsers != nil {
		logs.DataIsntEmply(rpbtn, "Request Users", rpb.RequestUsers)
	}
	rpb.RequestUsers = requs
	logs.DataWrittenSuccessfully(rpbtn, "Request Users")
}

func (rpb *replyKeyboardButton) WriteRequestChat(reqch *types.KeyboardButtonRequestChat) {
	if rpb.RequestChat != nil {
		logs.DataIsntEmply(rpbtn, "Request Chat", rpb.RequestChat)
	}
	rpb.RequestChat = reqch
	logs.DataWrittenSuccessfully(rpbtn, "Request Chat")
}

func (rpb *replyKeyboardButton) WriteRequestContact() {
	if rpb.RequestContact {
		logs.DataIsntEmply(rpbtn, "Request Contact", rpb.RequestContact)
	}
	rpb.RequestContact = true
	logs.SettedParam("Request Contact", rpbtn, true)
}

func (rpb *replyKeyboardButton) WriteRequestLocation() {
	if rpb.RequestLocation {
		logs.DataIsntEmply(rpbtn, "Request Location", rpb.RequestLocation)
	}
	rpb.RequestLocation = true
	logs.SettedParam("Request Location", rpbtn, true)
}

func (rpb *replyKeyboardButton) WriteRequestPoll(poll *types.KeyboardButtonPollType) {
	if rpb.RequestPoll != nil {
		logs.DataIsntEmply(rpbtn, "Request Poll", rpb.RequestPoll)
	}
	rpb.RequestPoll = poll
	logs.DataWrittenSuccessfully(rpbtn, "Request Poll")
}

func (rpb *replyKeyboardButton) WriteWebApp(webapp *types.WebAppInfo) {
	if rpb.WebApp != nil {
		logs.DataIsntEmply(rpbtn, "Web App", rpb.WebApp)
	}
	rpb.WebApp = webapp
	logs.DataWrittenSuccessfully(rpbtn, "Web App")
}
