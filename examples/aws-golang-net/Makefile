# package
build:
	docker-compose run --rm golang make _build
.PHONY: build

package: build
	docker-compose run --rm go-shim make _package
.PHONY: package

deploy:
	docker-compose run --rm sls make _deploy
.PHONY: deploy

# target to run within container
_clean:
	@rm -rf $(HANDLER) $(HANDLER).so $(PACKAGE).zip

_deps:
	@go get -d

_build: _clean _deps
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

_package:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip