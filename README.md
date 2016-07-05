# Alfred QR Code Workflow

[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://mit-license.org/2016)


alfred workflow to manipulate QR Codes on the fly, uses [json response](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/), currenty only supports [alred v3](https://www.alfredapp.com/). if you are an alfred v2 user, please update, it's definitely worth it.

## Installation

Download [QR-Code.alfredworkflow](https://github.com/fate-lovely/alfred-qrcode-workflow/raw/master/qrcode.alfredworkflow) and import to alfred v3(require Powerpack)

## Usage

### Generate QR Code

`qr [text]`, generate qr code using text, this has a 800ms delay in case of generating many intermediate texts,

if `text` starts with `@`, it will be considered as an url and prepends with `http://`, e.g. `qr @localhost:3000` will get `http://localhost:3000`

- Press `Cmd+Y` to preview qr codes
- Press `Enter` to open qr code in the default application

### List all generated QR Codes

`qr`, list all generated qr codes

### Clear all generated QR Codes

`qrclearall`, clear all generated qr codes

## License




