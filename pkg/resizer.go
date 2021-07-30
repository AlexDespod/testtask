package pkg

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func Resize(path string) {

	file, err := os.OpenFile(path, os.O_RDWR, 0666)

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatalf("%s", err)
	}

	m := resize.Resize(100, 0, img, resize.Lanczos3)

	file.Seek(0, 0)

	err = jpeg.Encode(file, m, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

}
