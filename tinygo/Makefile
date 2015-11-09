.PHONY: all binary

default: binary

all: clean binary
	docker build -t tinygo .

binary:
	go build -a -tags "netgo static_build" -installsuffix netgo tinygo.go

clean:
	-rm tinygo
	-docker rmi tinygo
