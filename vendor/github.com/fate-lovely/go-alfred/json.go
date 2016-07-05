/*
* @Author: CJ Ting
* @Date:   2016-07-05 11:15:57
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-05 15:02:35
 */

package alfred

import (
	"encoding/json"
)

type Item struct {
	Uid          string `json:"uid"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
	Icon         Icon   `json:"icon"`
	Valid        bool   `json:"valid,omitempty"`
	Quicklookurl string `json:"quicklookurl"`
	Mods         Mod    `json:"mods"`
	Text         struct {
		Copy      string `json:"copy"`
		Largetype string `json:"largetype"`
	} `json:"text"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Mod struct {
	Alt   ModItem `json:"alt"`
	Cmd   ModItem `json:"alt"`
	Shift ModItem `json:"cmd"`
}

type ModItem struct {
	Valid    bool   `json:"valid"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
}

type alfredResponse struct {
	Items []*Item `json:"items"`
}

var items []*Item

// add Item to response
func AddItem(item Item) {
	items = append(items, &item)
}

func JSON() (string, error) {
	var response = alfredResponse{}
	response.Items = items
	bytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
