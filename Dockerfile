FROM eawsy/aws-lambda-go-shim:latest

ENV SERVERLESS serverless@1.21.1
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN true && \
    yum -q -e 0 -y update || true && \
    yum -q -e 0 -y install git || true && \
	go get -u github.com/golang/dep/cmd/dep
RUN curl --silent --location https://rpm.nodesource.com/setup_6.x | bash - \
 	&& yum install -y nodejs
RUN npm install -g $SERVERLESS