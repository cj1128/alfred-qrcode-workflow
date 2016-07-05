.PHONY: build release

.default: build

build:
	godep go build -o workflow/qr

release:
	godep go build -ldflags="-s -w" -o workflow/qr
	upx --brute workflow/qr

