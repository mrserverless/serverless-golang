# Serverless AWS Kinesis Stream
Serverless AWS Kinesis events example inspired by https://github.com/pmuens/serverless-kinesis-streams

## Installation
Setup and deploy a new project called `your-app`:

```bash
cd $GOPATH/src/your-path/
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-kinesis -n your-app
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

## Usage
Publish a message
```bash
make publish
```

Observe the logs to see `data-receiver` stream triggering `data-logger`
```bash
make logger
```

## Cleanup
```bash
make remove
```
