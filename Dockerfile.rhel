FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.22-builder-multi-openshift-4.17 AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v ./...
#RUN CGO_ENABLED=0 GOOS=linux go install -v ./...

RUN ls 
RUN pwd

FROM registry.ci.openshift.org/ocp/builder:rhel-9-enterprise-base-multi-openshift-4.17

COPY --from=build /go/src/app/hello /
RUN ls
RUN pwd
USER 1001
EXPOSE 3000
CMD ["/hello"]
