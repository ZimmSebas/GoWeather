GO = go

.PHONY = build run

GoWeather: build 
	./$@

build: 
	$(GO) build
