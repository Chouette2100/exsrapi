// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
/*
Ver.0.0.1 ConfigのFncを削除する。

Ver.0.1.0 ApiLiveCurrentUser()の引数roomidをstringとしたことへ対応する。ApiLiveCurrentUser()実行時のroomidにRoomid[0]を渡す。
Ver.1.0.0 下位の関数で戻り値をstatusかｒerrに変更したことに対応する。

*/
package exsrapi

import (
	"fmt"
	"log"

	"net/http"

	"github.com/juju/persistent-cookiejar"

	"github.com/Chouette2100/srapi/v2"
)

/*
ファンレベルの達成状況を取得する

	使用しているSHOWROOMのAPI
		srapi.ApiLiveCurrentUser()
		srapi.CsrfToken()
		srapi.UserLogin()
		srapi.ApiActivefanNextlevel()

*/
func TestOfApiActivefanNextlevel(filename string) (err error) {

	//	設定情報
	type Config struct {
		SR_acct string   //	SHOWROOMのアカウント名
		SR_pswd string   //	SHOWROOMのパスワード
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
		return err
	}
	//	すべての処理が終了したらcookiejarを保存する。
	defer jar.Save()

	//	httpクライアントを作る
	client := &http.Client{}
	client.Jar = jar

	//	SHOWROOMにログインした状態にあるか？
	lcu, err := srapi.ApiLiveCurrentUser(client, config.Roomid[0])
	if err != nil {
		err = fmt.Errorf("srapi.ApiLiveCurrentUser: %w", err)
		return err
	}
	//	log.Printf("----------------------------------------------------\n")
	//	log.Printf("%+v\n", lcu)
	//	log.Printf("----------------------------------------------------\n")

	userid := ""
	if lcu.User_id == 0 {
		//	ログインしていない

		//	csrftokenを取得する
		csrftoken, err := srapi.ApiCsrftoken(client)
		if err != nil {
			err = fmt.Errorf("srapi.ApiCsrftoken: %w", err)
			return err
		}

		//	SHOWROOMのサービスにログインする。
		ul, err := srapi.ApiUserLogin(client, csrftoken, config.SR_acct, config.SR_pswd)
		if err != nil {
			err = fmt.Errorf("srapi.ApiUserLogin: %w", err)
			return err
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
		afnl, err := srapi.ApiActivefanNextlevel(client, userid, roomid)
		if err != nil {
			err = fmt.Errorf("srapi.ApiActivefanNextlevel: %w", err)
			return err
		}
		fmt.Printf("current level = %d\n", afnl.Level)
		fmt.Printf("next level =    %d\n", afnl.Next_level.Level)
		for _, c := range afnl.Next_level.Conditions {
			fmt.Printf("%s\n", c.Label)
			for _, cd := range c.Condition_details {
				fmt.Printf("  %-12s (目標)%5d %-10s (実績)%5d %-10s\n", cd.Label, cd.Goal, cd.Unit, cd.Value, cd.Unit)
			}
		}
	}
	return nil
}
