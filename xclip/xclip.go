package xclip

import (
	"io"
	"os/exec"
	"scann/log"
	"time"
)

// SetClip 设置剪贴板
func SetClip(text string) error {

	log.I("set clip text: ", text)

	// 实现 echo "text" | xclip -selection clipboard
	echo := exec.Command("echo", "-n", text)
	xclip := exec.Command("xclip", "-selection", "clipboard")

	echoOut, _ := echo.StdoutPipe()
	defer echoOut.Close()

	xclipIn, _ := xclip.StdinPipe()
	defer xclipIn.Close()

	err := xclip.Start()
	if err != nil {
		return err
	}

	err = echo.Start()
	if err != nil {
		return err
	}

	io.Copy(xclipIn, echoOut)
	xclipIn.Close()

	err = echo.Wait()
	if err != nil {
		return err
	}

	err = xclip.Wait()
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 200)

	return nil
}
