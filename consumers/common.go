package consumers

type EmailNotifyMessage struct {
	Sender   string   `json:"sender"`
	PWD      string   `json:"pwd"`
	Content  string   `json:"content"`
	Receiver []string `json:"receiver"`
}
