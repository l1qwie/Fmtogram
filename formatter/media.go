package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/logs"
)

const (
	photoMethod     string = "sendPhoto"
	audipMethod     string = "sendAudio"
	documentMethod  string = "sendDocument"
	videoMethod     string = "sendVideo"
	animationMethod string = "sendAnimation"
	voiceMethod     string = "sendVoice"
	videoNoteMethod string = "sendVideoNote"
	mediaGroup      string = "sendMediaGroup"
)

type handleMedia interface {
	uniqueFields() (string, error)
}

type handleInputMedia interface {
	uniqueInputFields(*multipart.Writer, int) error
}

type mediaFields struct {
	data   handleMedia
	method string
}

func (fm *Formatter) whereToAdd(mediadata, mediatype string, where int) {
	if len(fm.kindofmedia) == 1 {
		fm.Media.Group = append(fm.Media.Group, InputMedia{Type: fm.mediatype[0], Media: fm.Media.mediaFile.mediaData})
		fm.Media.Group = append(fm.Media.Group, InputMedia{Type: mediatype, Media: mediadata})
	} else if len(fm.kindofmedia) >= 2 {
		fm.Media.Group = append(fm.Media.Group, InputMedia{Type: mediatype, Media: mediadata})
	} else {
		fm.Media.mediaFile.mediaData = mediadata
	}
	fm.kindofmedia = append(fm.kindofmedia, where)
	fm.mediatype = append(fm.mediatype, mediatype)
}

func (fm *Formatter) AddPhotoFromStorage(path string) {
	fm.whereToAdd(path, "photo", fromStorage)
	logs.DataWrittenSuccessfully("A Photo From The Storage")
}

func (fm *Formatter) AddPhotoFromTG(photoID string) {
	fm.whereToAdd(photoID, "photo", fromTelegram)
	logs.DataWrittenSuccessfully("A Photo Telegram ID")
}

func (fm *Formatter) AddPhotoFromInternet(URL string) {
	fm.whereToAdd(URL, "photo", fromInternet)
	logs.DataWrittenSuccessfully("A Photo URL From The Internet")
}

func (fm *Formatter) AddAudioFromStorage(path string) {
	fm.whereToAdd(path, "audio", fromStorage)
	logs.DataWrittenSuccessfully("An Audio Path From The Storage")
}

func (fm *Formatter) AddAudioFromTG(audioID string) {
	fm.whereToAdd(audioID, "audio", fromTelegram)
	logs.DataWrittenSuccessfully("An Audio Telegram ID")
}

func (fm *Formatter) AddAudioFromInternet(URL string) {
	fm.whereToAdd(URL, "audio", fromInternet)
	logs.DataWrittenSuccessfully("An Audio URL From The Internet")
}

func (fm *Formatter) AddDocumentFromStorage(path string) {
	fm.whereToAdd(path, "document", fromStorage)
	logs.DataWrittenSuccessfully("A Document Path From The Storage")
}

func (fm *Formatter) AddDocumentFromTG(documentID string) {
	fm.whereToAdd(documentID, "document", fromTelegram)
	logs.DataWrittenSuccessfully("A Document Telegram ID")
}

func (fm *Formatter) AddDocumentFromInternet(URL string) {
	fm.whereToAdd(URL, "document", fromInternet)
	logs.DataWrittenSuccessfully("A Document URL From The Internet")
}

func (fm *Formatter) AddVideoFromStorage(path string) {
	fm.whereToAdd(path, "video", fromStorage)
	logs.DataWrittenSuccessfully("A Video From The Storage")
}

func (fm *Formatter) AddVideoFromTG(videoID string) {
	fm.whereToAdd(videoID, "video", fromTelegram)
	logs.DataWrittenSuccessfully("A Video Telegram ID")
}

func (fm *Formatter) AddVideoFromInternet(URL string) {
	fm.whereToAdd(URL, "video", fromInternet)
	logs.DataWrittenSuccessfully("A Video URL From The Internet")
}

