/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"log"

	"net/http"

	"github.com/Chouette2100/srapi"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。

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
	rooms *[]srapi.RoomFollowing,
) (
	roomafnls []RoomAfnl,
	status int,
) {

	roomafnls = make([]RoomAfnl, 0)
	var afnl srapi.ActiveFanNextLevel
	for _, room := range *rooms {
		afnl, status = srapi.ApiActivefanNextlevel(client, userid, room.Room_id)
		if status != 0 {
			log.Printf("***** ApiActiveFanNextlevel() returned error. status=%d\n", status)
			return
		}
		roomafnls = append(roomafnls, RoomAfnl{room.Room_id, room.Main_name, afnl})
	}
	return
}
