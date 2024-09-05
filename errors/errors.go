package errors

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func AssertTest(compared, comparedfunc, comparator, comparatorfunc string) (err error) {
	body := "[TEST ERROR] The data doesn't match. The original passed value in %s (%s) is different from the final %s (%s)"
	err = fmt.Errorf(fmt.Sprintf(body, comparedfunc, compared, comparatorfunc, comparator))
	return err
}

func JustError(message string) (err error) {
	err = fmt.Errorf("[TEST ERROR] The length of the passed arrays does not match. %s", message)
	return err
}

func UpdatesMisstakes(part string) (err error) {
	body := fmt.Sprintf("[ERROR] Couldn't get any updates: %s", part)
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
	log.Printf("[ERROR] You have an error in user's data from the response of Telegram: %s", err)
}

func ChatIDIsNil() {
	log.Print("[ERROR] You're trying to delete a message, but you haven't indicated the ChatID where the bot should delete the message")
}

func MessageIDIsNil() {
	log.Print("[ERROR] You're trying to delete a message, but you haven't indicated the Deleted Message ID which bot should delete")
}

func RequestError(err error) {
	log.Printf("[ERROR] You have an error while trying to send a request to Telegram: %s", err)
}

func CantOpenFile(err error) {
	log.Printf("[ERROR] Fmtogram couldn't open the file to send after: %s", err)
}

func CantCopyFile(err error) {
	log.Printf("[ERROR] Fmtgoram couldn't copy a part to the file: %s", err)
}

func CantCreateFormFile(err error) {
	log.Printf("[ERROR] Fmtogram couldn't create a Form File: %s", err)
}

func CantWriteField(err error) {
	log.Printf("[ERROR] Fmtogram couldn't write a multipart field: %s", err)
}

func CantMarshalJSON(err error) {
	log.Printf("[ERROR] Fmtogram couldn't marshal a json object: %s", err)
}

func ChatIDIsMissed() error {
	return fmt.Errorf("[ERROR] You must provide chatID")
}
