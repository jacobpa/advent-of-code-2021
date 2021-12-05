all: linux windows darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/aoc advent-of-code-2021/cmd/aoc

windows:
	GOOS=windows GOARCH=amd64 go build -o build/windows/aoc.exe advent-of-code-2021/cmd/aoc

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/aoc advent-of-code-2021/cmd/aoc

clean:
	rm -rf build/*
