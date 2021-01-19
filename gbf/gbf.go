package gbf

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"github.com/zhshch2002/goreq"
)

type Paginate struct {
	Count   string `json:"count"`
	First   int    `json:"first"`
	Last    int    `json:"last"`
	Page    string `json:"page"`
	Current string `json:"current"`
	Next    int    `json:"next"`
	Prev    int    `json:"prev"`
	List    List   `json:"list"`
}

type List []Item

type Item struct {
	Rank        string `json:"rank"`
	UserID      string `json:"user_id"`
	Level       string `json:"level"`
	Name        string `json:"name"`
	Point       string `json:"point"`
}

func GetPage(page int) List {
	url := fmt.Sprintf("http://game.granbluefantasy.jp/teamraid055/rest_ranking_user/detail/%d/0", page)
	res := goreq.Get(url).AddHeaders(map[string]string{
		"Cookie":     "midship=S%3AolrGoZzfcYb3vHLWCOKvHtaxnvtHvvYE-y0wupvuFjvLDO9_gFHTDLyg4ChMuFH0rr-lztq7I-_wzbnOBevUpelVSs3TLkqGSEsN5zePwef7252V45pnTHqKAWY13j6olptCrZcLDCV8ntLGLhAtsZ1yIC6mRVmxRUCaFC_CcVwtRoFIeXpl8NOaXTBZnrNnGprmI-qXm7rf7Y_fuayzOL0PsPtyxJ96eI2EO81uoU66Cw%3D%3D",
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36"})
	resp := res.Do()
	json, _ := resp.JSON()
	if json.Type != gjson.JSON {
		fmt.Println(resp.Txt())
		panic("登陆失效")
	}
	str := []byte(json.Raw)
	paginate := &Paginate{}
	_ = jsoniter.Unmarshal(str, &paginate)
	return paginate.List
}
