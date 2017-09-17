# Serverless AWS Kinesis Stream
Serverless AWS APIGateway events example using: 

- [AWS Lambda Go Shim](https://github.com/eawsy/aws-lambda-go-shim)
- [Kinesis Stream Event](https://github.com/eawsy/aws-lambda-go-event/tree/master/service/lambda/runtime/event/kinesisstreamsevt)

## Usage
Setup and deploy a new project called `your-app`:

```bash
cd $GOPATH/src/your-path/
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-kinesis -n your-app
```

```bash
cd your-app
make DOTENV=.env.example dotenv
```

* create a new stream called `data-receiver` in AWS console
* fill in and correct any of the variables in .env. Especially AWS_KINESIS_ARN
* replace `WORKDIR` in .env with `/go/src/your-path/your-app`

```bash
make test build deploy
```

The lambda will trigger on all new messages published to the `data-receiver` stream

```bash
make remove
```
