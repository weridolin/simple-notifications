// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Avatar       string   `json:"avatar"`
	Role         []string `json:"role"`
	IsSuperAdmin bool     `json:"is_super_admin"`
	Age          int      `json:"age"`
	Gender       int8     `json:"gender"`
}

type UserInfoWithToken struct {
	UserInfo
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Username string `form:"username"`
}

type RegisterResp struct {
	BaseResponse
	Data UserInfo `json:"data"`
}

type LoginReq struct {
	Count    string `form:"count"`
	Password string `form:"password"`
}

type LoginResp struct {
	BaseResponse
	Data UserInfoWithToken `json:"data"`
}

type LogoutResp struct {
	BaseResponse
}

type UpdateUserInfoResp struct {
	BaseResponse
	Data UserInfo `json:"data"`
}

type UpdateUserInfoRep struct {
	UserInfo
}

type RefreshTokenPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	BaseResponse
	Data RefreshTokenPayload `json:"data"`
}

type ValidateResp struct {
	UserId string `json:"user_id"`
}

type ValidateTokenReq struct {
	Authorization string `header:"authorization"`
}

type RefreshTokenReq struct {
	Authorization string `header:"authorization"`
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PaginationParams struct {
	Page int `query:"page" validate:"required,min=1"`
	Size int `query:"size" validate:"required,min=1,max=1000"`
}
