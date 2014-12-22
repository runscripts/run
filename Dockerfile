FROM golang:1.4
MAINTAINER tobe tobeg3oogle@gmail.com

RUN apt-get update -y

# Install run
ADD . /go/src/github.com/runscripts/run
WORKDIR /go/src/github.com/runscripts/run
RUN make install

CMD run