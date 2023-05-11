package bilibili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
	clients "github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/common"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/database"
	tools "github.com/weridolin/simple-vedio-notifications/tools"
)

var logger = config.GetLogger()

const (
	domain          = "live.bilibili.com"
	cnName          = "哔哩哔哩"
	Host            = "api.bilibili.com"
	Origin          = "https://space.bilibili.com"
	Referer         = "https://space.bilibili.com/%d/video"
	GetVideoInfoUrl = "https://api.bilibili.com/x/space/arc/search?mid=%d&ps=30&tid=0&pn=1&keyword=&order=pubdate&platform=web"
)

//模拟浏览器请求头池
var BrowserUserAgentPool = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0}"}

type VideoInfoStatus struct {
	Season_id int
	View      int
	Danmaku   int
	Reply     int
	Favorite  int
	Coin      int
	Share     int
	Like      int
	Mtime     int
	Vt        int
	Vv        int
}

type VideoInfoMeta struct {
	Id         int
	Title      string
	Cover      string
	Mid        int
	Intro      string
	Sign_state int
	Attribute  int
	Stat       VideoInfoStatus
}

type VideoInfo struct {
	Comment          int
	Typeid           int
	Play             int
	Pic              string
	Subtitle         string
	Description      string
	Copyright        string
	Title            string
	Review           int
	Author           string
	Mid              int
	Created          int
	Length           string
	Video_review     int
	Aid              int
	Bvid             string
	Hide_click       bool
	Is_pay           int
	Is_union_video   int
	Is_steins_gate   int
	Is_live_playback int
	Meta             VideoInfoMeta
	Is_avoided       int
	Attribute        int
}

type BiliBiliTask struct {
	common.Meta
	Ups            database.Ups
	Period         tools.Period
	Error          error
	Result         interface{}
	EmailNotifiers []*database.EmailNotifier
}

func NewBiliBiliTask(period tools.Period, ups database.Ups, dbindex uint, name, description string,
	emailNotifiers []*database.EmailNotifier) *BiliBiliTask {
	t := &BiliBiliTask{
		Meta: common.Meta{
			DBIndex:     dbindex,
			CallBacks:   make([]func(), 0),
			Name:        name,
			Description: description,
		},
		EmailNotifiers: emailNotifiers,
		Period:         period,
		Ups:            ups,
	}
	t.CallBacks = append(t.CallBacks, t.UpdateResult, t.PublicEmailNotifyMessage)
	return t
}

func (t *BiliBiliTask) UpdateResult() {
	if t.Error != nil {
		logger.Println(t.DBIndex, " run error -> ", t.Error)

	}
	logger.Println(t.DBIndex, " update result")
}

func (t *BiliBiliTask) PublicEmailNotifyMessage() {
	if t.Error != nil && len(t.EmailNotifiers) > 0 {
		rabbitMq := clients.NewRabbitMQTopic("emailNotify", "bilibili.email.notify")
		for _, emailNotifier := range t.EmailNotifiers {
			message, err := json.Marshal(
				map[string]interface{}{
					"sender":   emailNotifier.Sender,
					"pwd":      emailNotifier.PWD,
					"content":  "todo",
					"receiver": emailNotifier.Receiver,
				})
			if err != nil {
				logger.Println("json marshal error", err)
			}
			rabbitMq.PublishTopic(message)
			logger.Panicln("send message to rabbitmq")
		}
	}
}

func (t *BiliBiliTask) Run() {
	//获取视频信息
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	for up_name, up_id := range t.Ups {
		logger.Println("get up public records", "up_name = ", up_name, "up_id = ", up_id)
		var data map[string]interface{}
		url := fmt.Sprintf(GetVideoInfoUrl, int(up_id.(float64)))
		logger.Println("url = ", url)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Host", Host)
		req.Header.Add("Origin", Origin)
		req.Header.Add("Referer", fmt.Sprintf(Referer, up_id))
		req.Header.Add("User-Agent", GetRandomUserAgent())
		resp, http_err := client.Do(req)
		if http_err != nil {
			logger.Panicln("http get err = ", http_err)
			// panic(http_err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		logger.Println("get result", string(body), "status code", resp.StatusCode)
		// if resp.StatusCode != 200 {
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			logger.Println("unmarshal err = ", err)
			// fmt.Println(resp.StatusCode, string(body))
			t.Error = err
			break
		}

		// 将返回的结果转换为结构体
		VideoListResponse := data["data"].(map[string]interface{})["list"].(map[string]interface{})["vlist"].([]interface{})
		var videoInfoList []VideoInfo
		if err := mapstructure.Decode(VideoListResponse, &videoInfoList); err != nil {
			t.Error = err
			logger.Println("mapstructure decode err = ", err)
		}
		logger.Println("获取 up ->", up_name, "作品信息", videoInfoList)
	}

	// 执行callback
	for _, callback := range t.CallBacks {
		callback()
	}
	// return videoInfoList
}

func (t *BiliBiliTask) GetUpInfo() {
	//获取up主信息
	fmt.Println("GetUpInfo")
}

func (t *BiliBiliTask) Stop() {
	fmt.Println("stop bilibili task")
}

func GetRandomUserAgent() string {
	//随机获取一个列表里面的元素
	return BrowserUserAgentPool[rand.Intn(len(BrowserUserAgentPool))]
}

func UpdateResultToRedis() {
	//更新结果到redis
	logger.Println("UpdateResultToRedis")
}
