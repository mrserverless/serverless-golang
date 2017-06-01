# Serverless Golang

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com) 

[serverless](https://serverless.com/) example using `golang` and `AWS Lambda`. This is different to other `node` shim based frameworks
because it is powered by low latency `python` runtime:

- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)

It works with standard Lambda events or API Gateway HTTP requests.

## Install

Use `serverless install` for easy installation:

	cd $GOPATH/src/your-path

	# event driven based on aws-lambda-go-shim
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-event -n your-project

	# go net style based on aws-lambda-go-net 
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-net -n your-project

## Usage

All you need is `make`, `docker` and `docker-compose`

Make sure you setup your environment variables in `.env` file based on the example in `env.template`.

Then siomply build and deploy the function by running:

	make package deploy

## Clients

Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build
strategic Microservices. Serving 1M+ customers with up to 500CCU.
