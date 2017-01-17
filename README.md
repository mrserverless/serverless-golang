# Serverless Golang
Simple Hello World [serverless](https://serverless.com/) service embedding a `golang` 1.8 function 
into a `python` runtime using [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim).

`aws-lambda-go-shim` takes care of:

1. build golang binary into `$HANDLER.so`
2. zip `$HANDLER.so` and python shims into `$PACKAGE.zip`

`serverless` then takes care of:

3. deploy function
4. invoke function

Works with standard serverless/Lambda invoke or API Gateway HTTP requests.

## Usage
Prior to running, you should have `go` and `docker` installed. Then get all the dependencies:

    make deps

Optionally update default environment variables for: `AWS_DEFAULT_REGION`, `ENV` and `PROFILE`. Then run:

    make 

To clean up

	make clean delete	
