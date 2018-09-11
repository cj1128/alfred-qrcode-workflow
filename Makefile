SHELL := /bin/bash

build:
	go build -o workflow/qr
.PHONY: build

bundle: build
	upx --brute workflow/qr
	cd workflow && zip --exclude .DS_Store -r ../QR-Code.alfredworkflow .
.PHONY: bundle

zbar:
	gcc -I/tmp/local/include -L/tmp/local/lib -lzbar -lpng -ljpeg zbar/zbar.c -o workflow/zbar
.PHONY: zbar
