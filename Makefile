IMAGE_LAMBDA_GO=eawsy/aws-lambda-go-shim

# aws-lambda-go-shim env vars
GOPATH ?= $(HOME)/go
HANDLER ?= handler
PACKAGE ?= package

# serverless.yml env vars
ENV ?= dev
AWS_DEFAULT_REGION ?= ap-southeast-2
AWS_PROFILE ?= default

all: clean dist deploy invoke
.PHONY: all

deps:
	go get -u -d github.com/eawsy/aws-lambda-go-core/...
	docker pull eawsy/aws-lambda-go-shim
.PHONY: deps

clean: _clean
.PHONY: clean

dist:
	@docker run --rm \
		-v "$(GOPATH)":/go -v "$(shell PWD)":/tmp \
		-e "HANDLER=$(HANDLER)" -e "PACKAGE=$(PACKAGE)" \
		eawsy/aws-lambda-go-shim make _dist
.PHONY: dist

# serverless targets
deploy:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls deploy
.PHONY: deploy

invoke:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls invoke -f hello
.PHONY: invoke

delete:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls remove -v
.PHONY: delete

# make targets inside docker container
_dist: _clean _build _pack _inject
	@chown $(shell stat -c '%u:%g' .) $(PACKAGE).zip
	@printf "build, pack, inject, go!\n"

_clean:
	@rm -rf $(PACKAGE).zip $(HANDLER).so

_build:
	@printf "build...\r"
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so

_pack:
	@printf "build, pack\r"
	@zip -q $(PACKAGE).zip $(HANDLER).so

_inject:
	@printf "build, pack, inject\r"
	@cd /; mv /shim $(HANDLER); zip -q -r /tmp/$(PACKAGE).zip $(HANDLER)

