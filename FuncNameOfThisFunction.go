package exsrapi
import (
	"runtime"
	"strings"
)
func FuncNameOfThisFunction() (
	funcname string,
){
	// 現在のスタックから、情報を取得
	skip := 1
	pt, _, _, ok := runtime.Caller(skip)
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
