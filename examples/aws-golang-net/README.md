# Serverless AWS Golang Net
Serverless AWS Go Net example using:
 
- [AWS Lambda Golang Net](https://github.com/eawsy/aws-lambda-go-net). 
- [Gorilla MUX](http://www.gorillatoolkit.org/pkg/mux) 

All CRUD operations are within a single Lambda Function behind `gorilla/mux`. This example is
very handy for porting over existing `golang/net` projects. However it is not readily compatible 
with other Event Source triggers such as `Kinesis` or `SNS`.

## Usage
Setup and deploy a new project called `your-app`:

```bash
cd $GOPATH/src/your-app   
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-net -n your-app
cp .env.example .env
# fill in and correct any of the variables in .env
# replace `WORKDIR` in .env with `/go/src/path/your-app`
make build deploy
```