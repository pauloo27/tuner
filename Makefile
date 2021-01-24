build:
	go build -v

run: build
	./tuner

install: build
	sudo cp ./tuner /usr/bin/

# (build but with a smaller binary)
dist:
	go build -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
pack: dist
	upx ./tuner
