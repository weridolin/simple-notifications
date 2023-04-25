package users

// type UserValidator struct {
// 	Username string `json:"username" binding:"exists,alphanum,min=4,max=255" form:"username"`
// 	Password string `json:"password" binding:"required,min=8,max=255" form:"password"`
// 	Email    string `json:"email" binding:"exists,email" form:"email"`
// 	Phone    string `json:"phone" form:"phone"`
// 	Avatar   string `json:"avatar" form:"avatar"`
// 	Age      int    `json:"age"  form:"age"`
// 	Sex      int8   `json:"sex" binding:"exists,oneof=0,1" form:"sex"`
// }

// func NewUserValidator() UserValidator {
// 	return UserValidator{}
// }

// func (u *UserValidator) Bind(c *gin.Context) error {
// 	err := tools.Bind(c, u)
// 	if err != nil {
// 		return err
// 	}
// 	// 处理下密码
// 	u.Password = tools.GetMD5Hash(u.Password)
// 	return nil
// }
