FROM eawsy/aws-lambda-go-shim:latest

RUN true \
	&& yum -q -e 0 -y update || true \
	&& yum -q -e 0 -y install git || true \
	&& go get -u github.com/golang/dep/cmd/dep