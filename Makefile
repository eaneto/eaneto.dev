all:
	rm -rf bin/*
	go build -o bin/eanetodev

run:
	./bin/eanetodev
