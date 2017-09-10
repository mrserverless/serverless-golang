# Serverless AWS Golang Event 
Serverless AWS APIGateway events example using: 

- [AWS Lambda Go Shim](https://github.com/eawsy/aws-lambda-go-shim)
- [API Gateway Proxy Event](https://github.com/eawsy/aws-lambda-go-event/tree/master/service/lambda/runtime/event/apigatewayproxyevt)

Each CRUD operation is it's own Lambda Function. This is convenient to hook into other 
Event Source triggers such as `Kinesis` or `SNS`.

## Usage
Setup and deploy a new project called `your-app`:

```bash
cd $GOPATH/src/path/
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-event -n your-app
cp .env.example .env
# fill in and correct any of the variables in .env
# replace `WORKDIR` in .env with `/go/src/path/your-app`
make test build deploy
make remove
```