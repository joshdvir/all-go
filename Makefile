
clean:
	go clean -i ./...

deps:
	go get -t ./...

build_osx:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o main .

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

serve:
	./main