func (fm *Formatter) AddAnimationAnimationFromStorage(path string) {
	fm.whereToAdd(path, "animation", fromStorage)
	logs.DataWrittenSuccessfully("An Animation From The Storage")
}

func (fm *Formatter) AddAnimationFromTG(animationID string) {
	fm.whereToAdd(animationID, "animation", fromTelegram)
	logs.DataWrittenSuccessfully("An Animation Telegram ID")
}

func (fm *Formatter) AddAnimationFromInternet(URL string) {
	fm.whereToAdd(URL, "animation", fromInternet)
	logs.DataWrittenSuccessfully("An Animation URL From The Internet")
}

func (fm *Formatter) AddVoiceFromStorage(path string) {
	fm.whereToAdd(path, "voice", fromStorage)
	logs.DataWrittenSuccessfully("A Voice From The Storage")
}

func (fm *Formatter) AddVoiceFromTG(voiceID string) {
	fm.whereToAdd(voiceID, "voice", fromTelegram)
	logs.DataWrittenSuccessfully("A Voice Telegram ID")
}

func (fm *Formatter) AddVoiceFromInternet(URL string) {
	fm.whereToAdd(URL, "voice", fromInternet)
	logs.DataWrittenSuccessfully("A Voice URL From The Internet")
}

func (fm *Formatter) AddVideoNoteFromStorage(path string) {
	fm.whereToAdd(path, "video-note", fromStorage)
	logs.DataWrittenSuccessfully("A VideoNote From The Storage")
}

func (fm *Formatter) AddVideoNoteFromTG(videoNoteID string) {
	fm.whereToAdd(videoNoteID, "video-note", fromTelegram)
	logs.DataWrittenSuccessfully("A VideoNote Telegram ID")
}

func (fm *Formatter) AddVideoNoteFromInternet(URL string) {
	fm.whereToAdd(URL, "video-note", fromInternet)
	logs.DataWrittenSuccessfully("A VideoNote URL From The Internet")
}

// func (fm *Formatter) AddMapOfMedia(arr []types.Media) {
// 	i := 0
// 	fm.Message.InputMedia = make([]types.Media, len(arr))
// 	for _, val := range arr {
// 		fm.Message.InputMedia[i].Type = val.Type
// 		fm.Message.InputMedia[i].Media = val.Media
// 		i++
// 	}
// }

// func (fm *Formatter) createMediaGroup(buf *bytes.Buffer) (string, error) {
// 	var (
// 		file    *os.File
// 		err     error
// 		part    io.Writer
// 		w, wrtr *multipart.Writer
// 	)
// 	bf := bytes.NewBuffer(nil)
// 	w = multipart.NewWriter(bf)
// 	for i := 0; i < len(fm.Message.InputMedia); i++ {
// 		file, err = os.Open(fm.Message.InputMedia[i].Media)
// 		if err == nil {
// 			err = w.WriteField("type", fm.Message.InputMedia[i].Type)
// 		}
// 		if err == nil {
// 			part, err = w.CreateFormFile("media", fm.Message.InputMedia[i].Media)
// 		}
// 		if err == nil {
// 			_, err = io.Copy(part, file)
// 		}
// 	}
// 	wrtr = multipart.NewWriter(buf)
// 	err = wrtr.WriteField("chat_id", fmt.Sprint(fm.Message.ChatID))
// 	if err == nil {
// 		fmt.Println(bf.String())
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Println()
// 		err = wrtr.WriteField("media", bf.String())
// 	}
// 	if err == nil {
// 		err = w.Close()
// 	}
// 	if err == nil {
// 		err = wrtr.Close()
// 	}

// 	return wrtr.FormDataContentType(), err
// }

