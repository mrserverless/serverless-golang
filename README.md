# Serverless Golang

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com) 

[serverless](https://serverless.com/) example using `golang` and `AWS Lambda`. This is different to other `node` shim based frameworks
because it is powered by low latency `python` runtime:

- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)


## Features

- works with AWS Lambda event sources or API Gateway HTTP requests.
- go vendor support with default [dep](https://github.com/golang/dep). Can be swapped out easily.
- use `docker` for easy testing and ensure consistent dependencies across `golang`, `python` and `serverless` 

## Install

Use `serverless install` for easy installation:

	cd $GOPATH/src/your-app

	# event driven based on aws-lambda-go-shim
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-event -n your-app

	# go net style based on aws-lambda-go-net 
    serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-net -n your-app

Note you must have `$GOPATH` set and your new project must reside in `$GOPATH/src/your-app`

## Usage

All you need is `make`, `docker` and `docker-compose`

Make sure you setup your environment variables in `.env` file based on the example in `env.template`.

Then simply build and deploy the function by running:

	make package deploy

## Clients

Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build
strategic Microservices. Serving 1M+ customers with up to 500CCU.
