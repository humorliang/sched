BIN_DIR=_output/bin

# If tag not explicitly set in users default to the git sha.
# TAG ?= ${shell (git describe --tags --abbrev=14 | sed "s/-g\([0-9a-f]\{14\}\)$/+\1/") 2>/dev/null || git rev-parse --verify --short HEAD}

.EXPORT_ALL_VARIABLES:

all: local

init:
	mkdir -p ${BIN_DIR}

local: init
	go build -o=${BIN_DIR}/scheduler-framework-redis -mod=vendor ./cmd/scheduler

build-linux: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=${BIN_DIR}/scheduler-framework-redis -mod=vendor  ./cmd/scheduler

image: build-linux
	docker build --no-cache . -t 192.168.15.128:5000/scheduler-framework-redis:latest
	docker push 192.168.15.128:5000/scheduler-framework-redis:latest

deploy: image
	kubectl apply -f deploy/deployment.yaml
	kubectl apply -f config/busybox.yaml

update:
	go mod download
	go mod tidy
	go mod vendor

clean:
	rm -rf _output/
	rm -f *.log

