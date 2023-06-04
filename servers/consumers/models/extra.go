package models

import (
	"database/sql/driver"
	"encoding/json"
)

type EmailReceiver []string

func (u EmailReceiver) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *EmailReceiver) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}
