all:
	go build -o bin/wio-init ../wio-init
	go build -o bin/wio-generate ../wio-generate
	PATH="./bin:${PATH}" wio-generate
	go build
