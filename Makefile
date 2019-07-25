SLS_VERSION = 1.48.3
IMAGE_NAME ?= muellermh/serverless-golang:$(SLS_VERSION)

dockerPull:
	docker pull $(IMAGE_NAME)

dockerBuild:
	docker build -t $(IMAGE_NAME) .

shell:
	docker run --rm -it -v $(CURDIR):/opt/app $(IMAGE_NAME) bash

gitTag:
	-git tag -d $(SLS_VERSION)
	-git push origin :refs/tags/$(SLS_VERSION)
	git tag $(TAG)
	git push origin $(SLS_VERSION)
.PHONY: gitTag