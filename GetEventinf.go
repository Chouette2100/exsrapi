/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	//	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Event_Inf struct {
	Event_ID    string
	I_Event_ID  int
	Event_name  string
	Period      string
	Dperiod     float64
	Start_time  time.Time
	Sstart_time string
	Start_date  float64
	End_time    time.Time
	Send_time   string
	NoEntry     int
	NoRoom      int //      ルーム数
	Intervalmin int
	Modmin      int
	Modsec      int
	Fromorder   int
	Toorder     int
	Resethh     int
	Resetmm     int
	Nobasis     int
	Maxdsp      int
	Cmap        int
	Target      int
	Rstatus     string
	Maxpoint    int //      グラフのy軸のスケールを固定する
	MaxPoint    int //      DBには該当するものはない（データ中最大のポイント値を意味し、内部的に使用する）
	Gscale      int //      DBのMaxpoint = 構造体の Maxpoint + Gscale
	Achk        int //      1: ブロック、2:ボックス、子ルーム未処理のあいだはそれぞれ +4
	Aclr        int //未使用

	//      Event_no    int
	EventStatus string //   "Over", "BeingHeld", "NotHeldYet"
	Pntbasis    int
	Ordbasis    int
	League_ids  string
	//      Status          string          //      "Confirmed":    イベント終了日翌日に確定した獲得ポイントが反映されている。
}

func GetEventinf(
	eventid string,
	eventinfo *Event_Inf,
) (
	err error,
) {

	//      画面からのデータ取得部分は次を参考にしました。
	//              はじめてのGo言語：Golangでスクレイピングをしてみた
	//              https://qiita.com/ryo_naka/items/a08d70f003fac7fb0808

	//      _url := "https://www.showroom-live.com/event/" + EventID
	//      _url = "file:///C:/Users/kohei47/Go/src/EventRoomList03/20210128-1143.html"
	//      _url = "file:20210128-1143.html"

	var doc *goquery.Document

	inputmode := "url"
	eventidorfilename := eventid

	/*
		_, _, status := SelectEventNoAndName(eventidorfilename)
		log.Printf(" status=%d\n", status)
		if status != 0 {
				return
		}
		(*eventinfo).Event_no = eventno
	*/

	if inputmode == "file" {

		//      ファイルからドキュメントを作成します
		f, e := os.Open(eventidorfilename)
		if e != nil {
			//      log.Fatal(e)
			log.Printf("err=[%s]\n", e.Error())

			return
		}
		defer f.Close()
		doc, err = goquery.NewDocumentFromReader(f)
		if err != nil {
			err = fmt.Errorf("goquery.NewDocumentFromReader(): %w", err)
			return
		}

		content, _ := doc.Find("head > meta:nth-child(6)").Attr("content")
		content_div := strings.Split(content, "/")
		eventinfo.Event_ID = content_div[len(content_div)-1]

	} else {
		//      URLからドキュメントを作成します
		_url := "https://www.showroom-live.com/event/" + eventidorfilename
		/*
		   doc, err = goquery.NewDocument(_url)
		*/
		resp, error := http.Get(_url)
		if error != nil {
			log.Printf("GetEventInfAndRoomList() http.Get() err=%s\n", error.Error())
			err = fmt.Errorf("http.Get(): %w", error)
			return
		}
		defer resp.Body.Close()

		doc, error = goquery.NewDocumentFromReader(resp.Body)
		if error != nil {
			log.Printf("GetEventInfAndRoomList() goquery.NewDocumentFromReader() err=<%s>.\n", error.Error())
			err = fmt.Errorf("goquery.NewDocumentFromReader(): %w", error)
			return
		}

		eventinfo.Event_ID = eventidorfilename
	}
	value, _ := doc.Find("#eventDetail").Attr("data-event-id")
	eventinfo.I_Event_ID, _ = strconv.Atoi(value)

	log.Printf(" eventid=%s (%d)\n", (*eventinfo).Event_ID, eventinfo.I_Event_ID)

	selector := doc.Find(".detail")
	eventinfo.Event_name = selector.Find(".tx-title").Text()
	if eventinfo.Event_name == "" {
		log.Printf("Event not found. Event_ID=%s\n", eventinfo.Event_ID)
		err = fmt.Errorf("event not found. ID=%s", eventinfo.Event_ID)
		return
	}
	eventinfo.Period = selector.Find(".info").Text()
	eventinfo.Period = strings.Replace(eventinfo.Period, "\u202f", " ", -1)
	period := strings.Split(eventinfo.Period, " - ")
	if inputmode == "url" {
		eventinfo.Start_time, _ = time.Parse("Jan 2, 2006 3:04 PM MST", period[0]+" JST")
		eventinfo.End_time, _ = time.Parse("Jan 2, 2006 3:04 PM MST", period[1]+" JST")
	} else {
		eventinfo.Start_time, _ = time.Parse("2006/01/02 15:04 MST", period[0]+" JST")
		eventinfo.End_time, _ = time.Parse("2006/01/02 15:04 MST", period[1]+" JST")
	}

	eventinfo.EventStatus = "BeingHeld"
	if eventinfo.Start_time.After(time.Now()) {
		eventinfo.EventStatus = "NotHeldYet"
	} else if eventinfo.End_time.Before(time.Now()) {
		eventinfo.EventStatus = "Over"
	}

	//      イベントに参加しているルームの数を求めます。
	//      参加ルーム数と表示されているルームの数は違うので注意。ここで取得しているのは参加ルーム数。
	SNoEntry := doc.Find("p.ta-r").Text()
	fmt.Sscanf(SNoEntry, "%d", &eventinfo.NoEntry)
	log.Printf("[%s]\n[%s] [%s] (*event).EventStatus=%s NoEntry=%d\n",
		eventinfo.Event_name,
		eventinfo.Start_time.Format("2006/01/02 15:04 MST"),
		eventinfo.End_time.Format("2006/01/02 15:04 MST"),
		eventinfo.EventStatus, eventinfo.NoEntry)

	return
}
