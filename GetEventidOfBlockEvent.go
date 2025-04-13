// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"fmt"
	"log"
	"strings"

	"encoding/json"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Block struct {
	Label   string
	Block_id int
}

type BlockInf struct {
	Show_rank_label string
	Block_list     []Block
}

type BlockInfList struct {
	Blockinf []BlockInf
}


//	ブロックイベントの子のイベントのeventidを取得する。
func GetEventidOfBlockEvent(
	eventid string,		//	ブロックイベントの親イベントのeventid
) (
	blockinflist     BlockInfList ,	//	このブロックイベントのラベルとブロック番号のペア

	err error,
) {

	//	画面からのデータ取得部分は次を参考にしました。
	//		はじめてのGo言語：Golangでスクレイピングをしてみた
	//		https://qiita.com/ryo_naka/items/a08d70f003fac7fb0808

	var doc *goquery.Document

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

	//	ブロック情報がJSONとして得られる
	//	tjson, bl := doc.Find(".js-event-lower-cate-section div div event-block").Attr("data-list")
	tjson, bl := doc.Find("#js-event-block > event-block").Attr("data-list")
	if !bl {
		err = fmt.Errorf("doc.Find().Attr(): %t", bl)
		return
	}
	//	tjson = tjson[ 1: len(tjson)-1]

	err = json.NewDecoder(strings.NewReader(tjson)).Decode(&blockinflist.Blockinf)
	if err != nil {
		err = fmt.Errorf("json.NewDecoder().Decode(): %w", err)
		return
	}

	return
}
