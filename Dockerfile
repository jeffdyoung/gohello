FROM golang:1.22 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v ./...
RUN mkdir /binout
RUN cp hello /binout/hello
#RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

RUN ls 
RUN pwd
RUN ls /binout

FROM registry.access.redhat.com/ubi9/ubi

COPY --from=build /binout/hello /
RUN ls
RUN pwd
USER 1001
EXPOSE 3000
CMD ["/hello"]
