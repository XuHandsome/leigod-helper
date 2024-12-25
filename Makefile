.PHONY: leigod all clean

leigod:
	go build -ldflags "-s -w" .

win: clean
	mkdir -p build
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	go generate
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/Poweroff.exe
	rm -rf ./resource.syso

all: clean win
	mkdir -p build
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/leigod-helper-linux-x64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/leigod-helper-darwin-x64


clean:
	rm -rf build/* ./resource.syso