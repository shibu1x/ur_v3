package models

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UR APIが返す部屋のJSON構造
type JsonRoom struct {
	PageIndex        string `json:"pageIndex"`
	RowMax           string `json:"rowMax"`
	RowMaxSp         string `json:"rowMaxSp"`
	RowMaxNext       string `json:"rowMaxNext"`
	PageMax          string `json:"pageMax"`
	AllCount         string `json:"allCount"`
	Block            string `json:"block"`
	Tdfk             string `json:"tdfk"`
	Shisya           string `json:"shisya"`
	Danchi           string `json:"danchi"`
	Shikibetu        string `json:"shikibetu"`
	FloorAll         string `json:"floorAll"`
	RoomDetailLink   string `json:"roomDetailLink"`
	RoomDetailLinkSp string `json:"roomDetailLinkSp"`
	System           []struct {
		IMG           string `json:"制度_IMG"`
		NAMING_FAILED string `json:"制度名"`
		HTML          string `json:"制度HTML"`
	} `json:"system"`
	Parking       any    `json:"parking"`
	Design        []any  `json:"design"`
	FeatureParam  []any  `json:"featureParam"`
	Traffic       any    `json:"traffic"`
	Place         any    `json:"place"`
	Kanris        any    `json:"kanris"`
	Kouzou        any    `json:"kouzou"`
	Soukosu       any    `json:"soukosu"`
	ID            string `json:"id"`
	Year          any    `json:"year"`
	Name          string `json:"name"`
	Shikikin      string `json:"shikikin"`
	Requirement   string `json:"requirement"`
	Madori        string `json:"madori"`
	Rent          string `json:"rent"`
	RentNormal    string `json:"rent_normal"`
	RentNormalCSS string `json:"rent_normal_css"`
	Commonfee     string `json:"commonfee"`
	CommonfeeSp   any    `json:"commonfee_sp"`
	Status        any    `json:"status"`
	Type          string `json:"type"`
	Floorspace    string `json:"floorspace"`
	Floor         string `json:"floor"`
	URLDetail     any    `json:"urlDetail"`
	URLDetailSp   any    `json:"urlDetail_sp"`
	Feature       any    `json:"feature"`
}

// 建物にある部屋の情報を取得する
func (h *House) UpdateRooms() {

	log.Println(h.Code)

	if time.Since(h.RoomsGotAt).Hours() < 12 {
		// 前回取得から一定時間経過していない
		return
	}

	ctx := context.Background()

	now := time.Now()
	for i := 0; i < 100; i++ {
		data := h.getRooms(i)
		if len(data) == 0 {
			break
		}

		var rooms []Room
		for _, v := range data {
			data, err := json.Marshal(v)
			if err != nil {
				continue
			}

			rooms = append(rooms, Room{
				HouseCode: h.Code,
				RoomCode:  v.ID,
				Status:    "ready",
				Data:      null.JSONFrom(data),
				GotAt:     now,
			})
		}

		for _, room := range rooms {
			log.Println(room.RoomCode)
			// upsert
			room.UpsertG(ctx, true, []string{"house_code", "room_code"}, boil.Whitelist("status", "got_at", "data", "updated_at"), boil.Whitelist("house_code", "room_code", "status", "got_at", "data", "created_at", "updated_at"))
		}

		pageMax, _ := strconv.Atoi(data[0].PageMax)
		if i >= pageMax-1 {
			break
		}
	}

	h.RoomsGotAt = now
	h.UpdateG(ctx, boil.Infer())

	h.updateStatusClosed()

	rooms, _ := Rooms(qm.Where("house_code=?", h.Code), qm.Where("status=?", "ready")).AllG(ctx)
	for _, room := range rooms {
		room.ParseJsonData()
	}
}

// 取得できなかったデータは受付終了とする
func (h *House) updateStatusClosed() {

	ctx := context.Background()

	rooms, err := Rooms(qm.Where("house_code=?", h.Code), qm.Where("status=?", "ready"), qm.Where("got_at<?", time.Now().Format("2006-01-02"))).AllG(ctx)
	if err != nil {
		return
	}

	rooms.UpdateAllG(ctx, M{"status": "closed"})
}

// UR API から建物の部屋を取得する
func (h *House) getRooms(index int) []JsonRoom {
	params := h.getSearchParams()
	for k, v := range map[string][]string{
		"orderByField": {"0"},
		"orderBySort":  {"0"},
		"pageIndex":    {strconv.Itoa(index)},
	} {
		params[k] = v
	}

	resp, err := http.Post(
		"https://chintai.r6.ur-net.go.jp/chintai/api/bukken/detail/detail_bukken_room/",
		"application/x-www-form-urlencoded",
		strings.NewReader(url.Values(params).Encode()),
	)

	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data []JsonRoom
	json.Unmarshal(body, &data)

	return data
}

// UR API に必要なパラメーターを取得する
func (h *House) getSearchParams() map[string][]string {
	return map[string][]string{
		"shisya":    {h.Code[:2]},
		"danchi":    {h.Code[3:6]},
		"shikibetu": {h.Code[len(h.Code)-1:]},
	}
}
