package exsrapi
import (
	"os"
	"strconv"
	"golang.org/x/sys/unix"
)
// 既存のロックファイルをチェックし、有効なプロセスが実行中かを確認
func CheckExistingLock(lockFilePath string) bool {
	// content, err := ioutil.ReadFile(lockFilePath)
	content, err := os.ReadFile(lockFilePath)
	if err != nil {
			// ファイルが存在しない場合はロックなし
			return false
	}

	// ロックファイルからPIDを読み取る
	pid, err := strconv.Atoi(string(content))
	if err != nil {
			// PIDが読み取れない場合は古いロックファイルと判断して削除
			os.Remove(lockFilePath)
			return false
	}

	// プロセスが実行中かチェック
	// Linuxでは /proc/{pid} にアクセスする方法もあります
	process, err := os.FindProcess(pid)
	if err != nil {
			// プロセスが見つからない場合はロックファイルを削除
			os.Remove(lockFilePath)
			return false
	}

	// プロセスにシグナル0を送信してチェック（実際にはシグナルは送信されない）
	// err = process.Signal(syscall.Signal(0))
	err = process.Signal(unix.Signal(0))
	if err != nil {
			// プロセスが存在しない場合はロックファイルを削除
			os.Remove(lockFilePath)
			return false
	}

	// プロセスは実行中
	return true
}