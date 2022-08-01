/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"fmt"
	"log"
	"time"
	"os"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
	"github.com/juju/persistent-cookiejar"

	"github.com/Chouette2100/srapi"
)

// 設定ファイルを読み込む
//	以下の記事を参考にさせていただきました。
//		【Go初学】設定ファイル、環境変数から設定情報を取得する
//			https://note.com/artefactnote/n/n8c22d1ac4b86
//
func LoadConfig(filePath string, config interface{}) (status int) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("LoadConfig() os.ReadFile() returned err=%s\n", err.Error())
		return -1
	}

	content = []byte(os.ExpandEnv(string(content)))
	//	log.Printf("content=%s\n", content)

	if err := yaml.Unmarshal(content, config); err != nil {
		log.Printf("LoadConfig() yaml.Unmarshal() returned err=%s\n", err.Error())
		return -2
	}

	//	log.Printf("\n")
	//	log.Printf("%+v\n", config)
	//	log.Printf("\n")

	return 0
}

//	ログファイルを作る。
func CreateLogfile(dsc1, dsc2 string) (logfile *os.File) {
	//      ログファイルの設定
	logfilename := os.Args[0]
	if dsc1 != "" {
		logfilename += "_" + dsc1
	}
	logfilename += "_" + time.Now().Format("20060102")
	if dsc2 != "" {
		logfilename += "_" + dsc2
	}
	logfilename += ".txt"
	var err error
	logfile, err = os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("cannnot open logfile: %s\n", logfilename+" "+err.Error())
		os.Exit(1)
	}

	//      log.SetOutput(logfile)
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	return
}

//	HTTPクライアントを作り、cookiejarをセットする。
func CreateNewClient(
	cookiename string,
) (
	client *http.Client,
	jar *cookiejar.Jar,
	err error,
) {
	//	Cookiejarを作る
	//	Filenameは、cookieを保存するファイル名
	jar, err = cookiejar.New(&cookiejar.Options{Filename: cookiename + "_cookies"})
	if err != nil {
		log.Printf("cookiejar.New() returned error %s\n", err.Error())
		return
	}

	//	httpクライアントを作る
	client = &http.Client{}
	client.Jar = jar

	return
}

//	Showroomのサービスにログインし、ユーザIDを取得する。
func LoginShowroom(
	client *http.Client,
	acct string,
	pswd string,
) (
	userid string,
	status int,
) {
	//	SHOWROOMにログインした状態にあるか？
	ud, status := srapi.ApiUserDetail(client)
	if status != 0 {
		return
	}
	//	log.Printf("----------------------------------------------------\n")
	//	log.Printf("%+v\n", ud)
	//	log.Printf("----------------------------------------------------\n")

	if ud.User_id == 0 {
		//	ログインしていない

		//	csrftokenを取得する
		csrftoken := srapi.ApiCsrftoken(client)

		//	SHOWROOMのサービスにログインする。
		var ul srapi.UserLogin
		ul, status = srapi.ApiUserLogin(client, csrftoken, acct, pswd)
		if status != 0 {
			log.Printf("***** ApiUserLogin() returned error. status=%d\n", status)
			return
		} else {
			log.Printf("login status. Ok = %d User_id=%s\n", ul.Ok, ul.User_id)
		}
		userid = ul.User_id

	} else {
		//      ログインしている
		userid = fmt.Sprintf("%d", ud.User_id)
	}
	return

}

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

