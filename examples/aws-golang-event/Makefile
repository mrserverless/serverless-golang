
# strip GOPATH to mount the correct docker WORKDIR
WORKDIR ?= $(CURDIR:$(GOPATH)%=/go%)

# package
build:
	docker-compose run --rm -w $(WORKDIR) golang make _build
.PHONY: build

test:
	docker-compose run --rm -w $(WORKDIR) golang make _test
.PHONY: test

package: build
	docker-compose run --rm go-shim make _package
.PHONY: package

deploy:
	docker-compose run --rm serverless make _deploy
.PHONY: deploy

# target to run within container
_clean:
	@rm -rf $(HANDLER) $(HANDLER).so $(PACKAGE).zip

_deps:
	@dep ensure -update

_build: _clean _deps
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

_test: _deps
	@go test

_package:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

_deploy:
	rm -fr .serverless
	sls deploy -v
