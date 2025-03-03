package exsrapi

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/Chouette2100/srapi/v2"
)

func GetRoominfFromEventOfBR(
	client *http.Client,
	eventid string,
	breg int,
	ereg int,
) (
	roomlistinf *srapi.RoomListInf,
	err error,
) {

	var doc *goquery.Document

	//	inputmode := "url"
	eventidorfilename := eventid

	//	URLからドキュメントを作成します
	_url := "https://www.showroom-live.com/event/" + eventidorfilename
	/*
		doc, err = goquery.NewDocument(_url)
	*/
	resp, errg := http.Get(_url)
	if errg != nil {
		err = fmt.Errorf("http.Get(url): %w", errg)
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	//	bufstr := buf.String()
	//	log.Printf("%s\n", bufstr)

	//	doc, error = goquery.NewDocumentFromReader(resp.Body)
	doc, err = goquery.NewDocumentFromReader(buf)
	if err != nil {
		err = fmt.Errorf("goquery.NewDocumentFromReader(): %w", err)
		return
	}

	cevent_id, exists := doc.Find("#eventDetail").Attr("data-event-id")
	if !exists {
		err = fmt.Errorf("data-event-id not found. Event_ID=%s", eventid)
		return
	}
	ieventid, _ := strconv.Atoi(cevent_id)

	/*
		selector := doc.Find(".detail")
		(*eventinfo).Event_name = selector.Find(".tx-title").Text()
		if (*eventinfo).Event_name == "" {
			log.Printf("Event not found. Event_ID=%s\n", (*eventinfo).Event_ID)
			status = -1
			return
		}
		(*eventinfo).Period = selector.Find(".info").Text()
		eventinfo.Period = strings.Replace(eventinfo.Period, "\u202f", " ", -1)
		period := strings.Split((*eventinfo).Period, " - ")
		if inputmode == "url" {
			(*eventinfo).Start_time, _ = time.Parse("Jan 2, 2006 3:04 PM MST", period[0]+" JST")
			(*eventinfo).End_time, _ = time.Parse("Jan 2, 2006 3:04 PM MST", period[1]+" JST")
		} else {
			(*eventinfo).Start_time, _ = time.Parse("2006/01/02 15:04 MST", period[0]+" JST")
			(*eventinfo).End_time, _ = time.Parse("2006/01/02 15:04 MST", period[1]+" JST")
		}

		(*eventinfo).EventStatus = "BeingHeld"
		if (*eventinfo).Start_time.After(time.Now()) {
			(*eventinfo).EventStatus = "NotHeldYet"
		} else if (*eventinfo).End_time.Before(time.Now()) {
			(*eventinfo).EventStatus = "Over"
		}

		//	イベントに参加しているルームの数を求めます。
		//	参加ルーム数と表示されているルームの数は違うので、ここで取得したルームの数を以下の処理で使うわけではありません。
		SNoEntry := doc.Find("p.ta-r").Text()
		fmt.Sscanf(SNoEntry, "%d", &((*eventinfo).NoEntry))
		//	log.Printf("[%s]\n[%s] [%s] (*event).EventStatus=%s NoEntry=%d\n",
		//		(*eventinfo).Event_name,
		//		(*eventinfo).Start_time.Format("2006/01/02 15:04 MST"),
		//		(*eventinfo).End_time.Format("2006/01/02 15:04 MST"),
		//		(*eventinfo).EventStatus, (*eventinfo).NoEntry)
		//	log.Printf("breg=%d ereg=%d\n", breg, ereg)

		//	eventno, _, _ := SelectEventNoAndName(eventidorfilename)
		//	log.Printf(" eventno=%d\n", eventno)
		//	(*eventinfo).Event_no = eventno
	*/

	eia := strings.Split(eventid, "?")
	bia := strings.Split(eia[1], "=")
	blockid, _ := strconv.Atoi(bia[1])

	ebr, err := srapi.GetEventBlockRanking(client, ieventid, blockid, breg, ereg)
	if err != nil {
		err = fmt.Errorf("GetEventBlockRanking(): %w", err)
		return
	}

	var rli srapi.RoomListInf
	roomlistinf = &rli
	roomlistinf.RoomList = make([]srapi.Room, 0)
	ReplaceString := "/r/"

	for i, br := range ebr.Block_ranking_list {

		var room srapi.Room

		//	roominfo.ID = br.Room_id
		room.Room_id, _ = strconv.Atoi(br.Room_id)

		room.Room_url_key = strings.Replace(br.Room_url_key, ReplaceString, "", -1)
		room.Room_url_key = strings.Replace(room.Room_url_key, "/", "", -1)

		room.Room_name = br.Room_name

		room.Rank = i + 1
		room.Point = br.Point

		roomlistinf.RoomList = append(roomlistinf.RoomList, room)

	}

	//	(*eventinfo).NoRoom = len(*roominfolist)

	//	log.Printf(" GetEventInfAndRoomList() len(*roominfolist)=%d\n", len(*roominfolist))

	return
}
