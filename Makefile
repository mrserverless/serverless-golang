SLS_VERSION = 1.17.0
IMAGE_NAME ?= yunspace/serverless-golang:$(SLS_VERSION)
TAG = $(SLS_VERSION)

dockerPull:
	docker pull $(IMAGE_NAME)
.PHONY: dockerPull

dockerBuild:
	docker build -t $(IMAGE_NAME) .
.PHONY: dockerBuild

dockerPull:
	docker run --rm -it -v $(CURDIR):/opt/app $(IMAGE_NAME) bash
.PHONY: dockerPull

gitTag:
	-git tag -d $(TAG)
	-git push origin :refs/tags/$(TAG)
	git tag $(TAG)
	git push origin $(TAG)
.PHONY: gitTag