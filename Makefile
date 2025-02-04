# Variables
NAME=rooms-api
#SHELL=(shell find . -name '*.go' -type f)

build:
	go build -o ${NAME}

run:
	./${NAME}