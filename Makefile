BINARY_NAME = tuner

build:
	go build -o $(BINARY_NAME) -v ./cmd

run: build
	./$(BINARY_NAME) 

install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/

lint:
	revive -formatter friendly -config revive.toml -exclude ./player/mpv ./... 

tidy:
	go mod tidy

# (build but with a smaller binary)
dist:
	go build -o $(BINARY_NAME) -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
pack: dist
	upx ./$(BINARY_NAME)
