package Handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/lonelyevil/khl"
	"io/ioutil"
	"net/http"
	"strings"
)

type TodayTopList struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name       string `json:"name"`
		LastUpdate string `json:"last_update"`
		List       []struct {
			Title string `json:"title"`
			Link  string `json:"link"`
			Other string `json:"other"`
		} `json:"list"`
	} `json:"data"`
	Time  int   `json:"time"`
	LogID int64 `json:"log_id"`
}

var id_map = map[string]string{"知乎": "mproPpoq6O", "微博": "KqndgxeLl9",
	"B站日榜": "74KvxwokxM", "B站综合": "VaobLKGdAj", "百度": "Jb0vmloB1G",
	"豆瓣新片": "mDOvnyBoEB", "豆瓣新剧": "nBe0JLBv37", "小黑盒": "47o8QYwvMm"}

func TodayTop(msg string) (string, error) {
	var index int
	var key string
	for key = range id_map {
		index = strings.Index(msg, key)
		if index != -1 {
			break
		}
	}
	if index == -1 {
		eventList := "小伙伴你好，羊驼目前支持的热搜查询包括：知乎热搜，微博热搜，百度热搜，" +
			"B站日榜、B站综合榜，豆瓣新片、豆瓣新剧，小黑盒新闻。查询时请加上‘热搜’的tag哦～"
		//println(eventList)
		return eventList, nil
	}

	index = index / 3
	keyRune := []rune(key)
	msgRune := []rune(msg)
	idString := string(msgRune[index : index+len(keyRune)])
	id, _ := id_map[idString]

	api := conf.Get().Alapi.Tpapi
	token := conf.Get().Alapi.Token
	payload := strings.NewReader("token=" + token + "&id=" + id)

	req, _ := http.NewRequest("POST", api, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	jBody, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(jBody))

	var body TodayTopList
	err := json.Unmarshal(jBody, &body)
	if err != nil {
		return "", err
	}
	eventList := body.Data.Name + "\n" + body.Data.LastUpdate + "\n"
	for eventIndex := 0; eventIndex < min(len(body.Data.List), 10); eventIndex++ {
		eventTitle := "标题：[" + body.Data.List[eventIndex].Title + "]"
		eventLink := "(" + body.Data.List[eventIndex].Link + ")"
		eventHeat := "热度：" + body.Data.List[eventIndex].Other
		eventList += fmt.Sprint(eventTitle, eventLink, "\n", eventHeat, "\n---\n")
	}
	return eventList, nil
}

func TodayTopHandler(ctx *khl.TextMessageContext, msg string) {
	if ctx.Extra.ChannelName != "Test" || ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}
	list, err := TodayTop(msg)
	if err != nil {
		return
	}
	println(list)

	//return to channel
	_, err = ctx.Session.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeKMarkdown,
			TargetID: ctx.Common.TargetID,
			Content:  list,
			Quote:    ctx.Common.MsgID,
		},
	})
	if err != nil {
		return
	}
}
