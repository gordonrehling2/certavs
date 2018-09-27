export TOOLS=./bin

deps: dep

dep:
	dep ensure

lint:
	golint $(go list ./... | grep -v /vendor/)

dbuild:
	docker build -t certavs .

drun:
	docker run -it --rm --name certavs certavs

run: 
	go run main.go
