/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"fmt"
	"log"

	"net/http"

	"github.com/juju/persistent-cookiejar"

	"github.com/Chouette2100/srapi"
)

/*
ファンレベルの達成状況を取得する

	使用しているSHOWROOMのAPI
		srapi.ApiLiveCurrentUser()
		srapi.CsrfToken()
		srapi.UserLogin()
		srapi.ApiActivefanNextlevel()

*/
func TestOfApiActivefanNextlevel(filename string) {

	//	設定情報
	type Config struct {
		SR_acct string   //	SHOWROOMのアカウント名
		SR_pswd string   //	SHOWROOMのパスワード
		Fnc     string   //	ファンレベルを知りたいファンコミュニティのID
		Roomid  []string //	ファンレベルを知りたい配信ルームのルームID
	}
	var config Config
	//	config.Roomid = make([]string, 0)

	/*
設定ファイルの例（１）
sr_acct: xxxxxxxxxx
sr_pswd: yyyyyyyyyy
roomid:
- "111111"
- "222222"

設定ファイルの例（２） アカウントとパスワードを環境変数で与える
sr_acct: ${SRACCT}
sr_pswd: ${SRPSWD}
roomid:
- "111111"
- "222222"
- "333333"

	 */
	//	設定ファイルを読み込む
	LoadConfig(filename, &config)

	//	Cookiejarを作る
	//	Filenameは、cookieを保存するファイル名
	jar, err := cookiejar.New(&cookiejar.Options{Filename: config.SR_acct + "_cookies"})
	if err != nil {
		log.Printf("cookiejar.New() returned error %s\n", err.Error())
		return
	}
	//	すべての処理が終了したらcookiejarを保存する。
	defer jar.Save()

	//	httpクライアントを作る
	client := &http.Client{}
	client.Jar = jar

	//	SHOWROOMにログインした状態にあるか？
	lcu, status := srapi.ApiLiveCurrentUser(client, 75721)
	if status != 0 {
		return
	}
	log.Printf("----------------------------------------------------\n")
	log.Printf("%+v\n", lcu)
	log.Printf("----------------------------------------------------\n")

	userid := ""
	if lcu.User_id == 0 {
		//	ログインしていない

		//	csrftokenを取得する
		csrftoken := srapi.ApiCsrftoken(client)

		//	SHOWROOMのサービスにログインする。
		ul, status := srapi.ApiUserLogin(client, csrftoken, config.SR_acct, config.SR_pswd)
		if status != 0 {
			log.Printf("***** ApiUserLogin() returned error. status=%d\n", status)
			return
		} else {
			log.Printf("login status. Ok = %d User_id=%d\n", ul.Ok, lcu.User_id)
		}
		userid = ul.User_id
	} else {
		//	ログインしている
		userid = fmt.Sprintf("%d", lcu.User_id)
	}

	for _, roomid := range config.Roomid {
		//		ファンレベルの詳細を知る。
		log.Printf("********************************************************************************\n")
		afnl, status := srapi.ApiActivefanNextlevel(client, userid, roomid)
		if status != 0 {
			log.Printf("***** ApiActiveFanNextlevel() returned error. status=%d\n", status)
			return
		}
		log.Printf("current level = %d\n", afnl.Level)
		log.Printf("next level =    %d\n", afnl.Next_level.Level)
		for _, c := range afnl.Next_level.Conditions {
			log.Printf("%s\n", c.Label)
			for _, cd := range c.Condition_details {
				log.Printf("  %-12s (目標)%5d %-10s (実績)%5d %-10s\n", cd.Label, cd.Goal, cd.Unit, cd.Value, cd.Unit)
			}
		}
	}

}
