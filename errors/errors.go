package errors

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/l1qwie/Fmtogram/logs"
)

var head string = "\n\nAnswer from byogram: "

func AssertTest(compared, comparedfunc, comparator, comparatorfunc string) (err error) {
	body := "[TEST LOG] Error: The data doesn't match. The original passed value in %s (%s) is different from the final %s (%s)"
	err = fmt.Errorf(fmt.Sprintf(body, comparedfunc, compared, comparatorfunc, comparator))
	return err
}

func JustError(message string) (err error) {
	err = fmt.Errorf("[TEST LOG] Error: the length of the passed arrays does not match. %s", message)
	return err
}

func UpdatesMisstakes(part string) (err error) {
	body := fmt.Sprintf(head, part)
	err = fmt.Errorf(body)

	return err
}

func MadeMisstake(err error) {
	log.Println("You have a misstake:", err)
	debug.PrintStack()
}

func MakeIntestines() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from log.Fatal:", r)
		debug.PrintStack()
		os.Exit(1)
	}
}

func ErrorInUserData(err error) {
	log.Printf("[%s] You have an error in user's data from the response of Telegram: %s", logs.Caption, err)
}

func ChatIDIsNil() {
	log.Print("[WARNING] You're trying to delete a message, but you haven't indicated the ChatID where the bot should delete the message")
}

func MessageIDIsNil() {
	log.Print("[WARNING] You're trying to delete a message, but you haven't indicated the Deleted Message ID which bot should delete")
}

func RequestError(err error) {
	log.Printf("[WARNING] You have an error while trying to send a request to Telegram: %s", err)
}
