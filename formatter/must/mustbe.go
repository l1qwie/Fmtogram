package must

import (
	"encoding/json"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/formatter/methods"
)

func (mch *MediaCheck) RequiredFields(method string, request []byte) error {
	var err error
	if method == methods.MediaGroup {
		err = json.Unmarshal(request, mch)
		if err == nil {
			if len(mch.Media) == 0 && mch.ChatID == nil {
				err = errors.MustBe([]string{"Media", "Chat"},
					[]string{"Media", "ChatID"},
					[]string{"fm.New{any kind of media}()", "fm.WriteChat{ID/Name/URL}()"})
			}
		}
	}
	return err
}

func (phch *PhotoCheck) RequiredFields(method string, request []byte) error {
	var err error
	if method == methods.MediaGroup {
		err = json.Unmarshal(request, phch)
		if err == nil {
			if phch.Photo == "" && phch.ChatID == nil {
				err = errors.MustBe([]string{"Media", "Chat"},
					[]string{"Photo", "ChatID"},
					[]string{"fm.NewPhoto()", "fm.WriteChat{ID/Name/URL}()"})
			}
		}
	}
	return err
}

func (ach *AudioCheck) RequiredFields(method string, request []byte) error {
	var err error
	if method == methods.MediaGroup {
		err = json.Unmarshal(request, ach)
		if err == nil {
			if ach.Audio == "" && ach.ChatID == nil {
				err = errors.MustBe([]string{"Media", "Chat"},
					[]string{"Audio", "ChatID"},
					[]string{"fm.NewAudio()", "fm.WriteChat{ID/Name/URL}()"})
			}
		}
	}
	return err
}

func (dch *DocumentCheck) RequiredFields(method string, request []byte) error {
	var err error
	if method == methods.MediaGroup {
		err = json.Unmarshal(request, dch)
		if err == nil {
			if dch.Document == "" && dch.ChatID == nil {
				err = errors.MustBe([]string{"Media", "Chat"},
					[]string{"Document", "ChatID"},
					[]string{"fm.NewDocument()", "fm.WriteChat{ID/Name/URL}()"})
			}
		}
	}
	return err
}

func (vch *VideoCheck) RequiredFields(method string, request []byte) error {
	var err error
	if method == methods.MediaGroup {
		err = json.Unmarshal(request, vch)
		if err == nil {
			if vch.Video == "" && vch.ChatID == nil {
				err = errors.MustBe([]string{"Media", "Chat"},
					[]string{"Video", "ChatID"},
					[]string{"fm.NewVideo()", "fm.WriteChat{ID/Name/URL}()"})
			}
		}
	}
	return err
}
