SHELL := /bin/bash

build:
	go build -o workflow/qr
.PHONY: build

bundle: build
	upx --brute workflow/qr
	cd workflow && zip -r ../QR-Code.alfredworkflow .
.PHONY: bundle
