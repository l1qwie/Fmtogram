package media

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	Storage          int    = 1
	Telegram         int    = 2
	Internet         int    = 3
	photoMethod      string = "sendPhoto"
	audioMethod      string = "sendAudio"
	documentMethod   string = "sendDocument"
	videoMethod      string = "sendVideo"
	animationMethod  string = "sendAnimation"
	voiceMethod      string = "sendVoice"
	videoNoteMethod  string = "sendVideoNote"
	mediaGroupMethod string = "sendMediaGroup"
)

func writeFileToMultipart(writer *multipart.Writer, fieldname, filename string) error {
	var part io.Writer
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()

		part, err = writer.CreateFormFile(fieldname, filename)

		if err == nil {
			_, err = io.Copy(part, file)
			if err != nil {
				errors.CantCopyFile(err)
			}
		} else {
			errors.CantCreateFormFile(err)
		}
	} else {
		errors.CantOpenFile(err)
	}
	return err
}

func addThumbnail(writer *multipart.Writer, thumbnail string, gottenfrom int) error {
	var err error
	if gottenfrom == Storage {
		err = writeFileToMultipart(writer, "thumbnail", thumbnail)
	} else {
		err = writer.WriteField("thumbnail", thumbnail)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	return err
}

func commonFields(writer *multipart.Writer, caption, parsemode string, captionEntities []*types.MessageEntity) error {
	var err error
	if caption != "" {
		err = writer.WriteField("caption", caption)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && parsemode != "" {
		err = writer.WriteField("parse_mode", parsemode)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && len(captionEntities) != 0 {
		body, err1 := json.Marshal(captionEntities)
		if err1 == nil {
			err = writer.WriteField("caption_entities", string(body))

			if err != nil {
				errors.CantWriteField(err)
			}

		} else {
			errors.CantMarshalJSON(err1)
		}
	}
	return err
}

func (ph *Photo) CreateFields(writer *multipart.Writer) (string, error) {
	err := commonFields(writer, ph.Caption, ph.ParseMode, ph.CaptionEntities)
	if err == nil {
		err = writeFileToMultipart(writer, "photo", ph.Photo)
	}
	if err == nil && ph.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ph.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	return photoMethod, err
}

func (vd *Video) CreateFields(writer *multipart.Writer) (string, error) {
	err := commonFields(writer, vd.Caption, vd.ParseMode, vd.CaptionEntities)
	if err == nil {
		err = writeFileToMultipart(writer, "video", vd.Video)
	}
	if err == nil && vd.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", vd.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", vd.Width))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", vd.Height))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.Thumbnail != "" {
		err = addThumbnail(writer, vd.Thumbnail, vd.ThumbnailGottenFrom)
	}
	if err == nil && vd.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && vd.SupportsStreaming {
		err = writer.WriteField("supports_streaming", "True")
	}
	return videoMethod, err
}

func (ad *Audio) CreateFields(writer *multipart.Writer) (string, error) {
	err := commonFields(writer, ad.Caption, ad.ParseMode, ad.CaptionEntities)
	if err == nil {
		err = writeFileToMultipart(writer, "audio", ad.Audio)
	}
	if err == nil && ad.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", ad.Duration))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ad.Performer != "" {
		err = writer.WriteField("performer", ad.Performer)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && ad.Thumbnail != "" {
		err = addThumbnail(writer, ad.Thumbnail, ad.ThumbnailGottenFrom)
	}
	return audioMethod, err
}

func (dc *Document) CreateFields(writer *multipart.Writer) (string, error) {
	err := commonFields(writer, dc.Caption, dc.ParseMode, dc.CaptionEntities)
	if err == nil {
		err = writeFileToMultipart(writer, "document", dc.Document)
	}
	if err == nil && dc.Thumbnail != "" {
		err = addThumbnail(writer, dc.Thumbnail, dc.ThumbnailGottenFrom)
	}
	if err == nil && dc.DisableContentTypeDetection {
		err = writer.WriteField("disable_content_type_detection", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	return documentMethod, err
}
