/*!
Copyright © 2022 chouette.21.00@gmail.com
Released under the MIT license
https://opensource.org/licenses/mit-license.php
*/
package exsrapi

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"net/http"

	"github.com/Chouette2100/srapi"
)
/*

星集め、種集めの対象とするルームの一覧を作成する。

Ver.0.0.0
Ver.1.0.0 メソッド ExtrRoomLiveByCtg() の名前を ExtrByCtg() に変更したことに対応する。
Ver.1.1.0 ExtrByCtg() の引数がポインターになったことへ対応する。

*/


//		除外リスト
type ExclList map[int]string //	除外ルームIDリスト

//	除外リストの読み込み
func (el ExclList) Read(
	category string,
	exclfilename string,
) (
	err error,
) {

	log.Printf("%-9s ******************* ExclList.Read() *******************\n", category)

	fileEXCL, err := os.OpenFile(exclfilename, os.O_RDONLY, 0644)
	if err != nil {
		//	除外リストファイルをオープンできなかった場合
		err = fmt.Errorf("os.OpenFile(): %w", err)
		return err
	}
	defer fileEXCL.Close()

	var roomid int
	var roomname string
	for {
		nitem, _ := fmt.Fscanf(fileEXCL, "%d%s\n", &roomid, &roomname)
		//	log.Printf(" nitem =%d err=<%v> url=%s tstr=<%s>\n", nitem, err, url, tstr)
		if nitem != 2 {
			if nitem != 0 {
				//	読み込みに失敗した場合
				err = errors.New("fmt.Fscanf(): unexpected number of items. no=" + fmt.Sprintf("%d", nitem))
				return err
			}
			//	読み込みが終わった場合
			break
		}
		el[roomid] = roomname
		log.Printf("%-9s excl. %-12d %s\n", category, roomid, roomname)
	}
	return nil
}

//	訪問済みルーム情報
type RoomVisit struct {
	Rvlfn     string
	Category     string
	Roomvisit map[int]time.Time
}

//		配信ルーム訪問情報の読み込み
func (r *RoomVisit) Restore(
	category string, //	ジャンル名、ログのヘッダーに使う
	rvlfn string, //	訪問ルームリストファイル名
	aplmin int, //	訪問ルームリストの有効時間(分)
) (
	err error,
) {

	log.Printf("%-9s ******************* RoomVisit.Restore() *******************\n", category)

	r.Rvlfn = rvlfn //	訪問済みルームリストファイル名、SaveRVL()での書き込み事故防止のために使う
	r.Category = category //	カテゴリー、SaveRVL()での書き込み事故防止のために使う

	if len((*r).Roomvisit) != 0 {
		//	すでに訪問済みルームリストがある場合は、それを使う。
		for roomid, rvtime := range (*r).Roomvisit {
			if (*r).Roomvisit[roomid].Before(time.Now().Add(time.Duration(-aplmin) * time.Minute)) {
				//	最後の訪問が一定時間以上前のルームは重複訪問対策用リストから削除する。
				delete((*r).Roomvisit, roomid)
				//	log.Printf("M %20s deleted. (ltime=%s)\n", surl, rvtime.Format("01-02 15:04:05"))
				log.Printf("%-9sM %-12d%s deleted.\n", category, roomid, rvtime.Format("01-02 15:04:05"))
			}
		}
	} else {
		//	訪問済みルームリストがない場合は初めての訪問候補リスト作成なのでファイルから読み込む。
		fileRVL, err := os.OpenFile(rvlfn, os.O_RDONLY, 0644)
		if err != nil {
			//	カテゴリーリストファイルをオープンできなかった場合
			err = fmt.Errorf("os.OpenFile(): %w", err)
			return err
		}
		defer fileRVL.Close()

		var tstr string
		var roomid int
		for {
			nitem, _ := fmt.Fscanf(fileRVL, "%q%d\n", &tstr, &roomid)
			//	log.Printf(" nitem =%d err=<%v> url=%s tstr=<%s>\n", nitem, err, url, tstr)
			if nitem != 2 {
				if nitem != 0 {
					//	読み込みに失敗した場合
					err = errors.New("fmt.Fscanf(): unexpected number of items. no=" + fmt.Sprintf("%d", nitem))
					return err
				}
				//	読み込みが終わった場合
				break
			}
			rvtime, err := time.Parse("2006/01/02 15:04:05 -0700 MST", tstr)
			if err != nil {
				//	日付の解析に失敗した場合
				err = fmt.Errorf("time.Parse(): %w", err)
				return err
			}

			log.Printf("%-9s\"%s\"\t%d\n", category, rvtime.Format("2006/01/02 15:04:05 -0700 MST"), roomid)

			if rvtime.Before(time.Now().Add(time.Duration(-aplmin) * time.Minute)) {
				//	一定時間経過した訪問情報は削除する
				log.Printf("%-9sF %-12d%s deleted.\n", category, roomid, rvtime.Format("01-02 15:04:05"))
				continue
			}
			(*r).Roomvisit[roomid] = rvtime
			//	log.Printf("F %12d used.    (ltime=%s)\n", roomid, rvtime.Format("01-02 15:04:05"))
		}
	}
	return nil
}

