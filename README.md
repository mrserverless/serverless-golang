# Serverless Golang

![alt text](./gopher.jpg "Serverless Golang Gopher")
> Gopher by [@flemay](https://github.com/flemay), inspired by [@ashleymcnamara](https://github.com/ashleymcnamara) and [Renee French](http://reneefrench.blogspot.com.au/) artworks.

[<img
src="https://rawgit.com/justserverless/awesome-serverless/master/logo_serverless.png"
align="right" width="100">](http://serverless.com)

[![serverless](http://public.serverless.com/badges/v3.svg)](http://www.serverless.com) Golang example projects. 

## Performance
For AWS, we leverage eawsy's python based [AWS Lambda Go Shim](https://github.com/eawsy/aws-lambda-go-shim) for superior 
performance compared to Node.js shims:

![Benchmark](https://raw.githubusercontent.com/eawsy/aws-lambda-go-shim/master/asset/bench_1000.png)

## Features
- works with AWS Lambda event sources or API Gateway HTTP requests.
- go vendor support with default [dep](https://github.com/golang/dep). Can be swapped out easily.
- use `docker` and `docker-compose` for easy testing with [localstack](https://github.com/localstack/localstack) (example coming soon) and ensure consistent dependencies across `golang`, `python` and `serverless`

## Usage
Prerequisites:
- have `serverless`, `go`, `make` and `docker`
- have correct `$GOPATH` and your new project must reside in `$GOPATH/src/path/your-app`

See each individual example for detailed instructions:

- [Serverless Golang Event](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-event/)
- [Serverless Golang Net](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-net/)
- [Serverless Golang Kinesis](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-kinesis/)

## Clients
Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build strategic Microservices across 4 business verticals.

## RoadMap
- ~~kinesis example~~
- dynamodb example with localstack tests #13
- event-gateway and FDK support #17
- sls plugins example #21
- Azure support #15
