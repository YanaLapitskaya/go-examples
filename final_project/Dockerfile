FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/go-exercises/final_project
COPY ./Gopkg.toml /go/src/go-exercises/final_project

# COPY ./app /go/src/go-exercises/final_project
WORKDIR /go/src/go-exercises/final_project

RUN dep ensure
# RUN go get ./
# RUN go build