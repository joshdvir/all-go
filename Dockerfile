FROM golang:1.10 AS build-env
ADD . /go/src/all-go

FROM scratch
COPY --from=build-env /go/src/all-go/dist/linux_amd64/all-go /all-go
EXPOSE 9292
ENTRYPOINT [ "/all-go" ]
