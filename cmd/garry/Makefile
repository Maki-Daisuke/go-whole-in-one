all:
	go build -o bin/garry-init ../garry-init
	go build -o bin/garry-generate ../garry-generate
	PATH="./bin:${PATH}" garry-generate
	go build
