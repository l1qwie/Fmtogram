package helper

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/types"
)

func ReturnText(Telegram *types.Telegram) (text string) {
	if Telegram.Result[0].Message.Text != "" {
		text = Telegram.Result[0].Message.Text
	} else if Telegram.Result[0].CallbackQuery.Data != "" {
		text = Telegram.Result[0].CallbackQuery.Data
	}
	return text
}

func ReturnChatId(Telegram *types.Telegram) (chatID int) {
	if Telegram.Result[0].Message.TypeFrom.UserID != 0 {
		chatID = Telegram.Result[0].Message.TypeFrom.UserID
	} else if Telegram.Result[0].CallbackQuery.TypeFrom.UserID != 0 {
		chatID = Telegram.Result[0].CallbackQuery.TypeFrom.UserID
	}
	return chatID
}

func ReturnName(Telegram *types.Telegram) (name string) {
	if Telegram.Result[0].Message.TypeFrom.Name != "" {
		name = Telegram.Result[0].Message.TypeFrom.Name
	} else if Telegram.Result[0].CallbackQuery.TypeFrom.Name != "" {
		name = Telegram.Result[0].CallbackQuery.TypeFrom.Name
	}
	return name
}

func ReturnLastName(Telegram *types.Telegram) (lastname string) {
	if Telegram.Result[0].Message.TypeFrom.LastName != "" {
		lastname = Telegram.Result[0].Message.TypeFrom.LastName
	} else if Telegram.Result[0].CallbackQuery.TypeFrom.LastName != "" {
		lastname = Telegram.Result[0].CallbackQuery.TypeFrom.LastName
	}
	return lastname
}

func ReturnUsername(Telegram *types.Telegram) (username string) {
	if Telegram.Result[0].Message.TypeFrom.Username != "" {
		username = Telegram.Result[0].Message.TypeFrom.Username
	} else if Telegram.Result[0].CallbackQuery.TypeFrom.Username != "" {
		username = Telegram.Result[0].CallbackQuery.TypeFrom.Username
	}
	return username
}

func ReturnPhone(tr *types.Telegram) (phone string) {
	if tr.Result[0].Message.TypeFrom.Phone != "" {
		phone = tr.Result[0].Message.TypeFrom.Phone
	} else if tr.Result[0].CallbackQuery.TypeFrom.Phone != "" {
		phone = tr.Result[0].CallbackQuery.TypeFrom.Phone
	}
	return phone
}

func ReturnLanguage(Telegram *types.Telegram) (language string) {
	if Telegram.Result[0].Message.TypeFrom.Language != "" {
		language = Telegram.Result[0].Message.TypeFrom.Language
	} else if Telegram.Result[0].CallbackQuery.TypeFrom.Language != "" {
		language = Telegram.Result[0].CallbackQuery.TypeFrom.Language
	}
	return language
}

func ReturnBotStatus(Telegram *types.Telegram) (botstatus bool) {

	m_isbot := Telegram.Result[0].Message.TypeFrom.IsBot
	cl_isbot := Telegram.Result[0].CallbackQuery.TypeFrom.IsBot

	if !m_isbot && !cl_isbot {
		botstatus = false
	} else if m_isbot && !cl_isbot || !m_isbot && cl_isbot {
		botstatus = true
	}
	return botstatus
}

func ReturnMessageId(message *types.Returned) (messageId int, err error) {
	if message.Ok {
		messageId = message.Result.MessageId
	} else {
		err = fmt.Errorf("we don't have a message id")
	}
	return messageId, err
}

