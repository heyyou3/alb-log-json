MAKEFILE_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GO_ENTRY_POINT := 'adapter/input/cli/main.go'
# TODO: バージョン情報は外部から読み込む形で実現したい
GO_LD_FLAG := -ldflags '-X "main.cliVersion=v0.1.0"'
APP_NAME := 'alblogjson'

build.darwin.arm64: clean
	@cd $(MAKEFILE_DIR) && \
	GOOS=darwin GOARCH=arm64 go build $(GO_LD_FLAG) -o $(APP_NAME) $(GO_ENTRY_POINT)

build.darwin.amd64: clean
	@cd $(MAKEFILE_DIR) && \
	GOOS=darwin GOARCH=amd64 go build $(GO_LD_FLAG) -o $(APP_NAME) $(GO_ENTRY_POINT)

build.linux.amd64: clean
	@cd $(MAKEFILE_DIR) && \
	GOOS=linux GOARCH=amd64 go build $(GO_LD_FLAG) -o $(APP_NAME) $(GO_ENTRY_POINT)

clean:
	@rm -f $(APP_NAME)
