package ui

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"scann/log"
	"scann/zbar"
)

func ShowAppTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	//systray.SetTitle("E")
	systray.SetTooltip("Pretty awesome超级棒")
	barCodeItem := systray.AddMenuItem("条码识别", "条码识别")
	quitItem := systray.AddMenuItem("退出", "退出程序")

	go func() {
		<-quitItem.ClickedCh
		log.I("Requesting quit")
		systray.Quit()
		log.I("Finished quitting")

	}()

	// 另一个 goroutine 来处理其他 item
	go func() {
		for {
			select {
			case <-barCodeItem.ClickedCh:
				log.I("条码识别")
				zbar.SmartScan()
			}
		}
	}()
}

func onExit() {
	// clean up here
}
