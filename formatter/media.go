package formatter

import (
	"bytes"
	"encoding/json"
	"log"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
	"github.com/l1qwie/Fmtogram/formatter/methods"
)

func (fm *Formatter) multipartMedia(buf *bytes.Buffer) (string, string, error) {
	fm.file = new(os.File)
	fm.writer = multipart.NewWriter(buf)
	method, err := fm.mediaStorage[0].MultipartFields(fm.writer, nil, 0, false)
	if err == nil {
		err = fm.m.MultipartFields(fm.writer, method)
	}
	if err == nil {
		err = fm.ch.MultipartFields(fm.writer)
	}
	if err == nil {
		err = fm.kb.MultipartFields(fm.writer)
	}
	if err == nil {
		err = fm.writer.Close()
	}
	// log.Print(buf.String())
	return method, fm.writer.FormDataContentType(), err
}

func (fm *Formatter) jsonMedia(buf *bytes.Buffer) (string, string, error) {
	var mesjson, chjson, kbjson, jsonreq []byte
	finalmap := make(map[string]interface{})
	method, mediajson, err := fm.mediaStorage[0].JsonFileds()
	if err == nil {
		fm.m.Caption = fm.m.Text
		mesjson, err = json.Marshal(fm.m)
	}
	if err == nil {
		chjson, err = json.Marshal(fm.ch)
	}
	if err == nil {
		kbjson, err = fm.kb.JsonFields()
		log.Print(string(kbjson))
	}
	if err == nil {
		err = json.Unmarshal(mediajson, &finalmap)
	}
	if err == nil {
		err = json.Unmarshal(mesjson, &finalmap)
	}
	if err == nil {
		err = json.Unmarshal(chjson, &finalmap)
	}
	if err == nil {
		finalmap["reply_markup"] = fm.kb
	}
	if err == nil {
		jsonreq, err = json.Marshal(finalmap)
	}
	if err == nil {
		buf.Write(jsonreq)
	}
	log.Print(string(jsonreq))
	return method, "application/json", err
}

func (fm *Formatter) uniqueMedia(buf *bytes.Buffer) (string, string, error) {
	var method, contenttype string
	strg, err := fm.mediaStorage[0].CheckGottenFrom(0)
	if !strg && err == nil {
		strg, err = fm.mediaStorage[0].CheckThumbnailGottenFrom(0)
	}
	if err == nil {
		if strg {
			method, contenttype, err = fm.multipartMedia(buf)
		} else {
			method, contenttype, err = fm.jsonMedia(buf)
		}
	}
	return method, contenttype, err
}

func (fm *Formatter) makeMediaRequest(buf *bytes.Buffer, mediabytes [][]byte, mesbody []byte) error {
	var (
		finalbody []byte
		err       error
	)
	jsonMap := make(map[string]interface{})
	media := make([]interface{}, len(fm.mediaStorage))
	for i := 0; i < len(mediabytes); i++ {
		err = json.Unmarshal(mediabytes[i], &media[i])
	}
	if err == nil {
		err = json.Unmarshal(mesbody, &jsonMap)
	}
	if err == nil {
		jsonMap["media"] = media

		finalbody, err = json.Marshal(&jsonMap)
		if err == nil {
			buf.Write(finalbody)
		}
		log.Print(string(finalbody))
	}
	return err
}

func (fm *Formatter) mediaGroup(buf *bytes.Buffer) (string, string, error) {
	var (
		contenttype     string
		err             error
		ok, found       bool
		jsbody, mesbody []byte
	)
	for i := 0; i < len(fm.mediaStorage) && err == nil; i++ {
		ok, err = fm.mediaStorage[i].CheckGottenFrom(i)
		if !ok {
			ok, err = fm.mediaStorage[i].CheckThumbnailGottenFrom(i)
			if ok {
				found = true
			}
		} else {
			found = true
		}
	}

	if !found && err == nil {
		mediabytes := make([][]byte, len(fm.mediaStorage))

		for i := 0; i < len(fm.mediaStorage) && err == nil; i++ {
			jsbody, err = json.Marshal(fm.mediaStorage[i])
			if err == nil {
				mediabytes[i] = jsbody
			}
		}
		mesbody, err = message.CreateJSON(fm.m, methods.MediaGroup)

		fm.makeMediaRequest(buf, mediabytes, mesbody)
		contenttype = "application/json"
	} else if found && err == nil {
		fm.writer = multipart.NewWriter(buf)
		group := make([]interface{}, len(fm.mediaStorage))
		for i := 0; i < len(fm.mediaStorage) && err == nil; i++ {
			_, err = fm.mediaStorage[i].MultipartFields(fm.writer, group, i, true)
		}
		if err == nil {
			err = media.Group(fm.writer, group)
		}
		if err == nil {
			err = fm.m.MultipartFields(fm.writer, methods.MediaGroup)
		}
		if err == nil {
			err = fm.writer.Close()
		}
		log.Print(buf.String())
		contenttype = fm.writer.FormDataContentType()
	}
	return methods.MediaGroup, contenttype, err
}

// func (fm *Formatter) mediaGroup(buf *bytes.Buffer) (string, string, error) {
// 	var (
// 		found       bool
// 		err         error
// 		contenttype string // just or testing. For other functions I have fm.contenttype
// 		media       []handleInputMedia
// 	)

// 	for i := 0; i < len(fm.kindofmedia) && !found; i++ {
// 		if fm.kindofmedia[i] == fromStorage {
// 			found = true
// 		}
// 	}

// 	if !found {
// 		_, err := fm.checkInputMediaType()
// 		if err == nil {
// 			fm.makeMediaRequest(buf)
// 			contenttype = "application/json"
// 		}
// 	} else {
// 		fm.writer = multipart.NewWriter(buf)

// 		media, err = fm.checkInputMediaType()

// 		for i := 0; i < len(media) && err == nil; i++ {
// 			err = fm.inputMedia(media[i], i)
// 		}

// 		if err == nil {
// 			err = fm.commonInputMediaFields()
// 		}
// 		if err == nil {
// 			err = fm.writer.Close()
// 		}
// 		if err == nil {
// 			contenttype = fm.writer.FormDataContentType()
// 		}
// 	}
// 	log.Print(buf.String())
// 	return mediaGroup, contenttype, err
// }