func (fm *Formatter) checkInputMediaType() ([]handleInputMedia, error) {
	var err error

	acceptedInputTypes := make(map[string]handleInputMedia, 4)
	acceptedInputTypes["photo"] = &photo{fm: fm}
	acceptedInputTypes["video"] = &video{fm: fm}
	acceptedInputTypes["audio"] = &audio{fm: fm}
	acceptedInputTypes["document"] = &document{fm: fm}

	media := make([]handleInputMedia, len(fm.mediatype))
	ok := true
	i := 0
	for i < len(fm.mediatype) && ok {
		media[i], ok = acceptedInputTypes[fm.mediatype[i]]
		i++
	}

	if !ok {
		err = fmt.Errorf("the type of media file you've chosen [%s] isn't supported yet. Please choose something among 'photo', 'video', 'document' or 'audio'", fm.mediatype[i])
	}
	return media, err
}

func (fm *Formatter) checkMediaType() ([]*mediaFields, error) {
	var err error

	acceptedTypes := make(map[string]*mediaFields, 7)
	acceptedTypes["photo"] = &mediaFields{data: &photo{fm}, method: photoMethod}
	acceptedTypes["audio"] = &mediaFields{data: &audio{fm}, method: audipMethod}
	acceptedTypes["document"] = &mediaFields{data: &document{fm}, method: documentMethod}
	acceptedTypes["video"] = &mediaFields{data: &video{fm}, method: videoMethod}
	acceptedTypes["animation"] = &mediaFields{data: &animation{fm}, method: animationMethod}
	acceptedTypes["voice"] = &mediaFields{data: &voice{fm}, method: voiceMethod}
	acceptedTypes["video-note"] = &mediaFields{data: &videoNote{fm}, method: videoNoteMethod}

	media := make([]*mediaFields, len(fm.mediatype))
	found := true
	i := 0

	for i < len(fm.mediatype) && found {
		media[i], found = acceptedTypes[fm.mediatype[i]]
		i++
	}

	if found {
		if len(fm.mediatype) == 1 {
			if fm.mediatype[0] == "photo" {
				fm.Media.Photo = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "audio" {
				fm.Media.Audio = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "document" {
				fm.Media.Document = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "video" {
				fm.Media.Video = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "animation" {
				fm.Media.Animation = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "voice" {
				fm.Media.Voice = fm.Media.mediaFile.mediaData
			} else if fm.mediatype[0] == "video-note" {
				fm.Media.VideoNote = fm.Media.mediaFile.mediaData
			}
		}
	} else {
		err = fmt.Errorf("the type of media [%s] isn't supported yet. Please make sure you choose a right type", fm.mediatype[i])
	}
	return media, err
}

func writeToFile(file *os.File, writer *multipart.Writer, mediatype, path string) error {
	defer file.Close()

	// // Создаем заголовки вручную
	// h := make(textproto.MIMEHeader)
	// h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, mediatype, path))
	// h.Set("Content-Type", "multipart/form-data") // Устанавливаем нужный Content-Type

	// // Создаем часть с этими заголовками
	// part, err := writer.CreatePart(h)
	// if err == nil {
	// 	// Копируем файл в часть
	// 	_, err = io.Copy(part, file)
	// }

	part, err := writer.CreateFormFile(mediatype, path)
	if err == nil {
		_, err = io.Copy(part, file)
	}
	return err
}

func (fm *Formatter) durationWidthHeight(writer *multipart.Writer) error {
	var err error
	if fm.Char.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.Char.Duration))
	}
	if err == nil && fm.Media.Width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", fm.Media.Width))
	}
	if err == nil && fm.Media.Height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", fm.Media.Height))
	}
	return err
}

func (fm *Formatter) showCapAbMediaHasSpoiler(writer *multipart.Writer) error {
	var err error
	if fm.Message.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")
	}
	if err == nil && fm.Char.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")
	}
	return err
}

func (fm *Formatter) thumbnail(file *os.File, writer *multipart.Writer) error {
	var (
		err  error
		part io.Writer
	)
	if fm.Media.Thumbnail.mediaData != "" {
		part, err = writer.CreateFormFile("thumbnail", fm.Media.Thumbnail.mediaData)
		if err == nil {
			_, err = io.Copy(part, file)
		}
	}
	return err
}

