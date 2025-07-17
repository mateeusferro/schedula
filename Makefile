BUILD_FOLDER=bin/
SRC_FOLDER=cmd/

build:
	go build -o ${BUILD_FOLDER} ${SRC_FOLDER}/main.go

run:
	go run ${SRC_FOLDER}/main.go

test:
	go test ./... -v