package bilibili

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/mitchellh/mapstructure"
	clients "github.com/weridolin/simple-vedio-notifications/clients"
	"github.com/weridolin/simple-vedio-notifications/common"
	config "github.com/weridolin/simple-vedio-notifications/configs"
	"github.com/weridolin/simple-vedio-notifications/database"
	"github.com/weridolin/simple-vedio-notifications/storage"
	tools "github.com/weridolin/simple-vedio-notifications/tools"
)

var logger = config.GetLogger()
var appConfig = config.GetAppConfig()

const (
	domain          = "live.bilibili.com"
	cnName          = "哔哩哔哩"
	Host            = "api.bilibili.com"
	Origin          = "https://space.bilibili.com"
	Referer         = "https://space.bilibili.com/%d/video"
	GetVideoInfoUrl = "https://api.bilibili.com/x/space/wbi/arc/search?mid=%d&ps=30&pn=1&order=pubdate&platform=web" // ps 每页大小 pn 页码
)

//模拟浏览器请求头池
var BrowserUserAgentPool = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}

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
	emailNotifiers []*database.EmailNotifier, storage storage.StorageInterface) *BiliBiliTask {
	t := &BiliBiliTask{
		Meta: common.Meta{
			DBIndex:     dbindex,
			CallBacks:   make([]func(), 0),
			Name:        name,
			Description: description,
			Storage:     storage,
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
	t.Storage.Save(t.Result)
	logger.Println(t.DBIndex, " update result")
}

func (t *BiliBiliTask) PublicEmailNotifyMessage() {
	content, err := t.RenderEmailNotifyTemplate()
	if err != nil {
		logger.Println("render email notify template error -> ", err)
		return
	}
	logger.Println("public email notify message", content, t.EmailNotifiers, t.Error)
	if t.Error != nil {
		logger.Println("error -> ", t.Error)
		return
	} else if len(t.EmailNotifiers) > 0 {
		rabbitMq := clients.NewRabbitMQ(tools.GetUUID())
		// rabbitMq.CreateExchange(common.EmailExchangeName, "topic").
		// 	CreateQueue(common.EmailMessageQueueName, true).
		// 	ExchangeBindQueue(common.EmailMessageQueueName, "*.email.*", common.EmailExchangeName)
		for _, emailNotifier := range t.EmailNotifiers {
			// content := consumers.RenderEmailContentTemplate("bilibili",t.Ups)
			message, err := json.Marshal(
				map[string]interface{}{
					"sender":   emailNotifier.Sender,
					"pwd":      emailNotifier.PWD,
					"content":  content,
					"receiver": emailNotifier.Receiver,
				})
			logger.Println("message -> ", string(message))
			if err != nil {
				logger.Println("json marshal error", err)
				return
			}
			rabbitMq.Publish(appConfig.EmailMessageExchangeName, "bilibili.email.notify", message)
			logger.Println("send message to rabbitmq")
		}
	}
}

func (t *BiliBiliTask) Run() {
	//获取视频信息
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var result = make(map[string][]VideoInfo)

	for up_name, up_id := range t.Ups {
		logger.Println("获取 [哔哩哔哩] UP 投稿记录", "up主名字:", up_name, "up主ID:", up_id)
		var data map[string]interface{}
		url := fmt.Sprintf(GetVideoInfoUrl, int(up_id.(float64)))
		logger.Println("url = ", url)
		req, _ := http.NewRequest("GET", url, nil)

		// 伪造请求头
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
		req.Header.Add("sec-ch-ua", "'Microsoft Edge';v='113', 'Chromium';v='113', 'Not-A.Brand';v='24'")
		req.Header.Add("sec-ch-ua-mobile", "?0")
		req.Header.Add("cache-control", "max-age=0")
		req.Header.Add("sec-ch-ua-platform", "Windows")
		req.Header.Add("sec-fetch-user", "?1")
		req.Header.Add("sec-fetch-mode", "navigate")
		req.Header.Add("sec-fetch-site", "none")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("User-Agent", GetRandomUserAgent())
		// req.Header.Add("Cookie", "innersign=0; buvid3=324A040F-38FD-78B7-0BB2-687CB2E5A03D22670infoc; i-wanna-go-back=-1; b_ut=7; b_lsid=5E951F7D_188148AD1D2; bsource=search_bing; _uuid=B10627D6A-C9D6-E9106-3B6E-24108AB6C4B7522042infoc; FEED_LIVE_VERSION=V8; header_theme_version=undefined; buvid_fp=60df8955b90249b31ca1d24a40946184; home_feed_column=5; browser_resolution=1856-903; buvid4=3FDB890D-946E-B821-674C-03925B57884F24094-023051317-VheKLZviuDHxVPxcieCAUg==; b_nut=1683971823; nostalgia_conf=-1; PVID=3")
		resp, http_err := client.Do(req)
		logger.Println("rep HEADER = ", req.Header)
		if http_err != nil {
			logger.Println("http get err = ", http_err)
			t.Error = http_err
			goto Callback
			// panic(http_err)
		}
		defer resp.Body.Close()
		var body []byte
		if !resp.Uncompressed {
			reader, err := gzip.NewReader(resp.Body)
			if err != nil {
				logger.Println("gzip reader err = ", err)
				t.Error = err
				goto Callback
			}
			body, _ = ioutil.ReadAll(reader)

		} else {
			body, _ = ioutil.ReadAll(resp.Body)
		}
		// logger.Println("get result", string(body), "status code", resp.StatusCode)
		// if resp.StatusCode != 200 {
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			logger.Println("unmarshal err = ", err)
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
		// logger.Println("获取 up ->", up_name, "作品信息", videoInfoList)
		result[up_name] = videoInfoList
	}

	t.Result = result

	// 执行callback
Callback:
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

func (t *BiliBiliTask) RenderEmailNotifyTemplate() (string, error) {
	//渲染邮件模板
	logger.Println("RenderEmailNotifyTemplate")
	var content string
	buffer := new(bytes.Buffer)
	root, _ := os.Getwd()
	templatePath := path.Join(root, "static", "templates", "EmailNotifyTemplate.html")
	data, err := ioutil.ReadFile(templatePath)
	if err != nil {
		logger.Println("read template err = ", err)
		return content, nil
	}
	template, err := template.New("webpage").Parse(string(data))
	if err != nil {
		logger.Println("template parse err = ", err)
		return "", nil
	}
	templateParams := struct {
		Title   string
		Updated string
		Items   []map[string]string
	}{
		Title:   "BiliBili Up主更新提醒",
		Updated: time.Now().Format("2006-01-02 15:04:05"),
		Items:   make([]map[string]string, 0),
	}
	for upName, vedioInfoList := range t.Result.(map[string][]VideoInfo) {
		//渲染模板
		for _, videoInfo := range vedioInfoList {
			templateParams.Items = append(templateParams.Items, map[string]string{
				"Platform":   "bilibili",
				"UpName":     upName,
				"VideoTitle": videoInfo.Title,
				"VideoUrl":   `https://www.bilibili.com/video/` + videoInfo.Bvid,
				"PublicTime": time.Unix(int64(videoInfo.Created), 0).Format("2006-01-02 15:04:05"),
			})
		}
	}

	err = template.Execute(buffer, templateParams)
	return buffer.String(), nil
}

func GetRandomUserAgent() string {
	//随机获取一个列表里面的元素
	return BrowserUserAgentPool[rand.Intn(len(BrowserUserAgentPool))]
}
