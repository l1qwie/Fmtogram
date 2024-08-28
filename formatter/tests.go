package formatter

// import (
// 	"fmt"

// 	"github.com/l1qwie/Fmtogram/errors"
// )

// // You can compare data of a photo you want to send and data of a photo you're expecting to be sent
// // As the first agrument you should give data about the expected photo. It could be path in your storage, url or telegramID.
// // Whatever you hand over in AddPhotoFrom[...](). As the second argument you can hand over true and then the function will
// // panic as soon as it will discover an error. Or false and the function will return you an error and you can handle it as you want.
// func (tfm *Formatter) AssertPhoto(path string, doPanic bool) error {
// 	var (
// 		function string
// 		err      error
// 	)
// 	if len(tfm.Message.Photo) > 0 {
// 		if tfm.Message.Photo != path {
// 			if tfm.kindofmedia[0] == fromStorage {
// 				function = "AddPhotoFromStorage"

// 			} else if tfm.kindofmedia[0] == fromInternet {
// 				function = "AddPhotoFromInternet"

// 			} else if tfm.kindofmedia[0] == fromTelegram {
// 				function = "AddPhotoFromTG"

// 			}
// 			err = errors.AssertTest(tfm.Message.Photo, function, path, "AssertPhoto")
// 		}
// 	} else {
// 		err = errors.AssertTest(fmt.Sprint(tfm.Message.Photo), function, path, "AssertPhoto")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // You can compare data of a video you want to send and data of a video you're expecting to be sent
// // As the first agrument you should give data about the expected video. It could be path in your storage, url or telegramID.
// // Whatever you hand over in AddVideoFrom[...](). As the second argument you can hand over true and then the function will
// // panic as soon as it will discover an error. Or false and the function will return you an error and you can handle it as you want.
// func (tfm *Formatter) AssertVideo(path string, doPanic bool) (err error) {
// 	var function string
// 	if len(tfm.Message.Video) > 0 {
// 		if tfm.Message.Video != path {
// 			if tfm.kindofmedia[0] == fromStorage {
// 				function = "AddVideoFromStorage"
// 			} else if tfm.kindofmedia[0] == fromInternet {
// 				function = "AddVideoFromInternet"
// 			} else if tfm.kindofmedia[0] == fromTelegram {
// 				function = "AddVideoFromTG"
// 			}
// 			err = errors.AssertTest(tfm.Message.Video, function, path, "AssertVideo")
// 		}
// 	} else {
// 		err = errors.AssertTest(fmt.Sprint(tfm.Message.Video), function, path, "AssertPhoto")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// func (tfm *Formatter) checkLengthOfKeyboardButtons(testdim []int, doPanic bool) error {
// 	var (
// 		dim []int
// 		err error
// 	)

// 	for i := 0; i < len(tfm.Keyboard.Keyboard); i++ {
// 		dim = append(dim, len(tfm.Keyboard.Keyboard[i]))
// 	}
// 	if len(testdim) == len(dim) {
// 		for i := 0; i < len(dim); i++ {
// 			if testdim[i] != dim[i] {
// 				err = errors.AssertTest(fmt.Sprint(dim), "SetIkbdDim", fmt.Sprint(testdim), "AssertInlineKeyboard")
// 				if doPanic {
// 					panic(err)
// 				}
// 			}
// 		}
// 	} else {
// 		err = errors.AssertTest(fmt.Sprint("length of slice is ", len(dim)), "SetIkbdDim", fmt.Sprint("length of slice is ", len(testdim)), "AssertInlineKeyboard")
// 		if doPanic {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// func (tfm *Formatter) checkDataOfButtos(testdim []int, kbNames, kbDatas, typeofbuttons []string, doPanic bool) error {
// 	var (
// 		counter int
// 		err     error
// 		message string
// 	)

// 	if len(kbNames) == len(kbDatas) && len(kbNames) == len(typeofbuttons) {
// 		for i := 0; i < len(testdim); i++ {
// 			for j := 0; j < testdim[i]; j++ {
// 				if tfm.Keyboard.Keyboard[i][j].Label != kbNames[counter] {
// 					err = errors.AssertTest(fmt.Sprint("name of buttons is ", tfm.Keyboard.Keyboard[i][j].Label), "WriteInlineButtonUrl/WriteInlineButtonCmd", fmt.Sprint("name of buttons is ", kbNames[counter]), "AssertInlineKeyboard")
// 					if doPanic {
// 						panic(err)
// 					}
// 				} else if typeofbuttons[i] == "url" && tfm.Keyboard.Keyboard[i][j].Url != kbDatas[counter] {
// 					err = errors.AssertTest(fmt.Sprint("url of button is ", tfm.Keyboard.Keyboard[i][j].Url), "WriteInlineButtonUrl", fmt.Sprint("url of button is ", kbDatas[counter]), "AssertInlineKeyboard")
// 					if doPanic {
// 						panic(err)
// 					}
// 				} else if typeofbuttons[i] == "cmd" && tfm.Keyboard.Keyboard[i][j].Cmd != kbDatas[counter] {
// 					err = errors.AssertTest(fmt.Sprint("cmd of button is ", tfm.Keyboard.Keyboard[i][j].Cmd), "WriteInlineButtonCmd", fmt.Sprint("cmd of button is ", kbDatas[counter]), "AssertInlineKeyboard")
// 					if doPanic {
// 						panic(err)
// 					}
// 				}
// 				counter++
// 			}
// 		}
// 	} else {

