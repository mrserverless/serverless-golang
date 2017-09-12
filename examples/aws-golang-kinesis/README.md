# Serverless AWS Kinesis Stream
Serverless AWS APIGateway events example using: 

- [AWS Lambda Go Shim](https://github.com/eawsy/aws-lambda-go-shim)
- [Kinesis Stream Event](https://github.com/eawsy/aws-lambda-go-event/tree/master/service/lambda/runtime/event/kinesisstreamsevt)

## Usage
Setup and deploy a new project called `your-app`:

```bash
# 1. install
cd $GOPATH/src/your-path/
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-kinesis -n your-app

# 2. configure
cp .env.example .env

# 3. create a new stream called `data-receiver` in AWS console
# 4. fill in and correct any of the variables in .env. Especially AWS_KINESIS_ARN
# 5. replace `WORKDIR` in .env with `/go/src/your-path/your-app`
# 6. deploy
make test build deploy

# 7. clean up
make remove
```

The lambda will trigger on all new messages published to the `data-receiver` stream
