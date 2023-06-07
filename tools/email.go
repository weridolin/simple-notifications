package tools

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

// import (
// 	"fmt"
// 	"net/smtp"

// 	"github.com/jordan-wright/email"
// 	config "github.com/weridolin/simple-vedio-notifications/configs"
// )

// var logger = config.GetLogger()

func SendEmail(receiver []string, subject string, content string, sender, pwd string) error {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("simple-Notification <%s>", sender)
	// 设置接收方的邮箱
	e.To = receiver
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.HTML = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", sender, pwd, "smtp.qq.com"))
	if err != nil {
		logx.Info("send email failed ->", err)
	}
	return err
}
