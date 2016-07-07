/*
* @Author: CJ Ting
* @Date:   2016-07-05 11:13:55
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-07 19:12:26
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	alfred "github.com/fate-lovely/go-alfred"
	qrcode "github.com/skip2/go-qrcode"
)

var meta map[string]string

const meatFilename = "meta.json"
const codesDir = "qrcodes"

var item = alfred.Item{
	Type: "file",
	Icon: alfred.Icon{
		Path: "icon.png",
	},
}

func init() {
	createDirIfNotExists(codesDir)
	readMeta()
}

func main() {
	action := os.Args[1]

	switch action {
	case "add":
		if os.Args[2] == "" {
			listAllQRCodes()
		} else {
			generateQRCode(os.Args[2])
		}
	case "clear":
		clearAllQRCodes()
	}
}

type MetaKeys []string

func (s MetaKeys) Len() int {
	return len(s)
}

func (s MetaKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MetaKeys) Less(i, j int) bool {
	return getTimestampFromMeatKey(s[i]) > getTimestampFromMeatKey(s[j])
}

func listAllQRCodes() {
	var metaKeys []string
	for k, _ := range meta {
		metaKeys = append(metaKeys, k)
	}

	sort.Sort(MetaKeys(metaKeys))

	for _, key := range metaKeys {
		content := getContentFromMeatKey(key)
		path := meta[key]
		item.Title = content
		item.Arg = path
		alfred.AddItem(item)
	}

	response, err := alfred.JSON()
	if err == nil {
		fmt.Print(response)
	} else {
		panic(err)
	}
}

func clearAllQRCodes() {
	meta = nil
	saveMeta()

	err := os.RemoveAll(codesDir)
	if err != nil {
		panic(err)
	}

	item.Title = "Clear all QR Codes successfully!"
	item.Type = ""
	item.Icon = alfred.Icon{}
	alfred.AddItem(item)
	response, err := alfred.JSON()
	fmt.Print(response)
}

func generateQRCode(content string) {

	// if begins with @, it's a url, add http://
	if string(content[0]) == "@" {
		content = fmt.Sprintf("http://%s", content[1:])
	}

	timestamp := time.Now().Unix()
	path := fmt.Sprintf("./%s/qrcode_%d.png", codesDir, timestamp)
	absPath, _ := filepath.Abs(path)

	metaKey := fmt.Sprintf("%s,%d", content, timestamp)
	meta[metaKey] = absPath
	saveMeta()

	err := qrcode.WriteFile(content, qrcode.Medium, 256, path)
	if err != nil {
		panic(err)
	}

	item.Title = content
	item.Arg = absPath
	alfred.AddItem(item)

	response, err := alfred.JSON()
	if err != nil {
		panic(err)
	}
	fmt.Print(response)
}

func readMeta() {
	if !checkExists(meatFilename) {
		meta = make(map[string]string)
		return
	}

	bytes, err := ioutil.ReadFile(meatFilename)
	if err == nil {
		json.Unmarshal(bytes, &meta)
		if meta == nil {
			meta = make(map[string]string)
		}
	} else {
		panic(err)
	}
}

func saveMeta() {
	bytes, err := json.Marshal(meta)
	if err == nil {
		err := ioutil.WriteFile(meatFilename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func checkExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDirIfNotExists(path string) error {
	if !checkExists(path) {
		err := os.Mkdir(path, 0744)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTimestampFromMeatKey(key string) int64 {
	i := strings.LastIndex(key, ",")
	res, _ := strconv.ParseInt(key[i+1:], 10, 64)
	return res
}

func getContentFromMeatKey(key string) string {
	i := strings.LastIndex(key, ",")
	return key[:i]
}
