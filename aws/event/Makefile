IMAGE_LAMBDA_GO=eawsy/aws-lambda-go-shim:latest
HANDLER ?= handler
PACKAGE ?= $(HANDLER)
GOPATH  ?= $(HOME)/go

# package
package: clean dockerPull
	$(call dockerShim, make _package)
.PHONY: package

# helper targets
deps:
	@go get -u -d
.PHONY: deps

clean:
	@rm -rf $(HANDLER) $(HANDLER).so $(PACKAGE).zip
.PHONY: clean

dockerPull:
	@docker pull $(IMAGE_LAMBDA_GO)
.PHONY: dockerPull

shell:
	$(call dockerShim, bash)
.PHONY: shell

# target to run within container
_build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

_package: _build
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

define dockerShim
	docker run --rm -ti \
	  -e HANDLER=$(HANDLER) \
	  -e PACKAGE=$(PACKAGE) \
	  -v $(GOPATH):/go \
	  -v $(CURDIR):/tmp \
	  -w /tmp \
	  $(IMAGE_LAMBDA_GO) $1
endef