func ReturnPhotosFileIdout(message *types.Returned) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(message.Result.Photo))
	if len(message.Result.Photo) > 0 {
		for i := 0; i < len(message.Result.Photo); i++ {
			fileIds[i] = message.Result.Photo[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileIds, err
}

func ReturnVideosFileIdout(message *types.Returned) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(message.Result.Video))
	if len(message.Result.Video) > 0 {
		for i := 0; i < len(message.Result.Video); i++ {
			fileIds[i] = message.Result.Video[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileIds, err
}

func ReturnPhotosFileIdfrom(tr *types.Telegram) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	j := 0
	if len(tr.Result[0].Message.Photo) > 0 {
		l := len(tr.Result[0].Message.Photo) / 4
		fileIds = make([]string, l)
		for i := 0; i < l; i++ {
			fileIds[i] = tr.Result[0].Message.Photo[j].FileId
			j = j + 4
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileIds, err
}

func ReturnVideosFileIdfrom(tr *types.Telegram) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(tr.Result[0].Message.Video))
	if len(tr.Result[0].Message.Video) > 0 {
		for i := 0; i < len(tr.Result[0].Message.Video); i++ {
			fileIds[i] = tr.Result[0].Message.Video[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileIds, err
}

func ReturnPhotoResp(resp *types.Returned) (fileid string, err error) {
	if len(resp.Result.Photo) > 0 {
		if len(resp.Result.Photo) < 5 {
			fileid = resp.Result.Photo[0].FileId
		} else {
			err = fmt.Errorf("there is not a one Photo. There are a few. You should use [ReturnMediaResp]")
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileid, err
}

func ReturnVideoResp(resp *types.Returned) (fileid string, err error) {
	if len(resp.Result.Video) > 0 {
		if len(resp.Result.Video) < 5 {
			fileid = resp.Result.Video[0].FileId
		} else {
			err = fmt.Errorf("there is not a one Video. There are a few. You should use [ReturnMediaResp]")
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileid, err
}

func ReturnMediaResp(resp *types.Returned) ([]types.Media, error) {
	var err error
	media := make([]types.Media, (len(resp.Result.Photo) + len(resp.Result.Video)))
	count := 0
	if len(resp.Result.Photo) > 0 || len(resp.Result.Video) > 0 {
		if len(resp.Result.Photo) < 4 && len(resp.Result.Video) < 4 {
			if len(resp.Result.Photo) > 1 {
				for i := range resp.Result.Photo {
					media[count].Media = resp.Result.Photo[i].FileId
					media[count].Type = "photo"
					count++
				}
			}
			if len(resp.Result.Video) > 1 {
				for i := range resp.Result.Video {
					media[count].Media = resp.Result.Video[i].FileId
					media[count].Type = "video"
					count++
				}
			}
		} else {
			err = fmt.Errorf("there are not a few Video or Photo. There is only one (maybe bouth of them photo+video). You should use [ReturnPhotoResp] or [ReturnVideoResp]")
		}
	} else {
		err = fmt.Errorf("we don't have any Videos or Photos")
	}
	return media, err
}

func ReturnPhotoReq(req *types.Telegram) (fileId string, err error) {
	if len(req.Result[0].Message.Photo) > 0 {
		l := len(req.Result[0].Message.Photo) / 4
		if l == 1 {
			fileId = req.Result[0].Message.Photo[0].FileId
		} else {
			err = fmt.Errorf("there is not a one Photo. There are a few. You should use [ReturnMediaReq]")
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileId, err
}

func ReturnVideoReq(req *types.Telegram) (fileId string, err error) {
	if len(req.Result[0].Message.Video) > 0 {
		l := len(req.Result[0].Message.Video) / 4
		if l == 1 {
			fileId = req.Result[0].Message.Video[0].FileId
		} else {
			err = fmt.Errorf("there is not a one Video. There are a few. You should use [ReturnMediaReq]")
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileId, err
}

func ReturnMediaReq(req *types.Telegram) ([]types.Media, error) {
	var err error
	media := make([]types.Media, (len(req.Result[0].Message.Photo) + len(req.Result[0].Message.Video)))
	count := 0
	if len(req.Result[0].Message.Photo) > 0 || len(req.Result[0].Message.Video) > 0 {
		if len(req.Result[0].Message.Photo) < 2 && len(req.Result[0].Message.Video) < 2 {
			if len(req.Result[0].Message.Photo) > 1 {
				for i := range req.Result[0].Message.Photo {
					media[count].Media = req.Result[0].Message.Photo[i].FileId
					media[count].Type = "photo"
					count++
				}
			}
			if len(req.Result[0].Message.Video) > 1 {
				for i := range req.Result[0].Message.Video {
					media[count].Media = req.Result[0].Message.Video[i].FileId
					media[count].Type = "video"
					count++
				}
			}
		} else {
			err = fmt.Errorf("there are not a few Video or Photo. There is only one (maybe bouth of them photo+video). You should use [ReturnPhotoResp] or [ReturnVideoResp]")
		}
	} else {
		err = fmt.Errorf("we don't have any Videos or Photos")
	}
	return media, err
}

func ReturnTypeOfChat(req *types.Telegram) string {
	var res string
	if req.Result[0].CallbackQuery.Message.Chat.Type != "" {
		res = req.Result[0].CallbackQuery.Message.Chat.Type
	} else if req.Result[0].Message.Chat.Type != "" {
		res = req.Result[0].Message.Chat.Type
	}
	return res
}
