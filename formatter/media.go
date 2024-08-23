package formatter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

func (fm *Formatter) AddPhotoFromStorage(path string) {
	logs.DataWrittenSuccessfully("A Photo From The Storage")
	fm.Message.Photo = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromTG(path string) {
	logs.DataWrittenSuccessfully("A Photo Telegram ID")
	fm.Message.Photo = path
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromInternet(path string) {
	logs.DataWrittenSuccessfully("A Photo Url From The Internet")
	fm.Message.Photo = path
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"photo"}
}
func (fm *Formatter) AddVideoFromStorage(path string) {
	logs.DataWrittenSuccessfully("A Video From The Storage")
	fm.Message.Video = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromTG(path string) {
	logs.DataWrittenSuccessfully("A Video Telegram ID")
	fm.Message.Video = path
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromInternet(path string) {
	logs.DataWrittenSuccessfully("A Video Url From The Internet")
	fm.Message.Video = path
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddMapOfMedia(arr []types.Media) {
	i := 0
	fm.Message.InputMedia = make([]types.Media, len(arr))
	for _, val := range arr {
		fm.Message.InputMedia[i].Type = val.Type
		fm.Message.InputMedia[i].Media = val.Media
		i++
	}
}

func (fm *Formatter) createMediaGroup(buf *bytes.Buffer) (string, error) {
	var (
		file    *os.File
		err     error
		part    io.Writer
		w, wrtr *multipart.Writer
	)
	bf := bytes.NewBuffer(nil)
	w = multipart.NewWriter(bf)
	for i := 0; i < len(fm.Message.InputMedia); i++ {
		file, err = os.Open(fm.Message.InputMedia[i].Media)
		if err == nil {
			err = w.WriteField("type", fm.Message.InputMedia[i].Type)
		}
		if err == nil {
			part, err = w.CreateFormFile("media", fm.Message.InputMedia[i].Media)
		}
		if err == nil {
			_, err = io.Copy(part, file)
		}
	}
	wrtr = multipart.NewWriter(buf)
	err = wrtr.WriteField("chat_id", fmt.Sprint(fm.Message.ChatID))
	if err == nil {
		fmt.Println(bf.String())
		fmt.Println()
		fmt.Println()
		fmt.Println()
		err = wrtr.WriteField("media", bf.String())
	}
	if err == nil {
		err = w.Close()
	}
	if err == nil {
		err = wrtr.Close()
	}

	return wrtr.FormDataContentType(), err
}

func (fm *Formatter) prepareMedia(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
	)

	writer = multipart.NewWriter(buf)
	if fm.mediatype[0] == "photo" {
		media = fm.Message.Photo
	} else if fm.mediatype[0] == "video" {
		media = fm.Message.Video
	}

	file, err = os.Open(media)
	if err == nil {
		part, err = writer.CreateFormFile(fm.mediatype[0], media)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Message.ChatID))
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.WriteField("reply_markup", fm.Message.ReplyMarkup)
	}
	if err == nil {
		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()
	return writer.FormDataContentType(), err
}

func (fm *Formatter) PrepareMediaForEdit(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
		media  string
	)
	writer = multipart.NewWriter(buf)
	if fm.mediatype[0] == "photo" {
		media = fm.Message.Photo
	} else if fm.mediatype[0] == "video" {
		media = fm.Message.Video
	}

	file, err = os.Open(media)
	if err == nil {
		part, err = writer.CreateFormFile(fm.mediatype[0], media)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()
	return writer.FormDataContentType(), err
}
