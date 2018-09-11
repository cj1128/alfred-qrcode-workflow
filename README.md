# Alfred QR Code Workflow

[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://mit-license.org/2016)

An Alfred workflow to create QR Codes or scan them.

<p align="center">
    <img src="./intro.gif" width="600px" />
</p>

## Installation

Download [QR-Code.alfredworkflow](https://github.com/fate-lovely/alfred-qrcode-workflow/raw/master/QR-Code.alfredworkflow) and import to Alfred (require Powerpack).

## Usage

### Generate QR Code

Type `qr [text]` to generate QR code, this has a 800ms delay.

- Press `Cmd+Y` to preview the QR Code
- Press `Enter` to open QR Code in Preview

### Scan QR Code

Use [zbar-with-gbk] to do the parsing.

Select the target QR Code image, press `Cmd + Alt + \`，select `Scan QR Code` in the popup file actions.

Once you get the result, you can

- Press `Cmd + L` to preview in large text
- Press `Enter` to copy to the clipboard

## License

Released under the [MIT license](http://mit-license.org/2016).

[zbar-with-gbk]: https://github.com/fate-lovely/zbar-with-gbk
