FROM golang:1.10

# Basic tools
RUN apt-get update
RUN apt-get install -y nano net-tools htop

RUN mkdir -p /go/src/github.com/laurabcn/gobcn
WORKDIR /go/src/github.com/laurabcn/gobcn

RUN go get -u github.com/golang/dep/cmd/dep
#RUN go get -u github.com/jteeuwen/go-bindata/...
#RUN go get github.com/DATA-DOG/godog/cmd/godog

# For debugging
RUN go get github.com/derekparker/delve/cmd/dlv

COPY . /go/src/github.com/laurabcn/gobcn

# Install dependencies, look at Gopkg.toml
RUN dep ensure

# Keep alive the container without the main process running
CMD tail -f /dev/null
