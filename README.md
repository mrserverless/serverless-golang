# Serverless Golang

[<img src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png" align="right" width="100">](http://serverless.com)
Simple Hello World [serverless](https://serverless.com/) service embedding a `golang` 1.8 function
into a `python` runtime using [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim).

`aws-lambda-go-shim` takes care of:

1. build golang binary into `$HANDLER.so`
1. zip `$HANDLER.so` and python shims into `$PACKAGE.zip`

`serverless` then takes care of:

1. deploy function
1. invoke function

Works with standard serverless/Lambda invoke or API Gateway HTTP requests.

## Why

- Because according to the following [benchmark](https://github.com/berezovskyi/lambda-test) Python has the fastest cold startup time. Every other golang lambda framework out there uses node shims which is less efficient.
- Because golang is awesome and a lot of people wants [AWS Lambda Go support](https://twitter.com/awscloud/status/659795641204260864).

[Discussion thread](https://github.com/serverless/serverless/issues/2712)

## Usage

Prior to running, you should have `go` and `docker` installed. Then get all the dependencies:

    make deps

Optionally update default environment variables for: `AWS_DEFAULT_REGION`, `ENV` and `PROFILE`. Then run:

    make

To clean up

    make clean delete

## Benchmarks

Coming soon
