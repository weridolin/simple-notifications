package consumers

type EmailNotifyMessage struct {
	Sender   string   `json:"sender"`
	PWD      string   `json:"pwd"`
	Content  string   `json:"content"`
	Receiver []string `json:"receiver"`
	Subject  string   `json:"subject"`
}

func RenderEmailContentTemplate(platform string, upName string, upID string, videoTitle string, videoURL string) string {
	return "您订阅的" + platform + "up主" + upName + "发布了新视频\n" + "视频标题：" + videoTitle + "\n" + "视频链接：" + videoURL + "\n"
}
