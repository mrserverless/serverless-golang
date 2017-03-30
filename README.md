# Serverless Golang

## What
[<img src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png" align="right" width="100">](http://serverless.com)
Simple Hello World [serverless](https://serverless.com/) service embedding a `golang` 1.8 function
into a `python` runtime using 

- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)
- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) - pending #4 refactoring

`eawsy` shims build process takes care of:

1. build golang binary into `$HANDLER.so`
1. zip `$HANDLER.so` and python shims into `$PACKAGE.zip`

`serverless` then takes care of:

1. deploy function
1. invoke function

Works with standard serverless/Lambda invoke or API Gateway HTTP requests.

## Why
Because:
- Python has the fastest cold startup time and thus minimal latency. Every other golang lambda framework out there uses node shims which is less efficient. See benchmarks: [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim and [lambda-test](https://github.com/berezovskyi/lambda-test)
- Golang is awesome and a lot of people wants [AWS Lambda Go support](https://twitter.com/awscloud/status/659795641204260864).

[Discussion thread](https://github.com/serverless/serverless/issues/2712)

## Who
Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build strategic Microservices. Serving 1M+ customers with up to 500CCU.

## Usage
Prior to running, you should have `go` and `docker` installed. Then get all the dependencies:

    make deps

Optionally update default environment variables for: `AWS_DEFAULT_REGION`, `ENV` and `PROFILE`. Then run:

    make

To clean up

    make clean delete

## ToDo
Make this into a easy to use plugin. Currently it's really a sample project.
