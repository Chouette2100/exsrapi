/*
!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	//	"errors"
	// "fmt"
	// "log"
	// "os"
	// "strconv"
	// "strings"
	"time"

	"net/http"

	// "github.com/PuerkitoBio/goquery"

	"github.com/Chouette2100/srapi/v2"
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
	Thinit      int
	Thdelta     int
	MaxPoint    int //      DBには該当するものはない（データ中最大のポイント値を意味し、内部的に使用する）
	Gscale      int //      DBのMaxpoint = 構造体の Maxpoint + Gscale
	Achk        int //      1: ブロック、2:ボックス、子ルーム未処理のあいだはそれぞれ +4
	Aclr        int //		gtplの制御のために一時的に使用する

	EventStatus string //   "Over", "BeingHeld", "NotHeldYet"
	Pntbasis    int
	Ordbasis    int
	League_ids  string
	Valid       bool //	データとして有効か？（内部処理で処理の分岐に使う）
	//      Event_no    int
	//      Status          string          //      "Confirmed":    イベント終了日翌日に確定した獲得ポイントが反映されている。
}

// この関数はもともとイベントページをスクレイピングしていましたが、
// イベント情報を取得するためのAPIを使用するように変更されました。

// GetEventinf は、イベント情報を取得します。
//	イベント情報は、Event_Inf構造体に格納されます。
//	イベント情報は、イベントIDを指定して取得します。
func GetEventinf(
	eventid string,
	eventinfo *Event_Inf,
) (
	err error,
) {

	client := &http.Client{}
	var ea *srapi.EventAbstraction
	ea, err = srapi.ApiEventAbstraction(client, eventid)

	*eventinfo = Event_Inf{
		Event_ID: eventid,
		I_Event_ID: ea.EventID,
		Event_name: ea.Title,
		Start_time: time.Unix(int64(ea.EventStartAt), 0),
		End_time: time.Unix(int64(ea.EventEndAt), 0),
	}
	eventinfo.Period = eventinfo.Start_time.Format("2006/01/02 15:04") + " - " + eventinfo.End_time.Format("2006/01/02 15:04")

	return
}
