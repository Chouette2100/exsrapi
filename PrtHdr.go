package exsrapi

import (
	"log"
)

var Hdrspc = ""

func PrtHdr() (
	fncname string,
) {
	fncname = FuncNameOfThisFunction(2) + "()"
	log.Printf("%s****< %s ****", Hdrspc, fncname)
	Hdrspc += "  "
	return
}
