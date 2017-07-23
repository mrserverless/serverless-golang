

# entrypoints
build: clean deps
	docker-compose run --rm sls-go make _build
.PHONY: build

test: deps
	docker-compose run --rm sls-go make _test
.PHONY: test

package: build
	docker-compose run --rm sls-go make _package
.PHONY: package

deploy:
	docker-compose run --rm sls-go make _deploy
.PHONY: deploy

remove:
	docker-compose run --rm serverless make _remove
.PHONY: deploy

shell:
	docker-compose run --rm sls-go bash
.PHONY: shell
# helpers
clean:
	docker-compose run --rm sls-go make _clean
.PHONY: clean

shell:
	docker-compose run --rm sls-go bash
.PHONY: clean

deps:
	docker-compose run --rm sls-go make _deps
.PHONY: deps

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

_remove:
	sls remove -v
