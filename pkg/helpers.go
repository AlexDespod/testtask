package pkg

import (
	"math/rand"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AlexDespod/testtask/shared"
)

// GetName generate random name for file . Return a path to this file
func GetName(filename string) string {

	filetype := filepath.Ext(filename)
	rand.Seed(time.Now().UnixNano())
	num := rand.Int()

	path := shared.PhotosDir + "\\" + strconv.Itoa(num) + filetype

	return path
}
