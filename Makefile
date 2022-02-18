BINARY_NAME = tuner
TEST_COMMAND = gotest

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

.PHONY: run
run: build
	./$(BINARY_NAME) 

.PHONY: test
test: 
	$(TEST_COMMAND) -v -cover -parallel 5 -failfast  ./... 

.PHONY: cover
cover: 
	$(TEST_COMMAND) -v -cover -coverprofile=coverage.out -parallel 5 -failfast  ./... 
	go tool cover -html=coverage.out

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install
install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/

.PHONY: lint
lint:
	revive -formatter friendly -config revive.toml ./... 

.PHONY: spell
spell:
	misspell -error ./**

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: gosec
gosec:
	gosec -tests ./... 

.PHONY: inspect
inspect: lint spell gosec staticcheck

# (build but with a smaller binary)
.PHONY: dist
dist:
	go build -o $(BINARY_NAME) -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
.PHONY: pack
pack: dist
	upx ./$(BINARY_NAME)

# "hot reload"
.PHONY: dev
dev:
	fiber dev -t ./cmd/$(BINARY_NAE)
