package main

import (
	"flag"
	"github.com/Ericwyn/GoTools/shell"
	"scann/log"
	"scann/screenshot"
	"scann/xclip"
	"scann/zbar"
)

// 支持 -bar 和 -qr 两个参数，分别表示条形码和二维码
var barCode = flag.Bool("bar", false, "bar code scan")
var qrCode = flag.Bool("qr", false, "qr code scan")

func main() {

	flag.Parse()

	if !*barCode && !*qrCode {
		*barCode = true
		*qrCode = true
		log.I("not set -bar or -qr, will scan both bar code and qr code")
	}

	filePath, screenShotResult := screenshot.RunScreenshotSyn()
	log.I("filePath: ", filePath, " screenShotResult: ", screenShotResult)

	var analysisResult string
	var err error
	if *barCode {
		analysisResult, err = zbar.AnalysisBarCodePng(filePath)
		if err == nil {
			log.I("barcode analysisResult: ", analysisResult)
			handlerScanResult(analysisResult, "条形码")
			return
		}
		log.E("AnalysisBarCodePng error: ", err)
	}

	if *qrCode {
		analysisResult, err = zbar.AnalysisQrCodePng(filePath)
		if err == nil {
			log.I("qrcode analysisResult: ", analysisResult)

			handlerScanResult(analysisResult, "二维码")
			return
		}
		log.E("AnalysisQrCodePng error: ", err)
	}

	log.E("can't parse any result")
	shell.RunShellRes("notify-send", "识别失败", "无法识别到条形码或二维码")
}

func handlerScanResult(analysisResult string, typ string) {
	xclip.SetClip(analysisResult)

	// 执行命令 notify-send "识别结果" "420605049401936109735000157095"
	shell.RunShellRes("notify-send", "识别结果", typ+": "+analysisResult)

	return
}
