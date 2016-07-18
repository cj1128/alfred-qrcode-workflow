/*
* @Author: CJ Ting
* @Date:   2016-07-05 11:15:57
* @Last Modified by:   CJ Ting
* @Last Modified time: 2016-07-18 17:44:50
 */

// Package alfred provides utility functions and structs for writing alfred workflows
package alfred

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

type Response struct {
	XMLName xml.Name `json:"-" xml:"items"`
	Items   []*Item  `json:"items"`
}

type Item struct {
	XMLName      xml.Name `json:"-" xml:"item"`
	Uid          string   `json:"uid,omitempty" xml:"uid,attr,omitempty"`
	Type         string   `json:"type,omitempty" xml:"type,attr"`
	Title        string   `json:"title" xml:"title"`
	Subtitle     string   `json:"subtitle" xml:"subtitle"`
	Arg          string   `json:"arg" xml:"arg,attr"`
	Autocomplete string   `json:"autocomplete" xml:"autocomplete,attr"`
	Icon         Icon     `json:"icon" xml:"icon"`
	Valid        bool     `json:"valid,omitempty" xml:"valid,attr"`
	Quicklookurl string   `json:"quicklookurl,omitempty" xml:"quicklookurl"`
	Mods         Mods     `json:"mods,omitempty"`
	Text         Text     `json:"text"`
}

type Icon struct {
	Type string `json:"type" xml:"type,attr"`
	Path string `json:"path" xml:",chardata"`
}

type Mods map[string]Mod

func (mods Mods) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{}

	nameKey := xml.Name{"", "key"}
	nameSubtitle := xml.Name{"", "subtitle"}
	nameValid := xml.Name{"", "valid"}
	nameArg := xml.Name{"", "arg"}

	keys := []string{"cmd", "ctrl", "shift", "fn", "alt"}
	for _, key := range keys {
		mod := mods[key]
		validStr := "no"
		if mod.Valid {
			validStr = "yes"
		}

		t := xml.StartElement{
			Name: xml.Name{"", "mod"},
			Attr: []xml.Attr{
				{nameKey, key},
				{nameSubtitle, mod.Subtitle},
				{nameValid, validStr},
				{nameArg, mod.Arg},
			},
		}
		tokens = append(tokens, t, xml.EndElement{t.Name})
	}

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

type Mod struct {
	// Key      string `json:"-" xml:"key,attr"`
	Valid    bool   `json:"valid" xml:"valid,attr"`
	Arg      string `json:"arg" xml:"arg,attr"`
	Subtitle string `json:"subtitle" xml:"subtitle,attr"`
}

type Text struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

func (text Text) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{}

	nameText := xml.Name{"", "text"}
	nameType := xml.Name{"", "type"}

	copyToken := xml.StartElement{
		Name: nameText,
		Attr: []xml.Attr{
			{nameType, "copy"},
		},
	}
	tokens = append(tokens, copyToken, xml.CharData(text.Copy), xml.EndElement{Name: copyToken.Name})

	largetypeToken := xml.StartElement{
		Name: nameText,
		Attr: []xml.Attr{
			{nameType, "largetype"},
		},
	}

	tokens = append(tokens, largetypeToken, xml.CharData(text.Largetype), xml.EndElement{Name: copyToken.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (i *Item) AddMod(key string, mod Mod) {
	if i.Mods == nil {
		i.Mods = make(map[string]Mod)
	}
	i.Mods[key] = mod
}

var defaultResponse = Response{}

// add Item to response
func (r *Response) AddItem(item Item) {
	r.Items = append(r.Items, &item)
}

func (r *Response) ItemLength() int {
	return len(r.Items)
}

// clear all items
func (r *Response) ClearItems() {
	r.Items = nil
}

func (r *Response) JSON() (string, error) {
	bytes, err := json.MarshalIndent(*r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *Response) XML() (string, error) {
	bytes, err := xml.MarshalIndent(*r, "", "  ")
	if err != nil {
		return "", err
	}
	result := xml.Header + string(bytes)

	// Replace valid="true/false" -> valid="yes/no"
	result = strings.Replace(result, `valid="true"`, `valid="yes"`, -1)
	result = strings.Replace(result, `valid="false"`, `valid="no"`, -1)
	return result, nil
}

// top level functions
func AddItem(item Item) {
	defaultResponse.AddItem(item)
}

func ItemLength() int {
	return defaultResponse.ItemLength()
}

func ClearItems() {
	defaultResponse.ClearItems()
}

func JSON() (string, error) {
	return defaultResponse.JSON()
}

func XML() (string, error) {
	return defaultResponse.XML()
}
