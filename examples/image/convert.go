//go:build external

package imageexample

import (
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func ConvertToWebP(inputPath, outputPath string) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return webp.Encode(outputFile, img, &webp.Options{Lossless: false})
}
