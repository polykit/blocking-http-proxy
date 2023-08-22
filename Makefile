TARGET = blocking-http-proxy
GOTARGET = github.com/polykit/$(TARGET)
REGISTRY ?= ghcr.io/polykit
VERSION ?= 0.0.1
IMAGE = $(REGISTRY)/$(BIN)
DOCKER ?= docker

all: container push

build:
	CGO_ENABLED=0 go build -v -o $(TARGET) $(GOTARGET)

container:
	$(DOCKER) build --network=host -t $(REGISTRY)/$(TARGET):latest -t $(REGISTRY)/$(TARGET):$(VERSION) .

push:
	$(DOCKER) push $(REGISTRY)/$(TARGET):latest
	$(DOCKER) push $(REGISTRY)/$(TARGET):$(VERSION)

.PHONY: all build container push

clean:
	rm -f $(TARGET)
	$(DOCKER) rmi $(REGISTRY)/$(TARGET):latest
	$(DOCKER) rmi $(REGISTRY)/$(TARGET):$(VERSION)