func (photo *photo) uniqueFields() (string, error) {
	err := writeToFile(photo.fm.file, photo.fm.writer, "photo", photo.fm.path)
	if err == nil {
		err = photo.fm.showCapAbMediaHasSpoiler(photo.fm.writer)
	}
	return photo.fm.writer.FormDataContentType(), err
}

func (photo *photo) uniqueInputFields(writer *multipart.Writer, numberInQueue int) error {
	file, err := os.Open(photo.fm.Media.Group[numberInQueue].Media)
	if err == nil {
		err = writer.WriteField("type", "photo")
	}
	if err == nil {
		path := fmt.Sprintf("attach://%s", photo.fm.Media.Group[numberInQueue].Media)
		err = writeToFile(file, writer, "media", path)
	}
	if err == nil && photo.fm.Message.Text != "" {
		err = writer.WriteField("caption", photo.fm.Message.Text)
	}
	if err == nil && photo.fm.Message.ParseMode != "" {
		err = writer.WriteField("parse_mode", photo.fm.Message.ParseMode)
	}
	if err == nil && len(photo.fm.Message.CaptionEntities) > 0 {
		body, err1 := json.Marshal(photo.fm.Message.CaptionEntities)
		if err1 == nil {
			err = writer.WriteField("caption_entities", string(body))
		}
	}
	if err == nil {
		err = photo.fm.showCapAbMediaHasSpoiler(writer)
	}
	return err
}

func (audio *audio) uniqueFields() (string, error) {
	err := writeToFile(audio.fm.file, audio.fm.writer, "audio", audio.fm.path)
	if err == nil && audio.fm.Char.Duration != 0 {
		err = audio.fm.writer.WriteField("duration", fmt.Sprintf("%d", audio.fm.Char.Duration))
	}
	if err == nil && audio.fm.Char.Performer != "" {
		err = audio.fm.writer.WriteField("performer", audio.fm.Char.Performer)
	}
	if err == nil {
		err = audio.fm.thumbnail(audio.fm.file, audio.fm.writer)
	}
	return audio.fm.writer.FormDataContentType(), err
}

func (audio *audio) uniqueInputFields(writer *multipart.Writer, numberInQueue int) error {
	return nil
}

func (doc *document) uniqueFields() (string, error) {
	err := writeToFile(doc.fm.file, doc.fm.writer, "document", doc.fm.path)
	if err == nil {
		err = doc.fm.thumbnail(doc.fm.file, doc.fm.writer)
	}
	if err == nil && doc.fm.Media.DisableContentType {
		err = doc.fm.writer.WriteField("disable_content_type_detection", "True")
	}
	return doc.fm.writer.FormDataContentType(), err
}

func (doc *document) uniqueInputFields(writer *multipart.Writer, numberInQueue int) error {
	return nil
}

func (video *video) uniqueFields() (string, error) {
	err := writeToFile(video.fm.file, video.fm.writer, "video", video.fm.path)
	if err == nil {
		err = video.fm.durationWidthHeight(video.fm.writer)
	}
	if err == nil {
		err = video.fm.thumbnail(video.fm.file, video.fm.writer)
	}
	if err == nil {
		err = video.fm.showCapAbMediaHasSpoiler(video.fm.writer)
	}
	if err == nil && video.fm.Media.SupportsStreaming {
		err = video.fm.writer.WriteField("supports_streaming", "True")
	}
	return video.fm.writer.FormDataContentType(), err
}

func (video *video) uniqueInputFields(writer *multipart.Writer, numberInQueue int) error {
	return nil
}

func (anim *animation) uniqueFields() (string, error) {
	err := writeToFile(anim.fm.file, anim.fm.writer, "animation", anim.fm.path)
	if err == nil {
		err = anim.fm.durationWidthHeight(anim.fm.writer)
	}
	if err == nil {
		err = anim.fm.thumbnail(anim.fm.file, anim.fm.writer)
	}
	if err == nil {
		err = anim.fm.showCapAbMediaHasSpoiler(anim.fm.writer)
	}
	return anim.fm.writer.FormDataContentType(), err
}

