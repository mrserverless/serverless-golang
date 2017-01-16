# Serverless Golang

Run shimless Golang in serverless framework using the eawsy/aws-lambda-go 

This example currently uses [eawsy/aws-lambda-go][https://github.com/eawsy/aws-lambda-go], which will be
deprecated soon. But it will migrate over to [eawsy/aws-lambda-go-shim][https://github.com/eawsy/aws-lambda-go-shim] 

## Usage
Prior to running, you should get `go` and `docker` installed. 

    # get dependencies
    make deps

    # package handler.zip and deploy to AWS
    make package deploy

