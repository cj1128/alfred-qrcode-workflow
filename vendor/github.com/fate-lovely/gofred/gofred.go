// Package gofred provides utility functions and structs for writing alfred workflows
package gofred

import (
	"encoding/json"
)

type Response struct {
	// between 0.1 - 5.0 seconds
	Rerun     float64           `json:"rerurn,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	Items     []*Item           `json:"items"`
}

// details can be found at
// https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type Item struct {
	Uid          string `json:"uid,omitempty"`
	Type         string `json:"type,omitempty"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Arg          string `json:"arg,omitempty"`
	Autocomplete string `json:"autocomplete,omitempty"`
	Icon         *Icon  `json:"icon,omitempty"`
	Valid        bool   `json:"valid"`
	Match        string `json:"match,omitempty"`
	Quicklookurl string `json:"quicklookurl,omitempty"`
	Mods         Mods   `json:"mods,omitempty"`
	Text         *Text  `json:"text,omitempty"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type modKey string

const (
	CtrlKey  modKey = "ctrl"
	FnKey    modKey = "fn"
	AltKey   modKey = "alt"
	CmdKey   modKey = "cmd"
	ShiftKey modKey = "shift"
)

type Mods map[modKey]*Mod

type Mod struct {
	Valid     bool              `json:"valid"`
	Arg       string            `json:"arg"`
	Subtitle  string            `json:"subtitle"`
	Icon      *Icon             `json:"icon,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type Text struct {
	Copy      string `json:"copy"`
	Largetype string `json:"largetype"`
}

func (i *Item) AddMod(key modKey, mod *Mod) {
	if i.Mods == nil {
		i.Mods = make(map[modKey]*Mod)
	}
	i.Mods[key] = mod
}

func (r *Response) AddItem(item *Item) {
	r.Items = append(r.Items, item)
}

func (r *Response) ItemLength() int {
	return len(r.Items)
}

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

var DefaultResponse = New()

func New() *Response {
	return &Response{}
}

func AddItem(item *Item) {
	DefaultResponse.AddItem(item)
}

func ClearItems() {
	DefaultResponse.ClearItems()
}

func JSON() (string, error) {
	return DefaultResponse.JSON()
}
