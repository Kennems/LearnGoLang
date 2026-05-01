package main

import (
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"os"
)

func ConvertToWebP(inputPath, outputPath string) error {
	// 使用imaging包打开输入图像文件
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	// 创建输出文件
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 将图像保存为WebP格式到输出文件
	err = webp.Encode(outputFile, img, &webp.Options{Lossless: false})
	if err != nil {
		return err
	}

	// 返回nil表示转换成功
	return nil
}

//func main() {
//	fmt.Println("begin")
//	err := ConvertToWebP("Demo.jpg", "Demo.webp")
//	fmt.Println("end")
//	if err == nil {
//		fmt.Println("转换成功！")
//	}
//}
