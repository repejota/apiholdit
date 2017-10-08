BINARY="apiholdit"
DOCKER-IMAGE="repejota/apiholdit"
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
PACKAGES = "./..."

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: install
install: prebuild
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: build
build: prebuild
	go build $(LDFLAGS) -v $(PACKAGES)

.PHONY: prebuild
prebuild:
	go-bindata -pkg ${BINARY} -o data.go data/...

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

.PHONY: clean
clean:
	go clean
	rm -rf data.go
	rm -rf coverage-all.out

# Docker

.PHONY: docker
docker: clean
	docker build -t $(DOCKER-IMAGE) .
	docker tag $(DOCKER-IMAGE) $(DOCKER-IMAGE):$(VERSION)
	docker push $(DOCKER-IMAGE)
	docker rmi $(DOCKER-IMAGE)
	docker rmi $(DOCKER-IMAGE):$(VERSION)

docker-scratch: clean build
	docker build -t $(DOCKER-IMAGE):scratch -f Dockerfile.scratch .
	docker tag $(DOCKER-IMAGE):scratch $(DOCKER-IMAGE):scratch-$(VERSION)
	docker push $(DOCKER-IMAGE)
	docker rmi $(DOCKER-IMAGE):scratch
	docker rmi $(DOCKER-IMAGE):scratch-$(VERSION)
	rm -rf qurl

docker-run: clean
	docker run -it --rm --name $(BINARY) $(DOCKER-IMAGE)

# Testing

test:
	go test -v $(PACKAGES)

cover:
	go test -cover $(PACKAGES)

cover-html:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(shell go list ./...),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out
	go tool cover -html=coverage-all.out

# Lint

lint:
	gometalinter --tests .

# Dependencies

deps:
	go get -u golang.org/x/image/font
	go get -u github.com/golang/freetype

dev-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/jteeuwen/go-bindata/...


# Documentation

godoc-serve:
	godoc -http=":9090"