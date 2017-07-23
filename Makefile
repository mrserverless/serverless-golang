SLS_VERSION = 1.20.1
IMAGE_NAME ?= yunspace/serverless-golang:$(SLS_VERSION)

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
	-git tag -d $(SLS_VERSION)
	-git push origin :refs/tags/$(SLS_VERSION)
	git tag $(TAG)
	git push origin $(SLS_VERSION)
.PHONY: gitTag