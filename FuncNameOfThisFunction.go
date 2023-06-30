package exsrapi
import (
	"runtime"
)
func FuncNameOfThisFunction() (
	funcname string,
){
	// 現在のスタックから、情報を取得
	pt, _, _, ok := runtime.Caller(2)
	if !ok {
		funcname = "unknown"
		return
	}
	
	// ポインターから関数名に変換する
	funcname = runtime.FuncForPC(pt).Name()
	return
}
