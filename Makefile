BINARY_NAME = tuner

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)

.PHONY: run
run: build
	./$(BINARY_NAME) 

.PHONY: test
test: 
	go test -v -cover -parallel 5 -failfast  ./... 

.PHONY: cover
cover: 
	go test -v -cover -coverprofile=coverage.out -parallel 5 -failfast  ./... 
	go tool cover -html=coverage.out

.PHONY: show-cover
show-cover: cover
	go tool cover -html=coverage.out

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install
install: build
	# that's terrible?
	sudo cp ./$(BINARY_NAME) /usr/local/bin/

.PHONY: lint
lint:
	revive -formatter friendly -config revive.toml ./... 

.PHONY: spell
spell:
	find . -name '*.go' -exec misspell -error {} +

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: gosec
gosec:
	gosec -tests ./... 

.PHONY: inspect
inspect: spell lint staticcheck gosec

.PHONY: install-inspect-tools
install-inspect-tools:
	go install github.com/mgechev/revive@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/client9/misspell/cmd/misspell@latest

# human readable test output
.PHONY: love
love:
ifeq ($(filter watch,$(MAKECMDGOALS)),watch)
	gotestsum --watch -- -cover ./...
else
	gotestsum -- -cover ./...
endif

.PHONY: install-dev-tools
install-dev-tools: install-swaggo
	go install gotest.tools/gotestsum@latest

# (build but with a smaller binary)
.PHONY: dist
dist:
	go build -o $(BINARY_NAME) -ldflags="-w -s" -gcflags=all=-l -v
