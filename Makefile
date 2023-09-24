.PHONY: dev build-mac build-windows sign-mac

dev:
	wails dev

build-mac:
	wails build -clean --platform darwin/arm64 -ldflags="-X 'misc.Env=prod'" -o mydata

build-windows-arm64:
	wails build -clean -platform windows/arm64 -ldflags="-X 'misc.Env=prod'" -o mydata-arm64.exe

build-windows-x86:
	wails build -clean -platform windows/386 -ldflags="-X 'misc.Env=prod'" -o mydata-x86.exe

build-windows-amd64:
	wails build -clean -platform windows/amd64 -ldflags="-X 'misc.Env=prod'" -o mydata-amd64.exe