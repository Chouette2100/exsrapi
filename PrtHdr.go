package exsrapi

import (
	"log"
)

var Hdrspc = ""

func PrtHdr() (
	fncname string,
) {
	fncname = FuncNameOfThisFunction() + "()"
	log.Printf("%s****< %s ****", Hdrspc, fncname)
	Hdrspc += "  "
	return
}
