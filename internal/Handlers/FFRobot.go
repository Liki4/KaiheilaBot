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

func RobotHandler(ctx *khl.TextMessageContext) {
	if ctx.Extra.ChannelName != "Test" || ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "/ff") {
		msg := strings.TrimSpace(ctx.Common.Content[3:])
		println(msg)
		api := conf.Get().FFRobot.RobotApi
		println(api)

		//http query
		var build strings.Builder
		build.WriteString(api)
		build.WriteString(msg)
		url := fmt.Sprintf("%s", build.String())
		println(url)

		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		j_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Println(string(j_body))

		type Body struct {
			Result  int    `json:"result"`
			Content string `json:"content"`
		}
		var body Body
		err = json.Unmarshal(j_body, &body)
		body.Content = strings.ReplaceAll(body.Content, "{br}", "\n")
		fmt.Println(body.Result)
		fmt.Println(body.Content)
		if err != nil {
			return
		}

		//return to channel
		_, err = ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				Type:     khl.MessageTypeKMarkdown,
				TargetID: ctx.Common.TargetID,
				Content:  body.Content,
				Quote:    ctx.Common.MsgID,
			},
		})
		if err != nil {
			return
		}
	}
}
