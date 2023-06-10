package models

import (
	"encoding/json"
)

type Ups map[string]interface{}

func (u Ups) Value() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Ups) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}

type EmailReceiver []string

func (u EmailReceiver) Value() ([]byte, error) {
	return json.Marshal(u)
}

func (u *EmailReceiver) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}
