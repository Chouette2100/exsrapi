// Copyright © 2024 chouette.21.00@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"bufio"
	"fmt"
	//	"html"
	"log"
	"os"
	//	"sort"
	//	"strconv"
	"strings"
	//	"time"
	//	"github.com/PuerkitoBio/goquery"
	//	"github.com/dustin/go-humanize"
	//	"html/template"
	//	"net/http"
	//	"database/sql"
	//	"github.com/Chouette2100/exsrapi"
	//	"github.com/Chouette2100/srdblib"
)

// データ取得対象となるルームのポイント基準値を設定する。
func SetThdata(eventinf *Event_Inf, thdata *Thdata) (err error) {

	//	イベントIDあるいはイベント名が特定の条件を満たす場合は、その条件に対応する基準値を設定する
	//      word1 word2 thinit thdelta　　イベントIDにword1が含まれ、イベントネームにword2が含まれるとき
	//      word1 None  thinit thdelta　　イベントIDにword1が含まれるとき
	//      None  word2 thinit thdelta　　イベント名にword2が含まれるとき
	//      None  None  thinit thdelta    既出の条件に一致するものがないとき（デフォルト値）

	eventinf.Thinit = 50
	eventinf.Thdelta = 5
	for _, v := range thdata.Ptdlist {
		if v.Elm[0] == "None" || strings.Contains(eventinf.Event_ID, v.Elm[0]) {
			if v.Elm[1] == "None" || strings.Contains(eventinf.Event_name, v.Elm[1]) {
				eventinf.Thinit = v.Thinit
				eventinf.Thdelta = v.Thdelta
				break
			}
		}
	}
	return
}

// データ取得対象とするルームの基準ポイント
type Ptdata struct {
	Elm     [2]string // Elm[0] 対象となるイベントID（の一部）、#lm[1] 対象となるイベント名（の一部）
	Thinit  int       // 基準ポイントの初期値
	Thdelta int       // Thinit * Thdelta * イベント開始後の経過時間(hour) を基準ポイントとする
}

type Thdata struct {
	Hh      int
	From    int
	To      int
	Ptdlist []Ptdata
}

// 実行パラメータとデータ取得対象となるルームのポイント基準値を読み込む
// Thdata
// Hh int, 開催中のイベントとhh時間以内に始まるイベントをデータ取得対象とする。
// From int, 順位がfromからtoまでの範囲をデータ取得の対象とする
// To int,
// Ptdlist *[]Ptdata,　データ取得対象とするポイント基準
func ReadThdata() (
	thdata *Thdata,
	err error,
) {
	thdata = new(Thdata)

	var pt Ptdata
	//	pt.Elm[0] = "block_id=0"
	//	pt.Elm[1] = "None"
	log.Printf("pt=%+v\n", pt)
	ptlist := make([]Ptdata, 0)

	//	thpoint.txt ファイルから設定値を読み込む
	var file *os.File
	file, err = os.Open("thpoint.txt")
	if err != nil {
		err = fmt.Errorf("ReadThpoint() err:%w", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	first := true
	n := 0
	for scanner.Scan() {
		buf := scanner.Text()
		if buf[0] == '#' {
			continue
		}
		if first {
			first = false
			n, err = fmt.Sscanf(buf, "%d%d%d\n", &thdata.Hh, &thdata.From, &thdata.To)
			if err != nil {
				err = fmt.Errorf("Scanf() err: %w", err)
				return
			}
			if n != 3 {
				//	入力データの形式が違っているか、先頭が"# "でコメントである。
				err = fmt.Errorf("Scanf() n= %d", n)
				return
			}
		} else {
			n, err = fmt.Sscanf(buf, "%s%s%d%d\n", &pt.Elm[0], &pt.Elm[1], &pt.Thinit, &pt.Thdelta)
			if err != nil {
				err = fmt.Errorf("Scanf() err: %w", err)
				return
			}
			if n != 4 {
				//	入力データの形式が違っているか、先頭が"# "でコメントである。
				log.Printf("Scanf() n= %d", n)
				continue
			}
			ptlist = append(ptlist, pt)
		}
	}
	if err = scanner.Err(); err != nil {
		err = fmt.Errorf("scanner.Err() err: %w", err)
		return
	}

	if len(ptlist) == 0 {
		err = fmt.Errorf("ReadThpoint() err:ptlist is empty")
		return
	}
	thdata.Ptdlist = ptlist
	return
}
