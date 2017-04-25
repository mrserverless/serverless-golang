# Serverless Golang Net

## Usage

Start a new project using this example:

	serverless install \
		-u https://github.com/yunspace/serverless-golang/tree/master/aws/net \
		-n my-sls-golang-net
		
Deploy the code:

	make dockerDist
	sls deploy
	
Test out the endpoint using a REST Client such as PostMan