/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"fmt"

	"net/http"

	"github.com/Chouette2100/srapi/v2"
)

/*

Ver.1.3.0 CrawlFollow()からGetFollowRooms()への置き換えにともなってrooms []srapi.RoomFollowedとした。

*/


//	配信者のリストから、ファンレベルの達成状況を調べる。
type RoomAfnl struct {
	Room_id   string //	配信者のID
	Main_name string //	配信者の名前
	Afnl      srapi.ActiveFanNextLevel
}

func GetActiveFanNextLevel(
	client *http.Client,
	userid string,
	rooms []srapi.RoomFollowed,
) (
	roomafnls []RoomAfnl,
	err error,
) {

	roomafnls = make([]RoomAfnl, 0)
	var afnl srapi.ActiveFanNextLevel
	for _, room := range rooms {
		afnl, err = srapi.ApiActivefanNextlevel(client, userid, room.Room_id)
		if err != nil {
			err = fmt.Errorf("srapi.ApiActivefanNextlevel: %w", err)
			return nil, err
		}
		roomafnls = append(roomafnls, RoomAfnl{room.Room_id, room.Room_name, afnl})
	}
	return roomafnls, nil
}
