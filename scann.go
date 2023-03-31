package main

import (
	"flag"
	"scann/log"
	"scann/ui"
	"scann/zbar"
)

// 支持 -bar 和 -qr 两个参数，分别表示条形码和二维码
var barCode = flag.Bool("bar", false, "bar code scan")
var qrCode = flag.Bool("qr", false, "qr code scan")

var uiMode = flag.Bool("ui", false, "ui mode")

func main() {
	flag.Parse()

	if *uiMode {
		log.I("run in ui mode")
		ui.ShowAppTray()
		return
	} else {
		zbar.SmartScan()
	}
}
