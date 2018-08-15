FROM golang:1.10

RUN apt-get update
RUN apt-get install -y nano net-tools htop git

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN mkdir -p /go/src/github.com/laurabcn/gobcn

WORKDIR /go/src/github.com/laurabcn/gobcn

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/laurabcn/gobcn





