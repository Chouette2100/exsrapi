// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"fmt"
	"log"
	//	"os"
	"strings"
	//	"time"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetEventidOfEventBox(
	eventid string, //	ボックスイベントの親のEvent_url_key（通常これをイベントIDと呼んでいます）
) (
	namelist []string, //	ボックスイベントの子Event_url_keyのリスト
	err error,
) {

	//	画面からのデータ取得部分は次を参考にしました。
	//		はじめてのGo言語：Golangでスクレイピングをしてみた
	//		https://qiita.com/ryo_naka/items/a08d70f003fac7fb0808

	//	_url := "https://www.showroom-live.com/event/" + EventID
	//	_url = "file:///C:/Users/kohei47/Go/src/EventRoomList03/20210128-1143.html"
	//	_url = "file:20210128-1143.html"

	var doc *goquery.Document

	namelist = make([]string, 0)

	//	URLからドキュメントを作成します
	_url := "https://www.showroom-live.com/event/" + eventid
	resp, error := http.Get(_url)
	if error != nil {
		log.Printf("GetEventInf() http.Get() err=%s\n", error.Error())
		err = fmt.Errorf("http.Get(): %w", error)
		return
	}
	defer resp.Body.Close()

	doc, error = goquery.NewDocumentFromReader(resp.Body)
	if error != nil {
		log.Printf("GetEventInf() goquery.NewDocumentFromReader() err=<%s>.\n", error.Error())
		err = fmt.Errorf("goquery.NewDocumentFromReader(): %w", error)
		return
	}

	selector := []string{
		".event-insert-section .event-insert-section a",
		".event-float-list a",
		".description2 a",
		".description-html a",
	}
	for _, sel := range selector {
		//	イベントボックス内のすべてのイベントについて繰り返す。
		doc.Find(sel).EachWithBreak(func(i int, s *goquery.Selection) bool {

			eid, exists := s.Attr("href")
			if !exists {
				return false
			}

			//	各イベントのURLの最後の要素（イベントを識別する文字列の部分）を取得しリストにします。
			eida := strings.Split(eid, "/")
			if eida[len(eida)-2] == "event" {
				namelist = append(namelist, eida[len(eida)-1])
			} else {
				log.Printf("  ignored %s\n", eid)
			}

			return true

		})

		if len(namelist) > 0 {
			break
		}
	}

	return
}
