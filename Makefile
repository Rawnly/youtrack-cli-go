bin := youtrack


build:
	go build -o build/$(bin)

install:
	go build -o build/$(bin)
	mv ./build/$(bin) /usr/local/bin/

tar:
	tar -czf $(bin).tar.gz --directory=./build $(bin)
	shasum -a 256 $(bin).tar.gz

tag:
	git tag -a $(version) -m "Version: $(version)"
	git push --tags


