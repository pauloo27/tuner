build:
		go build

run: build
	./tuner

install: build
	sudo cp ./tuner /usr/bin/
