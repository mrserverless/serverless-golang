# Serverless Golang

This example uses [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) to build golang lambdas then
deploy to AWS either directly or via [serverless](https://serverless.com/) framework.

## Usage
Prior to running, you should have `go` and `docker` installed. Then get all the dependencies:

    make deps

To use `aws lambda` CLI command directly, set a environment variable for `AWS_ACCOUNT_NUMBER` and run:

    make aws-all
    make aws-invoke
    cat out.txt
    
To leverage `serverless`, set environment variables : `AWS_DEFAULT_REGION` and `ENV`. Then run:

    make sls-all
    make sls-invoke

To clean up

	make clean
	make aws-delete
    make sls-delete
    
