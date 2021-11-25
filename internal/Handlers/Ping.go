package Handlers

import (
	"github.com/lonelyevil/khl"
	"strings"
)

func PingHandler(ctx *khl.TextMessageContext) {
	if ctx.Extra.ChannelName != "Test" || ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "/ping") {
		_, err := ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "pong",
				Quote:    ctx.Common.MsgID,
			},
		})
		if err != nil {
			return
		}
	}
}
