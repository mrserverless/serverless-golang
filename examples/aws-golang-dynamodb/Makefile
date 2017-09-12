HANDLER=handler
PACKAGE=package

################
# Entry Points #
################
build: .env deps
	docker-compose run --rm sls-go make _build
.PHONY: build

test: .env deps localstack
	docker-compose run --rm sls-go make _test
	docker-compose down --remove-orphans
.PHONY: test

deploy: $(PACKAGE).zip .env
	docker-compose run --rm sls-go make _deploy
.PHONY: deploy

remove: .env
	docker-compose run --rm sls-go make _remove
.PHONY: remove

shellSls: .env
	docker-compose run --rm sls-go bash
.PHONY: shell

shellAws:
	docker-compose run --rm aws bash

###########
# helpers #
###########
# Create .env based on .env.template if .env does not exist
.env:
	@echo "Create .env with .env.template"
	cp .env.template .env

clean:
	docker-compose run --rm sls-go make _clean
.PHONY: clean

deps:
	docker-compose run --rm sls-go make _deps
.PHONY: deps

localstack: shutDown
	docker-compose up -d localstack

shutDown:
	docker-compose down --remove-orphans

# target to run within container
_clean:
	@rm -rf $(HANDLER).zip $(HANDLER).so $(PACKAGE).zip

_deps:
	@dep ensure

_build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

_test:
	@go test services

_deploy:
	rm -fr .serverless
	sls deploy -v

_remove:
	sls remove -v
