package Handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/lonelyevil/khl"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TodayEvent struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Date     string `json:"date"`
		Year     int    `json:"year"`
		Month    int    `json:"month"`
		Day      int    `json:"day"`
		Monthday string `json:"monthday"`
		Desc     string `json:"desc"`
	} `json:"data"`
	Time  int   `json:"time"`
	LogID int64 `json:"log_id"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Today() (string, error) {
	api := conf.Get().Tih.TihApi
	token := conf.Get().Tih.Token

	m := time.Now().Month()
	month := strconv.Itoa(int(m))
	day := strconv.Itoa(time.Now().Day())

	payload := strings.NewReader("token=" + token + "&monthday=" + string(month) + string(day) + "&page=1")
	//println("token=" + token + "&monthday=" + string(month) + string(day) + "&page=1")

	req, _ := http.NewRequest("POST", api, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	j_body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(j_body))

	var body TodayEvent
	err := json.Unmarshal(j_body, &body)
	if err != nil {
		return "", err
	}
	eventList := "---\n"
	for eventIndex := 0; eventIndex < min(len(body.Data), 5); eventIndex++ {
		event := body.Data[eventIndex].Desc
		eventList += fmt.Sprint(event, "\n---\n")
	}
	return eventList, nil
}

func TodayHandler(ctx *khl.TextMessageContext) {
	if ctx.Extra.ChannelName != "Test" || ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}
	list, err := Today()
	if err != nil {
		return
	}

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
