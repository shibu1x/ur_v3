package models

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`\d+`)
}

// 取得したjsonデータを解析してDBの列に登録する
func (r *Room) ParseJsonData() {
	ctx := context.Background()

	var data JsonRoom
	json.Unmarshal(r.Data.JSON, &data)

	r.Price = stringToInt(data.Rent)
	r.Fee = stringToInt(data.Commonfee)
	r.Type = null.StringFrom(data.Type)

	match := re.FindString(data.Floorspace)
	if match != "" {
		number, _ := strconv.Atoi(match)
		r.Space = number
	}
	r.Floor = stringToInt(data.Floor)
	r.LayoutURL = null.StringFrom(data.Madori)

	house, err := Houses(qm.Where("code=?", r.HouseCode)).OneG(ctx)
	if err != nil {
		return
	}

	pref, err := Prefs(qm.Where("code=?", house.PrefCode)).OneG(ctx)
	if err != nil {
		return
	}

	r.URL = null.StringFrom("https://www.ur-net.go.jp/chintai/" + pref.Region + "/" + pref.Code + "/" + r.HouseCode + "_room.html?JKSS=" + r.RoomCode)

	r.UpdateG(ctx, boil.Infer())
}

func stringToInt(str string) int {

	number, err := strconv.Atoi(strings.Join(re.FindAllString(str, -1), ""))
	if err != nil {
		// fmt.Println("Error converting to int:", err)
		return 0
	}
	return number
}
