# Serverless Golang

## What

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com) Simple Hello World
[serverless](https://serverless.com/) service embedding a `golang` 1.8 function
into a `python` runtime using

- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)
- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) -
  pending #4 refactoring

`eawsy` shims build process will:

1. build golang binary into `$HANDLER.so`
1. zip `$HANDLER.so` and python shims into `$PACKAGE.zip`

`serverless` then will:

1. deploy function on AWS
1. invoke function

It works with standard serverless/Lambda invocations or API Gateway HTTP
requests.

## Why

Because:

- Python has the fastest cold startup time and thus minimal cold start latency.
  Every other golang lambda framework out there uses node shims which are less
  efficient, as per these benchmarks:
  [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) and
  [lambda-test](https://github.com/berezovskyi/lambda-test)
- Golang is awesome and a lot of people
  [requested](https://twitter.com/awscloud/status/659795641204260864) AWS Lambda
  Go support.

It all started from this discussion
[thread](https://github.com/serverless/serverless/issues/2712)

## Who

Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build
strategic Microservices. Serving 1M+ customers with up to 500CCU.

## Usage

Install locally into your own Serverless project. For a generic event driven service, use:

    serverless install \
        -u https://github.com/yunspace/serverless-golang/aws/event \
        -n my-golang-event-project

For a `go/net` driven project, use:

    serverless install \
        -u https://github.com/yunspace/serverless-golang/net/event \
        -n my-golang-net-project

Prior to running, you should have `go`, `make` and `docker` installed. Then once
in your project directory run this to get all the build dependencies:

    make deps

Make sure to have set the usual environment
[variables](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-environment)
used for authentication on AWS.

Then deploy the function by running this command:

    make

To uninstall and clean up everything you just need to run:

    make clean delete

