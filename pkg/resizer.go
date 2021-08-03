package pkg

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"path/filepath"

	"os"

	"github.com/nfnt/resize"
)

//Resize can resize jpg or png formats
func Resize(path string, width, height uint) error {

	file, err := os.OpenFile(path, os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	filetype := filepath.Ext(path)

	switch filetype {
	case ".png":
		err = resizePNG(file, width, height)
	case ".jpg", ".jpeg":
		err = resizeJPG(file, width, height)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func resizeJPG(file *os.File, width, height uint) error {
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	m := resize.Resize(width, height, img, resize.Lanczos3)

	file.Seek(0, 0)

	err = jpeg.Encode(file, m, nil)

	if err != nil {
		return err
	}
	return nil
}

func resizePNG(file *os.File, width, height uint) error {
	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	m := resize.Resize(width, height, img, resize.Lanczos3)

	file.Seek(0, 0)

	err = png.Encode(file, m)

	if err != nil {
		return err
	}
	return nil
}
