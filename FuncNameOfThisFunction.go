// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi
import (
	"runtime"
	"strings"
)
func FuncNameOfThisFunction(
	level int,	// 1: This function, 2: Caller of this function, ...
) (
	funcname string,
){
	// 現在のスタックから、情報を取得
	pt, _, _, ok := runtime.Caller(level)
	if !ok {
		funcname = "unknown"
		return
	}
	
	// ポインターから関数名に変換する
	funcname = runtime.FuncForPC(pt).Name()
	fna := strings.Split(funcname, ".")
	funcname = fna[len(fna)-1]
	return
}
