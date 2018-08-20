package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/fate-lovely/gofred"
	qrcode "github.com/skip2/go-qrcode"
	qrcode2 "github.com/tuotoo/qrcode"
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
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	qrmatrix, err := qrcode2.Decode(file)

	var content string

	if err != nil {
		content = "Invlaid QR Code"
	} else {
		content = qrmatrix.Content
	}

	gofred.AddItem(&gofred.Item{
		Title:    content,
		Subtitle: "press Enter to copy",
		Valid:    true,
		Arg:      content,
		Text: &gofred.Text{
			Largetype: content,
		},
		Icon: &gofred.Icon{
			Path: " ",
		},
	})

	response, _ := gofred.JSON()
	fmt.Print(response)
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
