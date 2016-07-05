.PHONY: build bundle

.default: build

build:
	godep go build -o workflow/qr

bundle:
	godep go build -ldflags="-s -w" -o workflow/qr
	upx --brute workflow/qr
	godep go build -o workflow/qr
	cd workflow && rm -rf meta.json && rm -rf qrcodes
	cd workflow && zip -r ../qrcode.alfredworkflow .

