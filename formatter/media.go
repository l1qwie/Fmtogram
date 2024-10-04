package formatter

import (
	"encoding/json"

	"github.com/l1qwie/Fmtogram/formatter/media"
)

func uniqueMedia(msg *Message) error {
	var err error
	var mediajson []byte
	if !msg.fm.mh.evenone {
		err = msg.fm.mh.storage[msg.fm.mh.i-1].MultipartFields(msg.fm.writer, nil, 0, false)
	} else {
		mediajson, err = msg.fm.mh.storage[msg.fm.mh.i-1].JsonFileds()
		if err == nil {
			msg.fm.buf.Write(mediajson)
		}
	}
	return err
}

func mediaGroup(msg *Message) error {
	var (
		err    error
		jsbody []byte
		group  [len(msg.fm.mh.storage)]interface{}
	)
	mediaMap := make(map[string]interface{}, 1)
	if !msg.fm.mh.evenone {
		for i := 0; i < len(msg.fm.mh.storage) && err == nil; i++ {

			if msg.fm.mh.storage[i] != nil {
				group[i] = msg.fm.mh.storage[i]
			}
		}

		mediaMap["media"] = group
		jsbody, err = json.Marshal(mediaMap)
		if err == nil {
			msg.fm.buf.Write(jsbody)
		}
		msg.fm.contenttype = "application/json"
	} else {

		for i := 0; i < len(msg.fm.mh.storage) && err == nil; i++ {
			if msg.fm.mh.storage[i] != nil {
				err = msg.fm.mh.storage[i].MultipartFields(msg.fm.writer, &group, i, true)
			}
		}
		if err == nil {
			err = media.Group(msg.fm.writer, group)
		}
		msg.fm.contenttype = msg.fm.writer.FormDataContentType()
	}
	return err
}
