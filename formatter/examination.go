package formatter

import (
	"log"

	"github.com/l1qwie/Fmtogram/errors"
)

type checking interface {
	defaultString(string) bool
	defaultArray([]interface{}) bool
	defautBool(bool) bool
	defaultInt(int) bool
}

func (ph *photo) defaultString(data string) bool {
	return data == ""
}

func (ph *photo) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *photo) defautBool(data bool) bool {
	return !data
}

func (ph *photo) defaultInt(data int) bool {
	return data == 0
}

func (ph *video) defaultString(data string) bool {
	return data == ""
}

func (ph *video) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *video) defautBool(data bool) bool {
	return !data
}

func (ph *video) defaultInt(data int) bool {
	return data == 0
}

func (ph *audio) defaultString(data string) bool {
	return data == ""
}

func (ph *audio) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *audio) defautBool(data bool) bool {
	return !data
}

func (ph *audio) defaultInt(data int) bool {
	return data == 0
}

func (ph *document) defaultString(data string) bool {
	return data == ""
}

func (ph *document) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *document) defautBool(data bool) bool {
	return !data
}

func (ph *document) defaultInt(data int) bool {
	return data == 0
}

func isItEmply(ch checking, work int, data interface{}) bool {
	var ok bool
	switch d := data.(type) {
	case string:
		if work == checkString {
			ok = ch.defaultString(d)
		}
	case []interface{}:
		if work == checkArray {
			ok = ch.defaultArray(d)
		}
	case bool:
		if work == checkBool {
			ok = ch.defautBool(d)
		}
	case int:
		if work == checkInt {
			ok = ch.defaultInt(d)
		}
	default:
		log.Print("SOMETHING WENT WRONG!")
	}

	return ok
}

func requiredPhotoData(ph *photo, num int) error {
	var err error
	if ph.Photo == "" {
		err = errors.MissedRequiredField(interfacePhoto, "Photo{ID/FilePath/URL}", num, true)
	}
	if err == nil && ph.GottenFrom == 0 {
		err = errors.MissedRequiredField(interfacePhoto, "GottenFrom{media.TG, media.Storage, media.URL}", num, true)
	}
	return err
}

func requiredVideoData(vd *video, num int) error {
	var err error
	if vd.Video == "" {
		err = errors.MissedRequiredField(interfaceVideo, "Video{ID/FilePath/URL}", num, true)
	}
	if err == nil && vd.VideoGottenFrom == 0 {
		err = errors.MissedRequiredField(interfaceVideo, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (vd.Thumbnail != "" && vd.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(interfaceVideo, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}

func requiredAudioData(ad *audio, num int) error {
	var err error
	if ad.Audio == "" {
		err = errors.MissedRequiredField(interfaceAudio, "Audio{ID/FilePath/URL}", num, true)
	}
	if err == nil && ad.AudioGottenFrom == 0 {
		err = errors.MissedRequiredField(interfaceAudio, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (ad.Thumbnail != "" && ad.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(interfaceAudio, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}

func requiredDocumentData(dc *document, num int) error {
	var err error
	if dc.Document == "" {
		err = errors.MissedRequiredField(interfaceDocument, "Document{ID/FilePath/URL}", num, true)
	}
	if err == nil && dc.DocumentGottenFrom == 0 {
		err = errors.MissedRequiredField(interfaceDocument, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (dc.Thumbnail != "" && dc.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(interfaceDocument, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}