func (voice *voice) uniqueFields() (string, error) {
	err := writeToFile(voice.fm.file, voice.fm.writer, "voice", voice.fm.path)
	if err == nil && voice.fm.Char.Duration != 0 {
		err = voice.fm.writer.WriteField("duration", fmt.Sprintf("%d", voice.fm.Char.Duration))
	}
	return voice.fm.writer.FormDataContentType(), err
}

func (vn *videoNote) uniqueFields() (string, error) {
	err := writeToFile(vn.fm.file, vn.fm.writer, "video_note", vn.fm.path)
	if err == nil && vn.fm.Char.Duration != 0 {
		err = vn.fm.writer.WriteField("duration", fmt.Sprintf("%d", vn.fm.Char.Duration))
	}
	if err == nil && vn.fm.Media.Length != 0 {
		err = vn.fm.writer.WriteField("length", fmt.Sprintf("%d", vn.fm.Media.Length))
	}
	if err == nil {
		err = vn.fm.thumbnail(vn.fm.file, vn.fm.writer)
	}
	return vn.fm.writer.FormDataContentType(), err
}

func (fm *Formatter) commonMediaFields() error {
	var err error
	if fm.Business.BusinessConnectionID != "" {
		err = fm.writer.WriteField("business_connection_id", fm.Business.BusinessConnectionID)
	}
	if err == nil {
		err = fm.writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Chat.ChatID))
	}
	if err == nil && fm.Message.ThreadID != "" {
		err = fm.writer.WriteField("message_thread_id", fm.Message.ThreadID)
	}
	if err == nil && fm.Message.Text != "" && len(fm.kindofmedia) == 1 {
		err = fm.writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil && fm.Message.ParseMode != "" && len(fm.kindofmedia) == 1 {
		err = fm.writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil && len(fm.Message.CaptionEntities) > 0 && len(fm.kindofmedia) == 1 {
		body, err1 := json.Marshal(fm.Message.CaptionEntities)
		if err1 == nil {
			err = fm.writer.WriteField("caption_entities", string(body))
		}
	}
	if err == nil && fm.Char.DisableNotification {
		err = fm.writer.WriteField("disable_notification", "True")
	}
	if err == nil && fm.Char.ProtectContent {
		err = fm.writer.WriteField("protect_content", "True")
	}
	if err == nil && fm.Message.EffectID != "" {
		err = fm.writer.WriteField("message_effect_id", fm.Message.EffectID)
	}
	if err == nil && fm.Chat.ReplyParameters.MessageID != 0 {
		body, err1 := json.Marshal(fm.Chat.ReplyParameters)
		if err1 == nil {
			err = fm.writer.WriteField("reply_parameters", string(body))
		}
	}
	if err == nil && fm.Message.ReplyMarkup != nil && len(fm.kindofmedia) == 1 {
		body, err1 := json.Marshal(fm.Message.ReplyMarkup)
		if err1 == nil {
			err = fm.writer.WriteField("reply_markup", string(body))
		}
	}
	return err
}

func (fm *Formatter) commonInputMediaFields(buf *bytes.Buffer) error {
	var (
		err  error
		part io.Writer
	)
	if fm.Business.BusinessConnectionID != "" {
		err = fm.writer.WriteField("business_connection_id", fm.Business.BusinessConnectionID)
	}
	if err == nil {
		err = fm.writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Chat.ChatID))
	}
	if err == nil && fm.Message.ThreadID != "" {
		err = fm.writer.WriteField("message_thread_id", fm.Message.ThreadID)
	}
	if err == nil {
		part, err = fm.writer.CreateFormField("media")
	}
	if err == nil {
		// log.Print(part)
		_, err = io.Copy(part, buf)
	}
	return err
}

