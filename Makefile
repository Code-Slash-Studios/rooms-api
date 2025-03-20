# Variables
NAME=rooms-api
#SHELL=(shell find . -name '*.go' -type f)

build:
	go build -o ${NAME}

run:
	./${NAME}

podman:
	podman stop ${NAME}-container
	podman rm ${NAME}-container
	podman build -t ${NAME}-image .
	podman run -d -p 6000:6000 --name ${NAME}-container ${NAME}-image