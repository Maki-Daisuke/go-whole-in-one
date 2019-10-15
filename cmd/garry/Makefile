all:
	go build -o garry-init "../garry-init"
	go build -o garry-generate ../garry-generate
	PATH="./:${PATH}" garry-generate
	go build
