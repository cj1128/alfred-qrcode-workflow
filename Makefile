build:
	go build -o workflow/qr
.PHONY: build

bundle: build
	upx --brute workflow/qr
	cd workflow && rm -rf meta.json && rm -rf qrcodes
	cd workflow && zip -r ../QR-Code.alfredworkflow .
.PHONY: bundle
