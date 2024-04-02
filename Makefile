all:
	rm -rf bin/*
	go build -buildvcs=false -o bin/eanetodev

run:
	./bin/eanetodev
