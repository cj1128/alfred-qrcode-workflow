/*
* @Author: CJ Ting
* @Date:   2016-07-05 11:15:57
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-06 13:51:23
 */

package alfred

import (
	"encoding/json"
	"encoding/xml"
)

type Item struct {
	XMLName      xml.Name           `json:"-" xml:"item"`
	Uid          string             `json:"uid,omitempty" xml:"uid,attr"`
	Type         string             `json:"type" xml:"type,attr"`
	Title        string             `json:"title" xml:"title"`
	Subtitle     string             `json:"subtitle" xml:"subittle"`
	Arg          string             `json:"arg" xml:"arg,attr"`
	Autocomplete string             `json:"autocomplete" xml:"autocomplete,attr"`
	Icon         Icon               `json:"icon" xml:"icon"`
	Valid        bool               `json:"valid,omitempty" xml:"valid,attr"`
	Quicklookurl string             `json:"quicklookurl" xml:"quicklookurl"`
	Mods         map[string]ModItem `json:"mods,omitempty"`
	Text         struct {
		Copy      string `json:"copy" xml:""`
		Largetype string `json:"largetype" xml:""`
	} `json:"text"`
}

func (i *Item) AddMod(key string, modItem ModItem) {
	if i.Mods == nil {
		i.Mods = make(map[string]ModItem)
	}
	i.Mods[key] = modItem
}

type Icon struct {
	Type string `json:"type" xml:"type,attr"`
	Path string `json:"path" xml:",chardata"`
}

// type Mod struct {
// 	Alt   ModItem `json:"alt" xml:"mod"`
// 	Cmd   ModItem `json:"alt" xml:"mod"`
// 	Shift ModItem `json:"cmd" xml:"bbb"`
// 	Fn    ModItem `json:"fn" xml:"ccc"`
// }

type ModItem struct {
	// Key      string `json:"-" xml:"key,attr"`
	Valid    bool   `json:"valid" xml:"valid,attr"`
	Arg      string `json:"arg" xml:"arg,attr"`
	Subtitle string `json:"subtitle" xml:"subtitle,attr"`
}

type alfredResponse struct {
	XMLName xml.Name `json:"-" xml:"items"`
	Items   []*Item  `json:"items"`
}

var items []*Item
var response = alfredResponse{}

// add Item to response
func AddItem(item Item) {
	items = append(items, &item)
}

// clear all items
func ClearItems() {
	items = nil
}

func JSON() (string, error) {
	response.Items = items
	bytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func XML() (string, error) {
	response.Items = items
	bytes, err := xml.Marshal(response)
	if err != nil {
		return "", nil
	}
	return xml.Header + string(bytes), nil
}
