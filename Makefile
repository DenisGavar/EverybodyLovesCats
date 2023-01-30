.PHONY:
.SILENT:

build:
	go build -D ./.bin/bot cmd/bot/main/go

run: build
	./.bin/bot