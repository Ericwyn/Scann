package zbar

import (
	"github.com/Ericwyn/GoTools/shell"
	"scann/log"
	"scann/screenshot"
	"scann/xclip"
)

func SmartScan() {

	filePath, screenShotResult := screenshot.RunScreenshotSyn()
	log.I("filePath: ", filePath, " screenShotResult: ", screenShotResult)

	var analysisResult string
	var err error

	analysisResult, err = AnalysisBarCodePng(filePath)
	if err == nil {
		log.I("barcode analysisResult: ", analysisResult)
		handlerScanResult(analysisResult, "条形码")
		return
	}
	log.E("AnalysisBarCodePng error: ", err)

	analysisResult, err = AnalysisQrCodePng(filePath)
	if err == nil {
		log.I("qrcode analysisResult: ", analysisResult)

		handlerScanResult(analysisResult, "二维码")
		return
	}
	log.E("AnalysisQrCodePng error: ", err)

	log.E("can't parse any result")
	shell.RunShellRes("notify-send", "识别失败", "无法识别到条形码或二维码")
}

func handlerScanResult(analysisResult string, typ string) {
	xclip.SetClip(analysisResult)

	// 执行命令 notify-send "识别结果" "420605049401936109735000157095"
	shell.RunShellRes("notify-send", "识别结果", typ+": "+analysisResult)

	return
}
