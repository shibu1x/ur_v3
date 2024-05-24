package models

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var pattern *regexp.Regexp

func init() {
	pattern = regexp.MustCompile(`\.\./(\d+_\d+)\.html`)
}

// 都道府県にある建物の情報を取得する
func UpdateHousesAll() {
	prefs, err := Prefs(qm.Where("is_crawl=?", true)).AllG(context.Background())
	if err != nil {
		log.Println("pref no data")
		return
	}

	for _, pref := range prefs {
		log.Println(pref.Name)
		pref.updateHouses()
		pref.updateRooms()
	}
}

func (p *Pref) updateHouses() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ur-net.go.jp"),
	)

	var houses []House

	c.OnHTML("div.module_tables_apartment table tbody tr", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")

		matches := pattern.FindStringSubmatch(href)
		if len(matches) == 0 {
			fmt.Println("Did not match. " + href)
			return
		}

		houses = append(houses, House{Code: matches[1], PrefCode: p.Code, Name: e.ChildText("span.js-bukken-name")})
	})

	c.Visit("https://www.ur-net.go.jp/chintai/" + p.Region + "/" + p.Code + "/list/")

	for _, house := range houses {
		// upsert
		house.UpsertG(context.Background(), true, []string{"code"}, boil.Whitelist("pref_code", "name", "updated_at"), boil.Whitelist("code", "pref_code", "name", "created_at", "updated_at"))
	}
}

// 都道府県に紐ずく建物の部屋情報を取得する
func (p *Pref) updateRooms() {

	houses, err := Houses(qm.Where("pref_code=?", p.Code)).AllG(context.Background())
	if err != nil {
		panic(err)
	}

	for _, house := range houses {
		house.UpdateRooms()
	}
}
