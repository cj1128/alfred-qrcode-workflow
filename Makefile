SHELL := /bin/bash

build:
	go build -o workflow/qr
.PHONY: build

bundle: build
	# upx --brute workflow/qr
	cd workflow && zip --exclude .DS_Store -r ../QR-Code.alfredworkflow .
.PHONY: bundle

# we must static link the dependencies
zbar:
	gcc -Izbar/include -Lzbar -L/usr/lib -lz -liconv -lzbar -lpng -ljpeg zbar/zbar.c -o workflow/zbar
.PHONY: zbar
