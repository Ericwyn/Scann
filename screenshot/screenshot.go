package screenshot

import (
	"fmt"
	"github.com/Ericwyn/GoTools/shell"
	"math"
	"os"
	"scann/log"
	"time"
)

// var OCRTextTempPath = ""
var ScrPng = "/tmp/go-scann-screenshot-temp.png"

// RunScreenshotSyn 截图，返回图片的地址
// linux 下的实现，使用 gnome-screenshot
func RunScreenshotSyn() (string, bool) {
	shell.RunShellRes("gnome-screenshot", "-a", "-f", ScrPng)

	// 获取文件修改时间，看看是不是1s内修改的
	modTime := getFileModTime(ScrPng)

	if math.Abs((float64)(modTime-time.Now().Unix())) > 10 {
		fmt.Println("文件未修改，不做识别")
		return "", false
	}

	return ScrPng, true
}

func getFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.E("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.E("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
