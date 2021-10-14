build:
	go build -o youtrack

run:
	go build -o youtrack
	./youtrack issue SF-29

install:
	mv ./youtrack /usr/local/bin/
