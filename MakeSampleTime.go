// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi
import (
	"time"
)
/*
MakeSampleTime()
獲得ポイントを取得するタイミングをランダムに返す

5分に一回を前提として、240秒±40秒のように設定する。
*/
func MakeSampleTime(
	cval int, // ex. 240
	cvar int, // ex. 40
) (stm, sts int) {

	st := cval + int(time.Now().UnixNano()%int64(cvar*2)) - cvar

	stm = st / 60
	sts = st % 60

	return stm, sts
}
