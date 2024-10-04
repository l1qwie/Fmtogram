package media

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

func (ph *Photo) defaultString(data string) bool {
	return data == ""
}

func (ph *Photo) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Photo) defautBool(data bool) bool {
	return !data
}

func (ph *Photo) defaultInt(data int) bool {
	return data == 0
}

func (ph *Video) defaultString(data string) bool {
	return data == ""
}

func (ph *Video) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Video) defautBool(data bool) bool {
	return !data
}

func (ph *Video) defaultInt(data int) bool {
	return data == 0
}

func (ph *Audio) defaultString(data string) bool {
	return data == ""
}

func (ph *Audio) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Audio) defautBool(data bool) bool {
	return !data
}

func (ph *Audio) defaultInt(data int) bool {
	return data == 0
}

func (ph *Document) defaultString(data string) bool {
	return data == ""
}

func (ph *Document) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Document) defautBool(data bool) bool {
	return !data
}

func (ph *Document) defaultInt(data int) bool {
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

func RequiredPhotoData(ph *Photo, num int) error {
	var err error
	if ph.Photo == "" {
		err = errors.MissedRequiredField(objectPhoto, "Photo{ID/FilePath/URL}", num, true)
	}
	if err == nil && ph.GottenFrom == 0 {
		err = errors.MissedRequiredField(objectPhoto, "GottenFrom{media.TG, media.Storage, media.URL}", num, true)
	}
	return err
}

func RequiredVideoData(vd *Video, num int) error {
	var err error
	if vd.Video == "" {
		err = errors.MissedRequiredField(objectVideo, "Video{ID/FilePath/URL}", num, true)
	}
	if err == nil && vd.VideoGottenFrom == 0 {
		err = errors.MissedRequiredField(objectVideo, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (vd.Thumbnail != "" && vd.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(objectVideo, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}

func RequiredAudioData(ad *Audio, num int) error {
	var err error
	if ad.Audio == "" {
		err = errors.MissedRequiredField(objectAudio, "Audio{ID/FilePath/URL}", num, true)
	}
	if err == nil && ad.AudioGottenFrom == 0 {
		err = errors.MissedRequiredField(objectAudio, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (ad.Thumbnail != "" && ad.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(objectAudio, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}

func RequiredDocumentData(dc *Document, num int) error {
	var err error
	if dc.Document == "" {
		err = errors.MissedRequiredField(objectDocument, "Document{ID/FilePath/URL}", num, true)
	}
	if err == nil && dc.DocumentGottenFrom == 0 {
		err = errors.MissedRequiredField(objectDocument, "GottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	if err == nil && (dc.Thumbnail != "" && dc.ThumbnailGottenFrom == 0) {
		err = errors.MissedRequiredField(objectDocument, "ThumbnailGottenFrom{media.Telegram, media.Storage, media.Internet}", num, true)
	}
	return err
}
