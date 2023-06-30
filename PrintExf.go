package exsrapi

import (
	"log"
	"time"
)

func PrintExf(comment ...string) func() {
	start := time.Now()
	return func() {
		if len(Hdrspc) > 2 {
			Hdrspc = Hdrspc[2:]
		} else {
			Hdrspc = ""
		}
		fmtstr := Hdrspc
		if len(comment) != 0 {
			fmtstr += comment[0]
		}
		fmtstr += "****> "
		for _, c := range comment {
			fmtstr += c
		}
		fmtstr += " %s dt=%10.3fms\n"
		log.Printf(fmtstr, time.Now().Format("2006/01/02 15:04:05"), float64(time.Since(start).Microseconds())/1000.0)
	}
}
