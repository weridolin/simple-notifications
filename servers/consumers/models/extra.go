package models

import "encoding/json"

type EmailReceiver []string

func (u EmailReceiver) Value() (any, error) {
	return json.Marshal(u)
}

func (u *EmailReceiver) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}
