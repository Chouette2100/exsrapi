/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"log"

	"net/http"

	"github.com/juju/persistent-cookiejar"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。

*/

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
