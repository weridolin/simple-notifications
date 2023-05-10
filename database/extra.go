package database

import (
	"database/sql/driver"
	"encoding/json"
)

type Ups map[string]interface{}

func (u Ups) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *Ups) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}

type EmailReceiver []string

func (u EmailReceiver) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *EmailReceiver) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}
