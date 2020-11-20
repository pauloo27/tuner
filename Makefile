build:
	go build -v

run: build
	./tuner

install: build
	sudo cp ./tuner /usr/bin/
