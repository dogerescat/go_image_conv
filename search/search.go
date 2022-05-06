package search

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SearchImgFile(dir string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return err
		}
		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}
		slice := strings.Split(path, ".")
		switch filepath.Ext(path) {
		case ".jpeg", ".jpg":
			out, err := os.Create(slice[0] + ".png")
			if err != nil {
				return err
			}
			png.Encode(out, img)
			os.Remove(path)
		case ".png":
			out, err := os.Create(slice[0] + ".jpeg")
			if err != nil {
				return err
			}
			jpeg.Encode(out, img, &jpeg.Options{})
			os.Remove(path)
		}
		return nil
	})
	return err
}
