package util

import (
	"encoding/json"
	"math"
	"net/mail"
)

func StructToJSON(val interface{}) interface{} {
	jsonEncoded, _ := json.Marshal(val)
	var respJSON interface{}
	json.Unmarshal([]byte(jsonEncoded), &respJSON)
	return respJSON
}

func Contains(str []string, key string) bool {
	for _, v := range str {
		if v == key {
			return true
		}
	}
	return false
}

func IsInteger(val float64) bool {
	return math.Floor(val) == math.Ceil(val)
}

func ValidEmail(email string) (string, bool) {
	mailId, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return mailId.Address, true
}
