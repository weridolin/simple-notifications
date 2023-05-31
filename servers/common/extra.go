package common

import "encoding/json"

type Role []string

func (u Role) Value() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Role) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}
