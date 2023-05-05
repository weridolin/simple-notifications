package common

type PaginationValidator struct {
	Page int `json:"page" form:"page" binding:"required,min=1" query:"page,default=1" `
	Size int `json:"size" form:"size" binding:"required,min=1,max=100" query:"size,default=100" `
}
