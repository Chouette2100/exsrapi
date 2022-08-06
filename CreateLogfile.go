/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

/*

Ver.0.0.0
Ver.1.0.0 LoginShowroom()の戻り値 status を err に変更する。
Ver.-.-.- exsrapi.go から分離する。

*/

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
