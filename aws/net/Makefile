IMAGE_LAMBDA_GO=eawsy/aws-lambda-go-shim:latest

# serverless.yml env vars
ENV ?= dev
AWS_DEFAULT_REGION ?= ap-southeast-2
AWS_PROFILE ?= default

HANDLER ?= handler
PACKAGE ?= $(HANDLER)
GOPATH  ?= $(HOME)/go

all: clean dist deploy
.PHONY: all

# lambda shim targets
clean:
	@rm -rf $(HANDLER) $(HANDLER).so $(PACKAGE).zip
.PHONY: clean

deps:
	@go get -u -d github.com/eawsy/aws-lambda-go-core/...
	@docker pull $(IMAGE_LAMBDA_GO)
.PHONY: deps

dist: build pack

build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

pack:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

# serverless targets
deploy:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls deploy
.PHONY: deploy

remove:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls remove -v
.PHONY: delete

# dockerised target
dockerDist:
	$(call docker, make dist)
.PHONY: dockerAll

dockerShell:
	$(call docker, bash)
.PHONY: dockerShell

define docker
	docker run --rm -ti \
	  -e HANDLER=$(HANDLER) \
	  -e PACKAGE=$(PACKAGE) \
	  -v $(GOPATH):/go \
	  -v $(CURDIR):/tmp \
	  -w /tmp \
	  $(IMAGE_LAMBDA_GO) $1
endef