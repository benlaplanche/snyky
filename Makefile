hello:
	echo "Hello"

build:
	go build

run:
	go run main.go

install:
	go install github.com/benlaplanche/snyky

test:
	go test

dev:
	go build
	./snyky

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

publish:
	go build
	git commit -am 'latest build'
	git push

all: hello build