func (fm *Formatter) fromStorageMedia(buf *bytes.Buffer) (string, string, error) {
	fm.file = new(os.File)
	fm.writer = multipart.NewWriter(buf)
	media, err := fm.checkMediaType()
	if err == nil {
		fm.file, err = os.Open(fm.path)
	}
	if err == nil {
		fm.contenttype, err = media[0].data.uniqueFields()
	}
	if err == nil {
		err = fm.commonMediaFields()
	}
	if err == nil {
		defer fm.writer.Close()
	}
	return media[0].method, fm.writer.FormDataContentType(), err
}

func (fm *Formatter) makeMediaRequest(buf *bytes.Buffer) {
	var (
		messageByte, chatByte, mediaByte, businessByte, charByte, finalbody []byte
		err                                                                 error
	)
	messageMap := make(map[string]interface{})
	chatMap := make(map[string]interface{})
	mediaMap := make(map[string]interface{})
	businessMap := make(map[string]interface{})
	charMap := make(map[string]interface{})
	finalmap := make(map[string]interface{})
	commands := []interface{}{&fm.Message, &fm.Chat, &fm.Media, &fm.Business, &fm.Char}
	storageArr := [][]byte{messageByte, chatByte, mediaByte, businessByte, charByte}
	storageMaps := []map[string]interface{}{messageMap, chatMap, mediaMap, businessMap, charMap}
	if fm.Message.Text != "" {
		fm.Message.Caption = fm.Message.Text
	}
	for i := 0; i < len(commands) && err == nil; i++ {
		storageArr[i], err = json.Marshal(commands[i])
	}
	for i := 0; i < len(storageArr) && err == nil; i++ {
		err = json.Unmarshal(storageArr[i], &storageMaps[i])
	}
	if err == nil {
		for k, v := range messageMap {
			finalmap[k] = v
		}
		for k, v := range chatMap {
			finalmap[k] = v
		}
		for k, v := range mediaMap {
			finalmap[k] = v
		}
		for k, v := range businessMap {
			finalmap[k] = v
		}
		for k, v := range charMap {
			finalmap[k] = v
		}

		finalbody, err = json.Marshal(&finalmap)
		if err == nil {
			buf.Write(finalbody)
		}
		log.Print(string(finalbody))
	}
}

func (fm *Formatter) tgOrURLMedia(buf *bytes.Buffer) (string, string, error) {
	media, err := fm.checkMediaType()
	if err == nil {
		fm.makeMediaRequest(buf)
	}
	return media[0].method, "application/json", err
}

func (fm *Formatter) inputMedia(mediawriter *multipart.Writer, fields handleInputMedia, numberInQueue int) error {
	err := fields.uniqueInputFields(mediawriter, numberInQueue)
	return err
}

func (fm *Formatter) mediaGroup(buf *bytes.Buffer) (string, string, error) {
	var (
		found       bool
		err         error
		contenttype string // just or testing. For other functions I have fm.contenttype
		media       []handleInputMedia
	)

	for i := 0; i < len(fm.kindofmedia) && !found; i++ {
		if fm.kindofmedia[i] == fromStorage {
			found = true
		}
	}

	if !found {
		_, err := fm.checkInputMediaType()
		if err == nil {
			fm.makeMediaRequest(buf)
			contenttype = "application/json"
		}
	} else {
		fm.writer = multipart.NewWriter(buf)
		media, err = fm.checkInputMediaType()
		inputbuf := bytes.NewBuffer(nil)
		mediawriter := multipart.NewWriter(inputbuf)
		for i := 0; i < len(media) && err == nil; i++ {
			err = fm.inputMedia(mediawriter, media[i], i)
		}
		if err == nil {
			err = mediawriter.Close()
		}
		if err == nil {
			err = fm.commonInputMediaFields(inputbuf)
		}
		if err == nil {
			err = fm.writer.Close()
		}
		if err == nil {
			contenttype = fm.writer.FormDataContentType()
		}
	}
	log.Print(buf.String())
	return mediaGroup, contenttype, err
}
