IMAGE_LAMBDA_GO=eawsy/aws-lambda-go-shim

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

clean:
	make -f Makefile.shim clean
.PHONY: clean

dist:
	make -f Makefile.shim docker

# serverless targets
deploy:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls deploy
.PHONY: deploy

invoke:
	curl -v $(shell AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls info | grep GET | cut -f5 -d" ")
.PHONY: invoke

delete:
	AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} ENV=${ENV} sls remove -v
.PHONY: delete
