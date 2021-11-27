package Handlers

import (
	"encoding/json"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/lonelyevil/khl"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RobotContent struct {
	Result  int    `json:"result"`
	Content string `json:"content"`
}

func RobotCommunicate(msg string) (string, error) {
	api := conf.Get().FFRobot.RobotApi

	queryUrl := api + url.QueryEscape(url.QueryEscape(msg))
	//println(queryUrl)

	resp, err := http.Get(queryUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	jBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(j_body))

	var body RobotContent
	err = json.Unmarshal(jBody, &body)
	if err != nil {
		return "", err
	}

	body.Content = strings.ReplaceAll(body.Content, "{br}", "\n")
	//fmt.Println(body.Result)
	//fmt.Println(body.Content	)
	return body.Content, nil
}

func RobotHandler(ctx *khl.TextMessageContext) {
	if ctx.Extra.ChannelName != "Test" || ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "羊驼") {
		msg := strings.TrimSpace(ctx.Common.Content[6:])
		//println(msg)
		switch true {
		case strings.Index(msg, "历史上的今天") != -1:
			TodayInHistoryHandler(ctx)
		case strings.Index(msg, "热搜") != -1:
			TodayTopHandler(ctx, msg)
		default:
			body, err := RobotCommunicate(msg)
			if err != nil {
				return
			}

			//return to channel
			_, err = ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					Type:     khl.MessageTypeKMarkdown,
					TargetID: ctx.Common.TargetID,
					Content:  body,
					Quote:    ctx.Common.MsgID,
				},
			})
			if err != nil {
				return
			}
		}
	}
}
