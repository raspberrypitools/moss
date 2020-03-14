APP="moss"
MAIN = cmd/main.go
PREFIX = /usr/local
DOCKER_NAME=moss
DOCKER_TAG=latest

webApp:
	cd web && npm run build
build: 
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	go build -v -o $(APP) $(MAIN)
docker:
	env GOOS=linux go build -o $(APP).linux $(MAIN)
	docker build -t ${DOCKER_NAME}:${DOCKER_TAG} .
	

copy:
	 mkdir -p static&&cp -r web/dist/index.html static/index.html
	# cp -rf static/* web/dist/static/
update:
	go get -u ./...
package: build webApp copy
	rm -rf $(APP).tar.gz
	tar -zcvf $(APP).tar.gz  web/dist/  etc  $(APP).linux
install:
	mv -f $(TARGET) $(PREFIX)/bin
fmt:
	go fmt ./...
lint:
	go vet ./...
test:
	go test -short ./...
test-all: lint
	go test ./...
clean:
	rm -f $(TARGET))
	rm -f $(TARGET).exe
	rm -f $(PREFIX)/bin/$(TARGET)
.PHONY: build update package install fmt lint test test-all clean all %
