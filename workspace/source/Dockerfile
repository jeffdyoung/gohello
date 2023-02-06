FROM golang:1.18 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v ./...
#RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

RUN ls 
RUN pwd

FROM scratch

COPY --from=build /go/src/app/hello /
USER 1001
EXPOSE 3000
CMD ["/hello"]