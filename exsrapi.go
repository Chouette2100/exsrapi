/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (

	"log"
	"os"

	"gopkg.in/yaml.v2"

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

	log.Printf("\n")
	log.Printf("%+v\n", config)
	log.Printf("\n")

	return 0
}
