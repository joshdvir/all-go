FROM golang:1.10 AS build-env
ADD . /go/src/all-go
RUN cd /go/src/all-go && make deps && make build

FROM scratch
COPY --from=build-env /go/src/all-go/main /all-go
EXPOSE 9292
ENTRYPOINT [ "/all-go" ]
