# Serverless Golang

This example uses [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) to:

1. build golang binary into `$HANDLER.so`
2. packages everything into `$PACKAGE.zip`

And then uses [serverless](https://serverless.com/) to:

3. deploy function
4. invoke function  

## Usage
Prior to running, you should have `go` and `docker` installed. Then get all the dependencies:

    make deps

Optionally update default environment variables for: `AWS_DEFAULT_REGION`, `ENV` and `PROFILE`. Then run:

    make 

To clean up

	make clean delete	
