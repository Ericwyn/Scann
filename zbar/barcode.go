package zbar

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"scann/log"
)

// AnalysisBarCodePng 传入一个 png 图片的地址，返回条形码的内容
func AnalysisBarCodePng(pngPath string) (string, error) {
	//return ""

	// open and decode image file
	file, err := os.Open(pngPath)
	if err != nil {
		log.E("AnalysisBarCodePng open file error: ", err, " ,path: ", pngPath)
		return "", err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.E("AnalysisBarCodePng decode file error: ", err, " ,path: ", pngPath)
		return "", err
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.E("AnalysisBarCodePng NewBinaryBitmapFromImage error: ", err, " ,path: ", pngPath)
		return "", err
	}

	// decode image
	qrReader := oned.NewCode128Reader()
	result, err := qrReader.Decode(bmp, nil)

	if err != nil {
		log.E("AnalysisBarCodePng Decode error: ", err, " ,path: ", pngPath)
		return "", err
	}

	return fmt.Sprint(result), nil

}
