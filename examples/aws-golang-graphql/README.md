# Serverless AWS Golang Graphql
Serverless AWS Go Graphql example using:
 
- [AWS Lambda Golang Net](https://github.com/eawsy/aws-lambda-go-net). 
- [Graphql Go](https://github.com/neelance/graphql-go) 

## Usage
Setup and deploy a new project called `your-app`:

```bash
cd $GOPATH/src/your-path
serverless install -u https://github.com/yunspace/serverless-golang/tree/master/examples/aws-golang-graphql -n your-app
```

```bash  
cd your-app
make DOTENV=.env.example dotenv
```
* fill in and correct any of the variables in .env
* replace `WORKDIR` in .env with `/go/src/your-path/your-app`

```bash
make test build deploy
```

query your app with graphiql by installing [graphiql app](https://github.com/skevy/graphiql-app) and entering the provided gateway endpoint.

```bash
{
    hello
}
```

should return

```bash
{
  "data": {
    "hello": "serverless-golang graphql"
  }
} 
```

```bash
make remove
```
