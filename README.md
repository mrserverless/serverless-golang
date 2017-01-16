# Serverless Golang

This example uses [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) to build golang lambdas then
deploy to AWS either directly or via [serverless](https://serverless.com/) framework.

## Usage
Prior to running, you should have `go` and `docker` installed. 

    # get dependencies
    make deps

    # bundle package.zip and deploy to AWS directly
    make
    make aws-invoke
    cat out.txt
    
To leverage `serverless`, the following environment variables are need: `AWS_DEFAULT_REGION` and `ENV`. Then simply run:

    make dist sls-deploy
    make sls-invoke

To clean up

	make clean
	make aws-delete
    make sls-delete
    
