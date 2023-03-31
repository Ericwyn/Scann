package zbar

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"scann/log"
)

// AnalysisQrCodePng 传入一个 png 图片的地址，返回二维码的内容
func AnalysisQrCodePng(pngPath string) (string, error) {

	// open and decode image file
	file, err := os.Open(pngPath)
	if err != nil {
		log.E("AnalysisQrCodePng open file error: ", err, " ,path: ", pngPath)
		return "", err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.E("AnalysisQrCodePng decode file error: ", err, " ,path: ", pngPath)
		return "", err
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.E("AnalysisQrCodePng NewBinaryBitmapFromImage error: ", err, " ,path: ", pngPath)
		return "", err
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)

	if err != nil {
		log.E("AnalysisQrCodePng Decode error: ", err, " ,path: ", pngPath)
		return "", err
	}

	return fmt.Sprint(result), nil

}
