// Copyright © 2025 chouette2100@gmail.com
// Released under the MIT license
// https://opensource.org/licenses/mit-license.php
package exsrapi

import (
	"log"
)

var Hdrspc = ""

func PrtHdr() (
	fncname string,
) {
	fncname = FuncNameOfThisFunction(2) + "()"
	log.Printf("%s>>>>>>>>>>>>>>>>>> %s", Hdrspc, fncname)
	Hdrspc += "  "
	return
}
