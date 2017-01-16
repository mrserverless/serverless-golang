IMAGE_LAMBDA_GO=eawsy/aws-lambda-go

deps:
	go get -u -d github.com/eawsy/aws-lambda-go/...
.PHONY: deps

package:
	docker run --rm -v $(GOPATH):/go -v $(PWD):/tmp $(IMAGE_LAMBDA_GO)
.PHONY: package

deploy:
	sls deploy
.PHONY: deploy
