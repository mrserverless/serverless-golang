FROM golang:1.11.0-stretch

ENV SERVERLESS serverless@1.25.0
ENV GOPATH /go
ENV PATH $GOPATH/bin:/root/.yarn/bin:$PATH

RUN apt-get update && \
    apt-get install git
RUN	go get -u github.com/rancher/trash 
RUN curl --silent --location https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get install -y nodejs
RUN curl -o- -L https://yarnpkg.com/install.sh | bash
RUN npm install -g $SERVERLESS