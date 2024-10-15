package formatter

import (
	"encoding/json"
	"fmt"
)

func uniqueMedia(msg *Message) error {
	var err error
	var mediajson []byte
	if msg.fm.mh.evenone {
		err = msg.fm.mh.storage[msg.fm.mh.i-1].multipartFields(msg.fm.writer, nil, 0, false)
		msg.fm.contenttype = msg.fm.writer.FormDataContentType()
	} else {
		mediajson, err = msg.fm.mh.storage[msg.fm.mh.i-1].jsonFileds()
		if err == nil {
			msg.fm.buf.Write(mediajson)
		}
		msg.fm.contenttype = "application/json"
	}
	return err
}

func mediaGroup(msg *Message) error {
	var (
		err    error
		jsbody []byte
	)
	group := make([]interface{}, msg.fm.mh.amount)
	mediaMap := make(map[string]interface{}, 1)
	if !msg.fm.mh.evenone {
		for i := 0; i < len(msg.fm.mh.storage); i++ {

			if msg.fm.mh.storage[i] != nil {
				fmt.Println(msg.fm.mh.storage[i])
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
				err = msg.fm.mh.storage[i].multipartFields(msg.fm.writer, &group, i, true)
			}
		}
		if err == nil {
			err = putGroup(msg.fm.writer, group)
		}
		msg.fm.contenttype = msg.fm.writer.FormDataContentType()
	}
	return err
}
