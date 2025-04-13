// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"fmt"
	"log"

	"net/http"

	"github.com/Chouette2100/srapi/v2"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。

*/

//	Showroomのサービスにログインし、ユーザIDを取得する。
func LoginShowroom(
	client *http.Client,
	acct string,
	pswd string,
) (
	userid string,
	err error,
) {
	//	SHOWROOMにログインした状態にあるか？
	ud, err := srapi.ApiUserDetail(client)
	if err != nil {
		log.Printf(" err = %+v\n", err)
		return "0", err
	}
	//	log.Printf("----------------------------------------------------\n")
	//	log.Printf("%+v\n", ud)
	//	log.Printf("----------------------------------------------------\n")

	if ud.User_id == 0 {
		//	ログインしていない

		//	csrftokenを取得する
		csrftoken, err := srapi.ApiCsrftoken(client)
		if err != nil {
			err = fmt.Errorf("srapi.ApiCsrftoken: %w", err)
			return "0", err
		}

		//	SHOWROOMのサービスにログインする。
		var ul srapi.UserLogin
		ul, err = srapi.ApiUserLogin(client, csrftoken, acct, pswd)
		if err != nil {
			err = fmt.Errorf("srapi.ApiUserLogin: %w", err)
			return "0", err
		} else {
			log.Printf("login status. Ok = %d User_id=%s\n", ul.Ok, ul.User_id)
		}
		userid = ul.User_id

	} else {
		//      ログインしている
		userid = fmt.Sprintf("%d", ud.User_id)
	}
	return userid, nil

}
