# Go Alfred

[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://mit-license.org/2016)
[![Build Status](https://travis-ci.org/fate-lovely/go-alfred.svg?branch=master)](https://travis-ci.org/fate-lovely/go-alfred)
[![Go Doc](https://godoc.org/github.com/fate-lovely/go-alfred?status.svg)](https://godoc.org/github.com/fate-lovely/go-alfred)

a golang library for writing alfred workfl

## Installation

```go
go get github.com/fate-lovely/go-alfred
```

## Usage

Read the [doc](https://godoc.org/github.com/fate-lovely/go-alfred)

### Add items and get JSON / XML response strings

```go
import "github.com/fate-lovely/go-alfred"

func main() {
  item := alfred.Item{
    Title:    "this is title",
    Subtitle: "This is subtitle",
    Arg:      "this is arg",
    Icon: alfred.Icon{
      Type: "filetype",
      Path: "public.png",
    },
  }
  alfred.AddItem(item)

  item.Title = "this is title2"
  item.Subtitle = "this is subtitle2"
  item.Arg = "this is arg2"
  item.Icon.Path = "icon.png"

  alfred.AddItem(item)

  jsonStr, err := alfred.JSON()
  // xmlStr, err := res.XML()
}
```

### Use Response struct
```go
import "github.com/fate-lovely/go-alfred"

func main() {
  res := alfred.Response{}
  item := alfred.Item{
    Title:    "this is title",
    Subtitle: "This is subtitle",
    Icon: alfred.Icon{
      Type: "file",
      Path: "public.png",
    },
  }
  res.AddItem(item)
  jsonStr, err := res.JSON()
  // xmlStr, err := res.XML()
}
```

## License

Released under the [MIT license](http://mit-license.org/2016).

