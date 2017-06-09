
# strip GOPATH to mount the correct docker WORKDIR
WORKDIR ?= $(CURDIR:$(GOPATH)%=/go%)

# entrypoints
build: clean deps
	docker-compose run --rm -w $(WORKDIR) go-shim make _build
.PHONY: build

test: deps
	docker-compose run --rm -w $(WORKDIR) golang make _test
.PHONY: test

package: build
	docker-compose run --rm -w $(WORKDIR) go-shim make _package
.PHONY: package

deploy:
	docker-compose run --rm serverless make _deploy
.PHONY: deploy

# helpers
clean:
	docker-compose run --rm -w $(WORKDIR) golang make _clean
.PHONY: clean

deps:
	docker-compose run --rm -w $(WORKDIR) golang make _deps
.PHONY: build

# target to run within container
_clean:
	@rm -rf $(HANDLER) $(HANDLER).so $(PACKAGE).zip

_deps:
	@dep ensure -update

_build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

_test:
	@go test

_package:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

_deploy:
	rm -fr .serverless
	sls deploy -v
