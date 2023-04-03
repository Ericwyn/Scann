package main

import (
	"flag"
	"fmt"
	"scann/zbar"
)

var verFlag = flag.Bool("v", false, "show version")
var versionStr = "1.0.0, 2023-04-03"

func main() {
	flag.Parse()

	if *verFlag {
		fmt.Println(versionStr)
		return
	}

	zbar.SmartScan()
}
