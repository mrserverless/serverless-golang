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
- seamless integration with AWS Lambda event sources or API Gateway HTTP requests.
- use `docker` and `docker-compose` for easy testing with [localstack](https://github.com/localstack/localstack) (example coming soon) and ensure consistent dependencies across `golang`, `python` and `serverless`
- docker builder image immutably baked in with: 
  - amazon linux base image for building AWS Lambda
  - go 1.9
  - vendor support via [glide](https://github.com/Masterminds/glide), [trash](https://github.com/rancher/trash) and [dep](https://github.com/golang/dep)
  - python 2.7
  - node 6
  - serverless 1.22.0

## Usage
Prerequisites:
- have `serverless`, `go`, `make` and `docker`
- have correct `$GOPATH` and your new project must reside in `$GOPATH/src/path/your-app`

See each individual example for detailed instructions:

- [Serverless Golang Event](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-event/)
- [Serverless Golang Net](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-net/)
- [Serverless Golang Kinesis](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-kinesis/)
- [Serverless Golang Graphql](https://github.com/yunspace/serverless-golang/blob/master/examples/aws-golang-graphql/)

## Clients
Currently used by [amaysim Australia](https://www.amaysim.com.au/) to build strategic Microservices across 4 business verticals.

## RoadMap
- ~~kinesis example~~
- ~~graphql example~~
- dynamodb example with localstack tests #13
- event-gateway and FDK support #17
- sls plugins example #21
- Azure support #15