// 		if len(kbNames) == len(kbDatas) {
// 			message = "Length of kbNames = %d, but length of kbDatas = %d"
// 		}
// 		if len(kbNames) == len(typeofbuttons) {
// 			if message != "" {
// 				message += ". "
// 			}
// 			message += "Length of kbNames = %d, but length of typeofbuttons = %d"
// 		}
// 		err = errors.JustError(message)
// 		if doPanic {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // If you want to compare the inline-keyboard you're gonna send and your expectations use this function.
// // The first argument should be an array of int where the whole array will be just a string of buttons. So you should call this function
// // for every string of buttons you're expecting. So, an example of the first argument is [1, 1, 1, 1] (it means you have 4 buttons in one inline-keyboard string)
// // The second argument should be the names of the buttons. The third should be the data of the buttons. The fourth argument it is the types of the buttons (cmd or url).
// // And the last one should be true only if you want to do panic every time the function discovers an error
// func (tfm *Formatter) AssertInlineKeyboard(testdim []int, kbNames, kbDatas, typeofbuttons []string, doPanic bool) error {
// 	err := tfm.checkLengthOfKeyboardButtons(testdim, doPanic)
// 	if err == nil {
// 		err = tfm.checkDataOfButtos(testdim, kbNames, kbDatas, typeofbuttons, doPanic)
// 	}
// 	return err
// }

// // To check the text you're sending to some chat use this function.
// // It compares your data that you handed over in the past (hopefully in your bussines logic) and the data you handed over just right now
// // If the last argument is true the function will panic as soon as it discovers an error
// func (tfm *Formatter) AssertString(lineoftext string, doPanic bool) error {
// 	var err error
// 	if tfm.Message.Text != lineoftext {
// 		err = errors.AssertTest(fmt.Sprint(tfm.Message.Text), "WriteString", fmt.Sprint(lineoftext), "AssertString")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // Check the ChatId you'd handed over before and handed just right now
// // doPanic should be true if you want function panic as fast as it discovers en error
// func (tfm *Formatter) AssertChatId(chatID int, doPanic bool) error {
// 	var err error
// 	if chatid, ok := tfm.Message.ChatID.(int); ok {
// 		if chatid != chatID {
// 			err = errors.AssertTest(fmt.Sprint(tfm.Message.ChatID), "WriteChatId", fmt.Sprint(chatID), "AssertChatId")
// 		}
// 	} else {
// 		err = fmt.Errorf("the chatId isn't int as was expected")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // Check the Chat Name you'd handed over before and handed just right now
// // doPanic should be true if you want function panic as fast as it discovers en error
// func (tfm *Formatter) AssertChatName(chatName string, doPanic bool) error {
// 	var err error
// 	if chatid, ok := tfm.Message.ChatID.(string); ok {
// 		if chatid != fmt.Sprint("@", chatName) {
// 			err = errors.AssertTest(fmt.Sprint(tfm.Message.ChatID), "WriteChatName", fmt.Sprint(chatName), "AssertChatName")
// 		}
// 	} else {
// 		err = fmt.Errorf("the chatId isn't string as was expected")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // Check the Parse Mode you'd handed over before and handed just right now
// // doPanic should be true if you want function panic as fast as it discovers en error
// func (tfm *Formatter) AssertParseMode(parseMode string, doPanic bool) error {
// 	var err error
// 	if tfm.Message.ParseMode != parseMode {
// 		err = errors.AssertTest(fmt.Sprintf(tfm.Message.ParseMode), "WriteParseMode", fmt.Sprint(parseMode), "AssertParseMode")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // Check the Edited Message ID you'd handed over before and handed just right now
// // doPanic should be true if you want function panic as fast as it discovers en error
// func (tfm *Formatter) AssertEditMessageId(messageId int, doPanic bool) error {
// 	var err error
// 	if tfm.Message.MessageId != messageId {
// 		err = errors.AssertTest(fmt.Sprint(tfm.Message.MessageId), "WriteEditMesId", fmt.Sprint(messageId), "AssertEditMessageId")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }

// // Check the Deleted Message ID you'd handed over before and handed just right now
// // doPanic should be true if you want function panic as fast as it discovers en error
// func (tfm *Formatter) AssertDeleteMessageId(messageId int, doPanic bool) error {
// 	var err error
// 	if tfm.DeletedMessage.MessageId != messageId {
// 		err = errors.AssertTest(fmt.Sprint(tfm.DeletedMessage.MessageId), "WriteDeleteMesId", fmt.Sprint(messageId), "AssertDeleteMessageId")
// 	}
// 	if doPanic {
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	return err
// }
