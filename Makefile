PACKAGE=github.com/gotify/go-api-client
GOTIFY_VERSION=2.0.4
BUILD=./build
TEMP_SPEC=${BUILD}/gotify.json
SWAGGER=./.tools/swagger

ifeq ($(OS),Windows_NT)
    DOWNLOAD_SWAGGER=swagger_windows_amd64.exe
else
    DOWNLOAD_SWAGGER=swagger_linux_amd64
endif

clean-tools:
	rm -rf .tools

get-tools: clean-tools
	mkdir .tools || true
	wget -O ${SWAGGER} https://github.com/go-swagger/go-swagger/releases/download/v0.19.0/${DOWNLOAD_SWAGGER}
	chmod u+x .tools/*

install:
	go mod download

obtain-spec:
	mkdir build || true
	wget -O ${TEMP_SPEC} https://raw.githubusercontent.com/gotify/server/v${GOTIFY_VERSION}/docs/spec.json

generate: obtain-spec
	${SWAGGER} generate client -f ${TEMP_SPEC} --additional-initialism=rest --skip-models

test:
	${SWAGGER} generate client --help

clean:
	rm -rf client build

