BINARY_NAME = tuner

build:
	go build -o $(BINARY_NAME) -v ./cmd

test:
	go test -cover -parallel 5 -failfast  ./...

run: build
	./$(BINARY_NAME) 

install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/

lint:
	revive -formatter friendly -config revive.toml ./... 

spell:
	misspell -error ./**

staticcheck:
	staticcheck ./...

gosec:
	gosec -tests ./... 

inspect: lint spell gosec staticcheck

tidy:
	go mod tidy

dev:
	fiber dev -t ./cmd

# (build but with a smaller binary)
dist:
	go build -o $(BINARY_NAME) -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
pack: dist
	upx ./$(BINARY_NAME)
