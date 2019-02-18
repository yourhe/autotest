package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"math"

	"github.com/google/uuid"

	"github.com/nfnt/resize"
)

const DEFAULT_MAX_WIDTH float64 = 320
const DEFAULT_MAX_HEIGHT float64 = 240

// 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DEFAULT_MAX_WIDTH/float64(srcWidth), DEFAULT_MAX_HEIGHT/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// 生成缩略图
func MakeThumbnail(bs []byte) ([]byte, error) {
	r := bytes.NewReader(bs)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// // 需要保存的文件
	// imgfile, _ := os.Create(savePath)
	// defer imgfile.Close()

	// 以PNG格式保存文件
	buf := bytes.NewBuffer(nil)
	err = png.Encode(buf, m)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

//NewID uuid
func NewID() string {
	return uuid.New().String()
}
