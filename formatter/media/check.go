package media

import "log"

type checking interface {
	defaultString(string) bool
	defaultArray([]interface{}) bool
	defautBool(bool) bool
	defaultInt(int) bool
}

func (ph *Photo) defaultString(data string) bool {
	return data == ""
}

func (ph *Photo) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Photo) defautBool(data bool) bool {
	return !data
}

func (ph *Photo) defaultInt(data int) bool {
	return data == 0
}

func (ph *Video) defaultString(data string) bool {
	return data == ""
}

func (ph *Video) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Video) defautBool(data bool) bool {
	return !data
}

func (ph *Video) defaultInt(data int) bool {
	return data == 0
}

func (ph *Audio) defaultString(data string) bool {
	return data == ""
}

func (ph *Audio) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Audio) defautBool(data bool) bool {
	return !data
}

func (ph *Audio) defaultInt(data int) bool {
	return data == 0
}

func (ph *Document) defaultString(data string) bool {
	return data == ""
}

func (ph *Document) defaultArray(data []interface{}) bool {
	return len(data) == 0
}

func (ph *Document) defautBool(data bool) bool {
	return !data
}

func (ph *Document) defaultInt(data int) bool {
	return data == 0
}

func isItEmply(ch checking, work int, data interface{}) bool {
	var ok bool
	switch d := data.(type) {
	case string:
		if work == checkString {
			ok = ch.defaultString(d)
		}
	case []interface{}:
		if work == checkArray {
			ok = ch.defaultArray(d)
		}
	case bool:
		if work == checkBool {
			ok = ch.defautBool(d)
		}
	case int:
		if work == checkInt {
			ok = ch.defaultInt(d)
		}
	default:
		log.Print("SOMETHING WENT WRONG!")
	}

	return ok
}
