package xclip

import "testing"

// 测试 SetClip 方法
func TestSetClip(t *testing.T) {
	err := SetClip("420900589401936109735000084742")
	if err != nil {
		//return
		panic(err)
	}

}
