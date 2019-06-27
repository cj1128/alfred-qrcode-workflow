package main

import (
	"bytes"
	"encoding/base64"
	"strings"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"

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
		case "list":
			showHistoryList()
	}
}

func generateQRCode(content string) {
	filename := base64.StdEncoding.EncodeToString([]byte(content))
	path := fmt.Sprintf("/tmp/alfred_qrcode_%s.png", filename)

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

func showHistoryList() {
	files :=	getTmpFileList()
	prefix := "alfred_qrcode_"
	for _, f := range files {
		if strings.Contains(f.Name(), prefix) {
			name := strings.TrimRight(f.Name(), ".png")
			lastIndex := strings.Index(name, prefix) + len(prefix)
			title, err := base64.StdEncoding.DecodeString(name[lastIndex:])
			filepath := "/tmp/" + f.Name()

			if err != nil {
				panic(err)
			}
			gofred.AddItem(&gofred.Item{
				Title: string(title),
				Type:  "file",
				Arg:   filepath,
				Valid: true,
				Icon: &gofred.Icon{
					Type: "file",
					Path: filepath,
				},
			})
		
		}
	}

	response, _ := gofred.JSON()
	fmt.Print(response)
}

func getTmpFileList() []os.FileInfo {
	files, err := ioutil.ReadDir("/tmp")
	if err != nil {
		panic(err)
	}
	return files
}

func scanQRCode(filePath string) {
	cmd := exec.Command("./zbar", filePath)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	err := cmd.Run()

	if err != nil {
		panic(err)
	}

	content := stdout.String()
	if content == "" {
		panic("解码失败")
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
