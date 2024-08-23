package logs

import (
	"log"

	"github.com/l1qwie/Fmtogram/types"
)

var Caption string

func DebugOrInfo() {
	if types.DEBUG {
		Caption = "DEBUG"
	} else {
		Caption = "INFO"
		types.INFO = true
	}
}

func TurnOn() {
	log.Printf("[%s] The bot has been turned on", Caption)
}

func TurnOff() {
	log.Printf("[%s] The bot has been turned off", Caption)
}

func GetOffset() {
	if types.DEBUG {
		log.Panicf("[%s] Started trying to get offset from Telegram", Caption)
	}
}

func GottenOffset() {
	if types.DEBUG {
		log.Printf("[%s] The offset was given from Telegram and successfuly gotten by Fmtogram", Caption)
	}
}

func FoundSomeIformation() {
	if types.DEBUG {
		log.Printf("[%s] The bots has found some information about a user without any misstakes", Caption)
	}
}

func AlreadyInQueue() {
	if types.DEBUG {
		log.Printf("[%s] The user has already been placed in their queue", Caption)
	}
}

func FirstTimeUser() {
	if types.DEBUG {
		log.Printf("[%s] The user is a new one and has been just placed in their queue", Caption)
	}
}

func CallDeveloperFunc() {
	log.Printf("[%s] The imported function has just called by Fmtogram with some data", Caption)
}

func ReturnedIsntEmply() {
	if types.DEBUG {
		log.Printf("[%s] The returned data of the request to Telegram isn't emply and has some data", Caption)
	}
}

func ReturnedIsEmply() {
	if types.DEBUG {
		log.Printf("[%s] The returned data of the request to Telegram is emply and 'returned data' has been just genereted", Caption)
	}
}

func GottenRequest() {
	if types.DEBUG {
		log.Printf("[%s] Fmtogram has just gotten the request to send in to Telegram", Caption)
	}
}

func ErrorDuringSending() {
	log.Printf("[%s] You have an error during sending the request to Telegram", Caption)
}

func DataWrittenSuccessfully(text string) {
	if types.DEBUG {
		log.Printf("[%s] %s saved successfuly", Caption, text)
	}
}
