export TOOLS=./bin

deps: dep

dep:
	dep ensure

lint:
	golint $(go list ./... | grep -v /vendor/)

dbuild: build
	docker build -t certavs .

drun: dbuild
	docker run -it --rm --name certavs certavs

build:
	env GOOS=linux GOARCH=arm go build certavs.go

run: 
	go run certavs.go