//		配信ルーム訪問情報の書き出し
func (r *RoomVisit) Save(
) (
	err error,
) {

	log.Printf("%-9s ******************* RoomVisit.Save() *******************\n", r.Category)

	filervl, err := os.OpenFile(r.Rvlfn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		err = fmt.Errorf("os.OpenFile(): %w", err)
		return err
	}

	//	ルームリストをソートするためにいったんスライスにコピーする。
	type srv struct {
		vtm    time.Time
		roomid int
	}
	var srvl []srv
	for roomid, tm := range r.Roomvisit {
		srvl = append(srvl, srv{tm, roomid})
	}

	//	訪問時間でソートしておかないとひと目で状況が理解できない。
	sort.Slice(srvl, func(i, j int) bool {
		return srvl[i].vtm.Before(srvl[j].vtm)
	})

	//	訪問済みルームリストを書き出す。
	for _, srv := range srvl {
		fmt.Fprintf(filervl, "\"%s\"\t%d\n", srv.vtm.Format("2006/01/02 15:04:05 -0700 MST"), srv.roomid)
		log.Printf("%-9s\"%s\"\t%d\n", r.Category, srv.vtm.Format("2006/01/02 15:04:05 -0700 MST"), srv.roomid)
	}

	filervl.Close()
	return nil
}

//	星集め、種集めの対象とするルームの一覧を作成する。
func MkRoomsForStarCollec(
	client *http.Client, //	cookiejar付きHTTPクライアント
	category string, //	ジャンル名、ログのヘッダーにも使う
	aplmin int, //	訪問後経過時間の最大値(分)
	maxnoroom int, //	訪問候補として選択するルームの最大数
	excllist *ExclList, //	除外するルームのリスト
	roomvisit *map[int]time.Time, //	訪問済みルームリスト
) (
	lives *[]srapi.Live, //	訪問候補ルームリスト
	err error,
) {

	log.Printf("%-9s ******************* MkRoomsForStarCollec() *******************\n", category)

	//	現在配信中のすべてのルームのリストを取得する。
	Roomonlives, err := srapi.ApiLiveOnlives(client)
	if err != nil {
		err = fmt.Errorf("srapi.ApiLiveOnlives(): %w", err)
		return nil, err
	}

	//	指定したカテゴリーのルームのリストを取得する。
	lives_c, err := Roomonlives.ExtrByCtg(category)
	if err != nil {
		err = fmt.Errorf("srapi.ExtrRoomLiveByCtg(): %w", err)
		return nil, err
	}

	//	訪問候補ルームは配信を始めたばかりのルームから選ぶため開始時刻でソートする。
	sort.Sort(lives_c)

	lives = new([]srapi.Live)
	i := 0
	for _, live := range *lives_c {
		//	ジャンル内のすべてのルームについて繰り返す。
		reason, exist := (*excllist)[live.Room_id]
		starttime := time.Unix(live.Started_at, 0)
		if exist {
			//	除外リストに登録されているルームは除外する。
			log.Printf("%-9sB %-12d%s in EL. %s\n", category, live.Room_id, starttime.Format("01-02 15:04:05"), reason)
			continue
		}

		if starttime.Before(time.Now().Add(time.Duration(-aplmin) * time.Minute)) {
			//	配信開始から一定時間経過したルームの訪問情報は削除されており視聴の可否が判断できないので無視する
			log.Printf("%-9sO %-12d%s >%dmin. \n", category, live.Room_id, starttime.Format("01-02 15:04:05"), aplmin)
			continue
		}

		ltime, exist := (*roomvisit)[live.Room_id]
		if exist {
			if starttime.Before(ltime) {
				//	前回の視聴が終わってから今まで配信が続いている
				//	このケースでは星・種をもらえないのは確認済み
				//	delete(roomvisit, shortURL)
				log.Printf("%-9sL %-12d%s Can't get a star twice in one delivery! (ltime=%s)\n", category,
					live.Room_id, starttime.Format("01-02 15:04:06"), ltime.Format("15:04"))
				continue
			}
			if time.Since(ltime) < time.Hour {
				//	前回の視聴が終わってから新たに配信が行われているルーム
				//	ただし前回の視聴からまだ1時間経っていない <== ダメなの？
				log.Printf("%-9sP %-12d Only %5.1fm have passed after last visit. (ltime=%s)\n", category,
					live.Room_id, time.Since(ltime).Minutes(), ltime.Format("15:04:05"))
				//	continue	ひとまずやってみる
			}

		}
		//	取得したデータを戻り値（用の変数）に格納します。
		*lives = append(*lives, live)
		i++
		if i >= maxnoroom {
			break
		}

	}

	return lives, nil

}
