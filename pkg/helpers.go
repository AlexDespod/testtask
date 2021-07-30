package pkg

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetName(filename string) string {
	filetype := (strings.Split(filename, "."))[1]

	rand.Seed(time.Now().UnixNano())

	name := strconv.Itoa(rand.Int())

	dir, _ := os.Getwd()

	path := dir + "\\photos\\" + name + "." + filetype

	return path
}
