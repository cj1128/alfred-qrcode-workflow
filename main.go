package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/fate-lovely/gofred"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	action := os.Args[1]

	defer func() {
		if err := recover(); err != nil {
			handleError(fmt.Sprint(err))
		}
	}()

	switch action {
	case "generate":
		generateQRCode(os.Args[2])

	case "scan":
		scanQRCode(os.Args[2])
	}
}

func generateQRCode(content string) {
	md5Hash := md5.Sum([]byte(content))
	path := fmt.Sprintf("/tmp/alfred_qrcode_%s.png", hex.EncodeToString(md5Hash[:]))

	if !fileExists(path) {
		err := qrcode.WriteFile(content, qrcode.Medium, 256, path)
		if err != nil {
			panic(err)
		}
	}

	gofred.AddItem(&gofred.Item{
		Title: content,
		Type:  "file",
		Arg:   path,
		Valid: true,
		Icon: &gofred.Icon{
			Type: "file",
			Path: path,
		},
	})

	response, _ := gofred.JSON()
	fmt.Print(response)
}

func scanQRCode(filePath string) {

}

func handleError(msg string) {
	gofred.AddItem(&gofred.Item{
		Title:    "Error Occurred",
		Subtitle: msg,
		// hide icon
		Icon: &gofred.Icon{
			Path: " ",
		},
	})

	json, _ := gofred.JSON()
	fmt.Println(json)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
