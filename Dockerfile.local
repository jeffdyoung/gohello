FROM golang:1.15 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v ./...
RUN cp hello /binout
#RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

