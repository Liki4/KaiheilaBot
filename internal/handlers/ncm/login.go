package ncm

import (
	"encoding/json"
	"fmt"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/lonelyevil/khl"
	"strings"
)

type AccountStatus struct {
	Data struct {
		Code    int         `json:"code"`
		Account interface{} `json:"account"`
		Profile interface{} `json:"profile"`
	} `json:"data"`
}

func isLoggedin() bool {
	uri := "/login/status"
	url := conf.Get().Ncm.NcmApi + uri
	resp, err := Client.Get(url)
	if err != nil {
		return false
	}
	var accountStatus AccountStatus
	if err := json.NewDecoder(resp.Body).Decode(&accountStatus); err != nil {
		return false
	}
	if accountStatus.Data.Account == nil {
		return false
	}
	return true
}

//type LoginStatus struct {
//	Code int `json:"code"`
//}

type LoginStatus struct {
	LoginType int `json:"loginType"`
	Code      int `json:"code"`
	Account   struct {
		ID                 int    `json:"id"`
		UserName           string `json:"userName"`
		Type               int    `json:"type"`
		Status             int    `json:"status"`
		WhitelistAuthority int    `json:"whitelistAuthority"`
		CreateTime         int64  `json:"createTime"`
		Salt               string `json:"salt"`
		TokenVersion       int    `json:"tokenVersion"`
		Ban                int    `json:"ban"`
		BaoyueVersion      int    `json:"baoyueVersion"`
		DonateVersion      int    `json:"donateVersion"`
		VipType            int    `json:"vipType"`
		ViptypeVersion     int64  `json:"viptypeVersion"`
		AnonimousUser      bool   `json:"anonimousUser"`
	} `json:"account"`
	Token   string `json:"token"`
	Profile struct {
		Followed                  bool        `json:"followed"`
		BackgroundURL             string      `json:"backgroundUrl"`
		DetailDescription         string      `json:"detailDescription"`
		AvatarImgIDStr            string      `json:"avatarImgIdStr"`
		BackgroundImgIDStr        string      `json:"backgroundImgIdStr"`
		UserID                    int         `json:"userId"`
		UserType                  int         `json:"userType"`
		AccountStatus             int         `json:"accountStatus"`
		VipType                   int         `json:"vipType"`
		Gender                    int         `json:"gender"`
		AvatarImgID               int64       `json:"avatarImgId"`
		Nickname                  string      `json:"nickname"`
		BackgroundImgID           int64       `json:"backgroundImgId"`
		Birthday                  int64       `json:"birthday"`
		City                      int         `json:"city"`
		AvatarURL                 string      `json:"avatarUrl"`
		DefaultAvatar             bool        `json:"defaultAvatar"`
		Province                  int         `json:"province"`
		ExpertTags                interface{} `json:"expertTags"`
		Experts                   struct{}    `json:"experts"`
		RemarkName                interface{} `json:"remarkName"`
		AuthStatus                int         `json:"authStatus"`
		Mutual                    bool        `json:"mutual"`
		DjStatus                  int         `json:"djStatus"`
		Description               string      `json:"description"`
		Signature                 string      `json:"signature"`
		Authority                 int         `json:"authority"`
		Followeds                 int         `json:"followeds"`
		Follows                   int         `json:"follows"`
		EventCount                int         `json:"eventCount"`
		AvatarDetail              interface{} `json:"avatarDetail"`
		PlaylistCount             int         `json:"playlistCount"`
		PlaylistBeSubscribedCount int         `json:"playlistBeSubscribedCount"`
	} `json:"profile"`
	Bindings []struct {
		UserID       int    `json:"userId"`
		URL          string `json:"url"`
		Expired      bool   `json:"expired"`
		BindingTime  int64  `json:"bindingTime"`
		TokenJSONStr string `json:"tokenJsonStr"`
		ExpiresIn    int64  `json:"expiresIn"`
		RefreshTime  int    `json:"refreshTime"`
		ID           int64  `json:"id"`
		Type         int    `json:"type"`
	} `json:"bindings"`
	Cookie string `json:"cookie"`
}

func login() bool {
	uri := "/login/cellphone?phone=%s&md5_password=%s"
	url := conf.Get().Ncm.NcmApi + fmt.Sprintf(uri, conf.Get().Ncm.Phone, conf.Get().Ncm.Md5Pass)
	resp, err := Client.Get(url)
	if err != nil {
		return false
	}
	var loginStatus LoginStatus
	if err := json.NewDecoder(resp.Body).Decode(&loginStatus); err != nil {
		return false
	}
	return loginStatus.Code == 200
}

func LoginHandler(ctx *khl.TextMessageContext) {
	if ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "/login") {
		if isLoggedin() {
			_, err := ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					TargetID: ctx.Common.TargetID,
					Content:  "logged in.",
					Quote:    ctx.Common.MsgID,
				},
			})
			if err != nil {
				return
			}
		} else if login() {
			_, err := ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					TargetID: ctx.Common.TargetID,
					Content:  "login successful.",
					Quote:    ctx.Common.MsgID,
				},
			})
			if err != nil {
				return
			}
		} else {
			_, err := ctx.Session.MessageCreate(&khl.MessageCreate{
				MessageCreateBase: khl.MessageCreateBase{
					TargetID: ctx.Common.TargetID,
					Content:  "login failed.",
					Quote:    ctx.Common.MsgID,
				},
			})
			if err != nil {
				return
			}
		}
	}
}
