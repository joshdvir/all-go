FROM golang:1.10 AS build-env
ADD . /src
RUN cd /src && make deps && make build

FROM scratch
COPY --from=build-env /src/main /all-go
EXPOSE 9292
ENTRYPOINT [ "/all-go" ]
