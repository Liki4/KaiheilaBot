package ncm

import (
	"encoding/json"
	"fmt"
	"github.com/Liki4/KaiheilaBot/internal/conf"
	"github.com/lonelyevil/khl"
	"strings"
)

type SearchList struct {
	Result struct {
		Songs []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				ID        int           `json:"id"`
				Name      string        `json:"name"`
				PicURL    interface{}   `json:"picUrl"`
				Alias     []interface{} `json:"alias"`
				AlbumSize int           `json:"albumSize"`
				PicID     int           `json:"picId"`
				Img1V1URL string        `json:"img1v1Url"`
				Img1V1    int           `json:"img1v1"`
				Trans     interface{}   `json:"trans"`
			} `json:"artists"`
			Album struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Artist struct {
					ID        int           `json:"id"`
					Name      string        `json:"name"`
					PicURL    interface{}   `json:"picUrl"`
					Alias     []interface{} `json:"alias"`
					AlbumSize int           `json:"albumSize"`
					PicID     int           `json:"picId"`
					Img1V1URL string        `json:"img1v1Url"`
					Img1V1    int           `json:"img1v1"`
					Trans     interface{}   `json:"trans"`
				} `json:"artist"`
				PublishTime int64 `json:"publishTime"`
				Size        int   `json:"size"`
				CopyrightID int   `json:"copyrightId"`
				Status      int   `json:"status"`
				PicID       int64 `json:"picId"`
				Mark        int   `json:"mark"`
			} `json:"album,omitempty"`
			Duration    int           `json:"duration"`
			CopyrightID int           `json:"copyrightId"`
			Status      int           `json:"status"`
			Alias       []interface{} `json:"alias"`
			Rtype       int           `json:"rtype"`
			Ftype       int           `json:"ftype"`
			TransNames  []string      `json:"transNames,omitempty"`
			Mvid        int           `json:"mvid"`
			Fee         int           `json:"fee"`
			RURL        interface{}   `json:"rUrl"`
			Mark        int           `json:"mark"`
		} `json:"songs"`
		HasMore   bool `json:"hasMore"`
		SongCount int  `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

const SearchRespTpl string = `
%d.	Name:	%s
	Artist:	%s
	Album:	%s
---`

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func search(args string) (string, error) {
	uri := "/search?keywords=%s&type=1"
	url := conf.Get().Ncm.NcmApi + fmt.Sprintf(uri, args)
	resp, err := Client.Get(url)
	if err != nil {
		return "", err
	}
	var searchList SearchList
	if err = json.NewDecoder(resp.Body).Decode(&searchList); err != nil {
		return "", err
	}
	searchResp := "---"
	songs := searchList.Result.Songs
	for songsIndex := 0; songsIndex < min(len(songs), 3); songsIndex++ {
		song := songs[songsIndex]
		artists := ""
		for artistsIndex := 0; artistsIndex < len(song.Artists); artistsIndex++ {
			artist := song.Artists[artistsIndex]
			artists += artist.Name + "/"
		}
		artists = artists[:len(artists)-1]
		searchResp += fmt.Sprintf(SearchRespTpl, songsIndex+1, song.Name, artists, song.Album.Name)
	}
	return searchResp, nil
}

func SearchHandler(ctx *khl.TextMessageContext) {
	if ctx.Common.Type != khl.MessageTypeText || ctx.Extra.Author.Bot {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "/search") {
		resp, err := search(strings.TrimSpace(ctx.Common.Content[7:]))
		if err != nil {
			return
		}
		_, err = ctx.Session.MessageCreate(&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				Type:     khl.MessageTypeKMarkdown,
				TargetID: ctx.Common.TargetID,
				Content:  resp,
				Quote:    ctx.Common.MsgID,
			},
		})
		if err != nil {
			return
		}
	}
}
