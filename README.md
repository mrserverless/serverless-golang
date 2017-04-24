# Serverless Golang

## What

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com) 

[serverless](https://serverless.com/) example using `golang` and `AWS Lambda` together. Different to other frameworks
because it is powered by low latency `python` runtime instead of the more common `node.js` shim:

- [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim)
- [eawsy/aws-lambda-go-net](https://github.com/eawsy/aws-lambda-go-net)

It works with standard Lambda events or API Gateway HTTP requests.

## Usage

Use `serverless install` for easy installation. 

For a generic event driven service using `aws-lambda-go-shim`:

    serverless install \
        -u https://github.com/yunspace/serverless-golang/tree/master/aws/event \
        -n my-golang-event-project

For a `go/net` driven project using `aws-lambda-go-net`:

    serverless install \
        -u https://github.com/yunspace/serverless-golang/tree/master/aws/net \
        -n my-golang-net-project

Prior to running, you should have `go`, `make` and `docker` installed. Then once
in your project directory run this to get all the build dependencies:

    make deps

Make sure to have set the usual environment
[variables](http://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html#cli-environment)
used for authentication on AWS.

Then deploy the function by running this command:

    make dockerDist deploy

To uninstall and clean up everything you just need to run:

    make clean remove

## How

How does it work? Quite simple actually. `eawsy` shims build process will:

1. build golang binary into `$HANDLER.so`
2. zip `$HANDLER.so` and python shims into `$PACKAGE.zip`

`serverless` framework then will:

3. deploy function on AWS
4. invoke function

## Why

This project came about because:

- This discussion [thread](https://github.com/serverless/serverless/issues/2712)
- Python has the fastest cold startup time and thus minimal cold start latency.
  Every other golang lambda framework out there uses node shims which are less
  efficient, as per these benchmarks:
  [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) and
  [lambda-test](https://github.com/berezovskyi/lambda-test)
- Golang is awesome and a lot of people
  [requested](https://twitter.com/awscloud/status/659795641204260864) AWS Lambda
  Go support.

## Who

Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build
strategic Microservices. Serving 1M+ customers with up to 500CCU.
