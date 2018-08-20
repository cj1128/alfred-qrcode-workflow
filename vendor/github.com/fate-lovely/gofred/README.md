# Gofred

[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://mit-license.org/2016)
[![Build Status](https://travis-ci.org/fate-lovely/gofred.svg?branch=master)](https://travis-ci.org/fate-lovely/gofred)
[![Go Doc](https://godoc.org/github.com/fate-lovely/gofred?status.svg)](https://godoc.org/github.com/fate-lovely/gofred)

A golang library for writing alfred workflow.

## Installation

```go
go get github.com/fate-lovely/gofred
```

## Doc

[Go Doc].

### Example

**make sure to set `Valid` attribute of `Item` and `Mod`, its default value is `false`**.

```go
package main

import (
    "fmt"

    "github.com/fate-lovely/gofred"
)

func main() {
    gofred.AddItem(&gofred.Item{
        Title:    "titl1",
        Subtitle: "subtitle1",
        Arg:      "arg1",
        Valid:    true,
        Icon: &gofred.Icon{
            Type: "filetype",
            Path: "public.png",
        },
    })

    gofred.AddItem(&gofred.Item{
        Title:    "title2",
        Subtitle: "subtitle2",
        Arg:      "arg2",
        Valid:    true,
    })

    json, _ := gofred.JSON()
    fmt.Println(json)
}
// {
//   "items": [
//     {
//       "title": "titl1",
//       "subtitle": "subtitle1",
//       "arg": "arg1",
//       "icon": {
//         "type": "filetype",
//         "path": "public.png"
//       },
//       "valid": true
//     },
//     {
//       "title": "title2",
//       "subtitle": "subtitle2",
//       "arg": "arg2",
//       "valid": true
//     }
//   ]
// }
```

## Reference

- [Alfred JSON Response]

## License

Released under the [MIT license](http://mit-license.org/2016).

[Go Doc]: https://godoc.org/github.com/fate-lovely/gofred
[Alfred JSON Response]: https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
