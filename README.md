# Serverless Golang

## What

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com) 

[serverless](https://serverless.com/) example using `golang` and `AWS Lambda`. This is different to other `node` shim based frameworks
because it is powered by low latency `python` runtime:

- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)

It works with standard Lambda events or API Gateway HTTP requests.

## Usage

Use `serverless install` for easy installation. 

For a generic event driven service using `aws-lambda-go-shim`:

	cd $GOPATH/src/your-path
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/aws/event -n your-project

For a `go/net` driven project using `aws-lambda-go-net`:

	cd $GOPATH/src/your-path
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/aws/net -n your-project

Prior to running, you should have `go`, `make` and `docker` installed. 

Make sure to have set the usual environment
[variables](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-environment)
used for authentication on AWS.

Then deploy the function by running this command:

    make package
	sls deploy

## Usage

Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build
strategic Microservices. Serving 1M+ customers with up to 500CCU.

If your company is 
