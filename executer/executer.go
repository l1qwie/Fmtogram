package executer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/types"
)

const None int = -1

type Entry struct {
	UserId int
	Chu    chan *types.Telegram
	Chb    chan *types.Returned
}

type RegTable struct {
	Reg []Entry
}

func GetpostRequest(url string, Buffer *bytes.Buffer, contenttype string) (body []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
		client   *http.Client
	)
	request, err = http.NewRequest("POST", url, Buffer)
	if err == nil {
		request.Header.Set("Content-Type", contenttype)
		client = &http.Client{}
		response, err = client.Do(request)
		if err == nil {
			defer response.Body.Close()
		}
	}
	if err != nil {
		log.Printf("Ошибка при попытке отправить request к Telegram GetpostRequest(): %s", err)
	}
	return io.ReadAll(response.Body)
}

func heandlerMessage(response []byte, mes *types.Returned) (err error) {
	err = json.Unmarshal(response, &mes)
	if err != nil {
		err = nil
		tr := new(types.Telegram)
		err = json.Unmarshal(response, &tr)
		mes.Result.Photo = []types.Photo{{FileId: "TRASHID"}}
	}
	return err
}

func Send(buf *bytes.Buffer, function, contenttype string, unmarshal bool) (mes *types.Returned) {
	var (
		err  error
		url  string
		body []byte
	)
	url = fmt.Sprintf("%sbot%s/%s", types.HttpsRequest, types.BotID, function)
	body, err = GetpostRequest(url, buf, contenttype)
	if err == nil && unmarshal {
		mes = new(types.Returned)
		err = heandlerMessage(body, mes)
	} else if !unmarshal {
		mes = &types.Returned{
			Ok: false,
			Result: types.Message{
				MessageId: 0,
				Chat: types.Chat{
					Id: 0,
				},
			},
		}
	}
	if err != nil {
		log.Printf("Ошикба при попытке отправить request к Telegram Send(): %s", err)
	}
	return mes
}

func Updates(offset *int, Telegram *types.Telegram) (err error) {
	var (
		response *http.Response
		body     []byte
	)
	url := fmt.Sprintf(types.HttpsRequest+"bot%s/getUpdates?limit=1&offset=%d", types.BotID, *offset)
	response, err = http.Get(url)
	if err == nil {
		defer response.Body.Close()
		body, err = io.ReadAll(response.Body)
	}
	if err == nil {
		err = handlerTelegram(body, Telegram)
	}
	return err
}

func firstMisstake(response []byte) (string, bool) {
	var (
		monitor bool
		mes     string
	)
	dresp := new(types.BadResponse)
	err := json.Unmarshal(response, &dresp)
	if err == nil {
		if dresp != nil {
			if !dresp.Ok {
				mes = fmt.Sprintf("Telegram вернул ошибку: %s", dresp.Description)
				monitor = true
			}
		}
	} else {
		mes = fmt.Sprintf("Ошибка при попытке Unmarshal: %s", err)
	}
	return mes, monitor
}

func secondMisstake(response []byte, tr *types.Telegram) string {
	var mes string
	err := json.Unmarshal(response, &tr)
	if err == nil {
		if !tr.Ok {
			if tr.Error != nil {
				mes = fmt.Sprintf("Telegram вернул ошибку: %s", tr.Error.Message)
			}
		}
	} else {
		mes = fmt.Sprintf("Ошибка при попытке Unmarshal: %s", err)
	}
	return mes
}

func handlerTelegram(response []byte, Telegram *types.Telegram) error {
	var (
		monitor bool
		descrip string
		err     error
	)
	descrip, monitor = firstMisstake(response)
	if !monitor {
		descrip = secondMisstake(response, Telegram)
	}
	if descrip != "" {
		err = fmt.Errorf(descrip)
	}
	return err
}

func (reg *RegTable) Seeker(chatID int) (index int) {
	var i int
	index = None
	if len(reg.Reg) != 0 {
		for i < len(reg.Reg) && reg.Reg[i].UserId != chatID {
			i++
		}
		if i < len(reg.Reg) {
			if reg.Reg[i].UserId == chatID {
				index = i
			}
		}
	}
	return index
}

func (reg *RegTable) NewIndex() (newIndex int) {
	newIndex = len(reg.Reg)
	reg.Reg = append(reg.Reg, Entry{})
	return newIndex
}

func handlerOffsetResponse(response []byte, offset *int) (err error) {
	var Telegram types.JustForUpdate

	err = json.Unmarshal(response, &Telegram)
	if err == nil {
		if len(Telegram.Result) > 0 {
			for _, storage := range Telegram.Result {
				*offset = storage.ID
			}
		} else {
			err = errors.UpdatesMisstakes("Telegram returned an empty data of Telegram")
		}
	}
	return err
}

func RequestOffset(TelebotToken string, offset *int) (err error) {
	var (
		response *http.Response
		body     []byte
	)

	Url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?limit=1", url.PathEscape(TelebotToken))
	response, err = http.Get(Url)
	if err == nil {
		defer response.Body.Close()
		body, err = io.ReadAll(response.Body)
	}
	if err == nil {
		err = handlerOffsetResponse(body, offset)
	}

	return err
}
