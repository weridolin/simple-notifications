package tools

import "github.com/gofrs/uuid"

func GetUUID() string {
	//获取UUID4
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
