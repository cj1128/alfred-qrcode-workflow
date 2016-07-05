# Alfred QR Code Workflow

alfred workflow to manipulate QR Codes on the fly, uses [json response](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/), currenty only supports [alred v3](https://www.alfredapp.com/). if you are an alfred v2 user, please update, it's definitely worth it.

## Generate QR Code

`qr [text]` this has a 800ms delay in case of generating many intermediate texts

if `text` starts with `@`, it will be considered as an url and prepends with `http://`

`qr @localhost:3000` will get `http://localhost:3000`

## List all generated QR Codes

`qr`

## Clear all generated QR Codes

`qrclearall`




