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
